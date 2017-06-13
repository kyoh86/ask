package ask

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

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
		prompt += " [" + p.Default + "]"
	}
	if p.Column > 0 {
		prompt = fmt.Sprintf(fmt.Sprintf("%%%ds", p.Column), prompt)
	}
	return prompt + ": "
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
	if s.Prototype.Optional || input != "" {
		return nil
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

type Doer interface {
	Do() error
}

type DoFunc func() error

func (d DoFunc) Do() error { return d() }

type Parser interface {
	Parse(string) error
}

type ParseFunc func(string) error

func (f ParseFunc) Parse(input string) error {
	return f(input)
}

var Skip = errors.New("skip")

func (s *Service) Var(parse Parser) Doer {
	return DoFunc(func() error {
		return s.Ask(parse)
	})
}

func (s *Service) AskFunc(parse ParseFunc) error {
	return s.Ask(parse)
}

func (s *Service) Ask(parse Parser) error {
	if bef := s.Prototype.Before; bef != nil {
		switch err := bef(); err {
		case nil:
			// continue
		case Skip:
			return nil
		default:
			return err
		}
	}
	for i := 0; s.Prototype.Limit < 1 || i < s.Prototype.Limit; i++ {
		if bef := s.Prototype.BeforeEach; bef != nil {
			switch err := bef(i); err {
			case nil:
				// continue
			case Skip:
				return nil
			default:
				return err
			}
		}
		if err := s.askOnce(parse); err != nil {
			fmt.Fprint(s.writer(), strings.Title(err.Error()))
		} else {
			return nil
		}
	}
	return errors.New("asked over the limit")
}

func (s *Service) askOnce(parse Parser) error {
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
		parse.Parse,
	} {
		switch err := p(input); err {
		case nil:
			// continue
		case Skip:
			return nil
		default:
			return err
		}
	}
	return nil
}
