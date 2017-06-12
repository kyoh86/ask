package main

import (
	"fmt"
	"regexp"

	"github.com/kyoh86/ask"
)

func main() {
	{
		out, err := ask.Hidden(true).Message("hidden").String()
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	{
		out, err := ask.Default("foo").Message("default").String()
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	{
		out, err := ask.Env("ENVAR").Message("envar").String()
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	{
		out, err := ask.Validation(func(s string) error { return nil }).Message("validation").String()
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	{
		out, err := ask.Matcher(regexp.MustCompile(".*")).Message("matcher").String()
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	{
		out, err := ask.Enum([]string{"foo", "bar"}).Message("enum").String()
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	{
		out, err := ask.Column(42).Message("column").String()
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	{
		out, err := ask.Optional(true).Message("optional").String()
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	// if the answer is invalid, it'll return error except for retrying just twice.
	{
		out, err := ask.Limit(2).Message("limit(2)").String()
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
}
