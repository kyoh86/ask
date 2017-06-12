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

type Handler struct {
	Prototype
}

var static Handler

// env gets value from environment variable
func (h *Handler) env() string {
	return os.Getenv(h.Prototype.Env)
}

// defaultValue gets default value from the prototype and environment variable
func (h *Handler) defaultValue() string {
	def := h.Prototype.Default
	if env := h.env(); env != "" {
		def = env
	}
	return def
}

// prompt makes text to mean that the program is waiting for user to type in some answer.
// A message should take form a question.
func (h *Handler) prompt() string {
	proto := h.Prototype
	prompt := proto.Message
	def := h.defaultValue()
	if def != "" {
		prompt += " [" + proto.Default + "]"
	}
	if proto.Column > 0 {
		prompt = fmt.Sprintf(fmt.Sprintf("%%%ds", proto.Column), prompt)
	}
	return prompt + ": "
}

func (h *Handler) reader() io.Reader {
	if h.Prototype.Reader != nil {
		return h.Prototype.Reader
	}
	return os.Stdin
}

func (h *Handler) writer() io.Writer {
	if h.Prototype.Writer != nil {
		return h.Prototype.Writer
	}
	return os.Stdout
}

func (h *Handler) isOptional(input string) error {
	if h.Prototype.Optional || input != "" {
		return nil
	}
	return errors.New("empty")
}

func (h *Handler) isInEnum(input string) error {
	enum := h.Prototype.Enum
	var hint string
	switch len(enum) {
	case 0:
		return nil
	case 1:
		hint = enum[0]
	case 2:
		hint = strings.Join(enum, " or ")
	default:
		hint = strings.Join(enum[0:len(enum)-1], ", ") + enum[len(enum)-1]
	}
	for _, e := range enum {
		if e == input {
			return nil
		}
	}
	return fmt.Errorf("not in options %s", hint)
}

func (h *Handler) isMatched(input string) error {
	matcher := h.Prototype.Matcher
	if matcher == nil {
		return nil
	}
	if matcher.MatchString(input) {
		return nil
	}
	return fmt.Errorf("unmatched for %s", h.Prototype.Matcher.String())
}

func (h *Handler) isValid(input string) error {
	validator := h.Prototype.Validation
	if validator == nil {
		return nil
	}
	if err := validator(input); err != nil {
		return err
	}
	return nil
}

func (h *Handler) getInput() (string, error) {
	if h.Prototype.Hidden && h.Prototype.Reader == nil {
		buf, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		fmt.Fprintln(h.writer())
		return string(buf), err
	}

	scanner := bufio.NewScanner(h.reader())
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", io.EOF
}

type Doer interface {
	Do() error
}

type doerImpl func() error

func (d doerImpl) Do() error { return d() }

type Collector []Doer

func (c Collector) Do() error {
	for _, d := range c {
		if err := d.Do(); err != nil {
			return err
		}
	}
	return nil
}

func (c *Collector) Push(d Doer) {
	*c = append(*c, d)
}

type Parser func(string) error

func (h *Handler) Ask(parse Parser) error {
	for i := 0; h.Prototype.Limit < 1 || i < h.Prototype.Limit; i++ {
		if err := h.AskOnce(parse); err != nil {
			fmt.Fprintf(h.writer(), "invalid input: %s\n", err)
		} else {
			return nil
		}
	}
	return errors.New("asked over the limit")
}

func (h *Handler) AskOnce(parse Parser) error {
	prompt := h.prompt()

	fmt.Fprint(h.writer(), prompt)

	input, err := h.getInput()
	if err != nil {
		return err
	}
	if input == "" {
		if def := h.defaultValue(); def != "" {
			input = def
		}
	}

	for _, p := range []Parser{
		h.isOptional,
		h.isInEnum,
		h.isMatched,
		h.isValid,
		parse,
	} {
		if err := p(input); err != nil {
			return err
		}
	}
	return nil
}
