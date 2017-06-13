// Code generated by ask-gen. DO NOT EDIT!!

package ask

import (
	"io"
	"regexp"
)

// Prototype gives handler some options
type Prototype struct {
	Hidden     bool
	Default    string
	Envar      string
	Validation func(string) error
	Matcher    *regexp.Regexp
	Enum       []string
	Column     int
	Optional   bool
	Limit      int
	Reader     io.Reader
	Writer     io.Writer
	Message    string
	Before     func() error
	BeforeEach func(int) error
}

// Hidden makes new Handler with new Hidden option
func (h Handler) Hidden(v bool) *Handler {
	n := h.Prototype
	n.Hidden = v
	return &Handler{Prototype: n}
}

// Hidden makes new Handler with new Hidden option
func Hidden(v bool) *Handler {
	return static.Hidden(v)
}

// Default makes new Handler with new Default option
func (h Handler) Default(v string) *Handler {
	n := h.Prototype
	n.Default = v
	return &Handler{Prototype: n}
}

// Default makes new Handler with new Default option
func Default(v string) *Handler {
	return static.Default(v)
}

// Envar makes new Handler with new Envar option
func (h Handler) Envar(v string) *Handler {
	n := h.Prototype
	n.Envar = v
	return &Handler{Prototype: n}
}

// Envar makes new Handler with new Envar option
func Envar(v string) *Handler {
	return static.Envar(v)
}

// Validation makes new Handler with new Validation option
func (h Handler) Validation(v func(string) error) *Handler {
	n := h.Prototype
	n.Validation = v
	return &Handler{Prototype: n}
}

// Validation makes new Handler with new Validation option
func Validation(v func(string) error) *Handler {
	return static.Validation(v)
}

// Matcher makes new Handler with new Matcher option
func (h Handler) Matcher(v *regexp.Regexp) *Handler {
	n := h.Prototype
	n.Matcher = v
	return &Handler{Prototype: n}
}

// Matcher makes new Handler with new Matcher option
func Matcher(v *regexp.Regexp) *Handler {
	return static.Matcher(v)
}

// Enum makes new Handler with new Enum option
func (h Handler) Enum(v []string) *Handler {
	n := h.Prototype
	n.Enum = v
	return &Handler{Prototype: n}
}

// Enum makes new Handler with new Enum option
func Enum(v []string) *Handler {
	return static.Enum(v)
}

// Column makes new Handler with new Column option
func (h Handler) Column(v int) *Handler {
	n := h.Prototype
	n.Column = v
	return &Handler{Prototype: n}
}

// Column makes new Handler with new Column option
func Column(v int) *Handler {
	return static.Column(v)
}

// Optional makes new Handler with new Optional option
func (h Handler) Optional(v bool) *Handler {
	n := h.Prototype
	n.Optional = v
	return &Handler{Prototype: n}
}

// Optional makes new Handler with new Optional option
func Optional(v bool) *Handler {
	return static.Optional(v)
}

// Limit makes new Handler with new Limit option
func (h Handler) Limit(v int) *Handler {
	n := h.Prototype
	n.Limit = v
	return &Handler{Prototype: n}
}

// Limit makes new Handler with new Limit option
func Limit(v int) *Handler {
	return static.Limit(v)
}

// Reader makes new Handler with new Reader option
func (h Handler) Reader(v io.Reader) *Handler {
	n := h.Prototype
	n.Reader = v
	return &Handler{Prototype: n}
}

// Reader makes new Handler with new Reader option
func Reader(v io.Reader) *Handler {
	return static.Reader(v)
}

// Writer makes new Handler with new Writer option
func (h Handler) Writer(v io.Writer) *Handler {
	n := h.Prototype
	n.Writer = v
	return &Handler{Prototype: n}
}

// Writer makes new Handler with new Writer option
func Writer(v io.Writer) *Handler {
	return static.Writer(v)
}

// Message makes new Handler with new Message option
func (h Handler) Message(v string) *Handler {
	n := h.Prototype
	n.Message = v
	return &Handler{Prototype: n}
}

// Message makes new Handler with new Message option
func Message(v string) *Handler {
	return static.Message(v)
}

// Before makes new Handler with new Before option
func (h Handler) Before(v func() error) *Handler {
	n := h.Prototype
	n.Before = v
	return &Handler{Prototype: n}
}

// Before makes new Handler with new Before option
func Before(v func() error) *Handler {
	return static.Before(v)
}

// BeforeEach makes new Handler with new BeforeEach option
func (h Handler) BeforeEach(v func(int) error) *Handler {
	n := h.Prototype
	n.BeforeEach = v
	return &Handler{Prototype: n}
}

// BeforeEach makes new Handler with new BeforeEach option
func BeforeEach(v func(int) error) *Handler {
	return static.BeforeEach(v)
}
