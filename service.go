package ask

import (
	"bufio"
	"encoding"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

// Service is the handler to ask values.
type Service struct {
	Prototype
}

var static Service

// envar gets value from environment variable
func (s *Service) envar() string {
	return os.Getenv(s.Prototype.Envar)
}

// defaultValue gets default value from the prototype and environment variable
func (s *Service) defaultValue() string {
	def := s.Prototype.Default
	if env := s.envar(); env != "" {
		def = env
	}
	return def
}

// prompt makes text to mean that the program is waiting for user to type in some answer.
// A message should take form a question.
func (s *Service) prompt() string {
	p := s.Prototype
	prompt := p.Message
	def := s.defaultValue()
	if def != "" {
		if p.Hidden {
			def = "*****"
		}
		prompt += " [" + def + "]"
	}
	if p.Column > 0 {
		prompt = fmt.Sprintf(fmt.Sprintf("%%%ds", p.Column), prompt)
	}
	return prompt + " : "
}

func (s *Service) reader() io.Reader {
	if s.Prototype.Reader != nil {
		return s.Prototype.Reader
	}
	return os.Stdin
}

func (s *Service) writer() io.Writer {
	if s.Prototype.Writer != nil {
		return s.Prototype.Writer
	}
	return os.Stdout
}

func (s *Service) isOptional(input string) error {
	if input != "" {
		return nil
	}
	if s.Prototype.Optional {
		return ErrSkip
	}
	return errors.New("it must not be empty")
}

func (s *Service) isInEnum(input string) error {
	enum := s.Prototype.Enum
	for _, e := range enum {
		if e == input {
			return nil
		}
	}

	var hint string
	switch len(enum) {
	case 0:
		return nil
	case 1:
		hint = enum[0]
	case 2:
		hint = strings.Join(enum, " OR ")
	default:
		hint = strings.Join(enum[0:len(enum)-1], ", ") + " OR " + enum[len(enum)-1]
	}
	return fmt.Errorf("it must be %s", hint)
}

func (s *Service) isMatched(input string) error {
	matcher := s.Prototype.Matcher
	if matcher == nil {
		return nil
	}
	if matcher.MatchString(input) {
		return nil
	}
	return fmt.Errorf("it must be matched for %s", s.Prototype.Matcher.String())
}

func (s *Service) isValid(input string) error {
	validator := s.Prototype.Validation
	if validator == nil {
		return nil
	}
	if err := validator(input); err != nil {
		return err
	}
	return nil
}

func (s *Service) getInput() (string, error) {
	if s.Prototype.Hidden && s.Prototype.Reader == nil {
		buf, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		fmt.Fprintln(s.writer())
		return string(buf), err
	}

	scanner := bufio.NewScanner(s.reader())
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", io.EOF
}

// Doer will do.
type Doer interface {
	Do() error
}

type doFunc func() error

func (d doFunc) Do() error { return d() }

// ErrSkip will skip to ask a value.
var ErrSkip = errors.New("skip")

// Var sets encoding.TextUnmarshaler.
func (s *Service) Var(parse encoding.TextUnmarshaler) Doer {
	return doFunc(func() error {
		return s.Ask(parse)
	})
}

type UnmarshalFunc func([]byte) error

func (f UnmarshalFunc) UnmarshalText(raw []byte) error {
	return f(raw)
}

type ParseFunc func(string) error

func (f ParseFunc) UnmarshalText(raw []byte) error {
	return f(string(raw))
}

// AskFunc will get a value from input and pass it to unmarshal func.
func (s *Service) AskUnmarshalFunc(unmarshal UnmarshalFunc) error {
	return s.Ask(unmarshal)
}

// AskParseFunc will get a value from input and pass it to parser func.
func (s *Service) AskParseFunc(parse ParseFunc) error {
	return s.Ask(parse)
}

// Ask will get a value from input and pass it to encoding.TextUnmarshaler.
func (s *Service) Ask(unmarshaler encoding.TextUnmarshaler) error {
	if bef := s.Prototype.Before; bef != nil {
		switch err := bef(); err {
		case nil:
			// continue
		default:
			return err
		}
	}
	for i := 0; s.Prototype.Limit < 1 || i < s.Prototype.Limit; i++ {
		if bef := s.Prototype.BeforeEach; bef != nil {
			switch err := bef(i); err {
			case nil:
				// continue
			default:
				return err
			}
		}
		switch err := s.askOnce(unmarshaler); err {
		case ErrSkip:
			return nil
		case nil:
			return nil
		default:
			fmt.Fprintln(s.writer(), err.Error())
		}
	}
	return errors.New("asked over the limit")
}

func (s *Service) askOnce(unmarshaler encoding.TextUnmarshaler) error {
	prompt := s.prompt()

	fmt.Fprint(s.writer(), prompt)

	input, err := s.getInput()
	if err != nil {
		return err
	}
	if input == "" {
		if def := s.defaultValue(); def != "" {
			input = def
		}
	}

	for _, p := range []ParseFunc{
		s.isOptional,
		s.isInEnum,
		s.isMatched,
		s.isValid,
		func(s string) error { return unmarshaler.UnmarshalText([]byte(s)) },
	} {
		switch err := p(input); err {
		case nil:
			// continue
		default:
			return err
		}
	}
	return nil
}
