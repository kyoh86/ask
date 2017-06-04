package main

import (
	"fmt"
	"regexp"

	"github.com/kyoh86/ask"
)

func main() {
	{
		out, err := ask.Hidden(true).String("hidden")
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	{
		out, err := ask.Default("foo").String("default")
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	{
		out, err := ask.Env("ENVAR").String("envar")
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	{
		out, err := ask.Validation(func(s string) error { return nil }).String("validation")
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	{
		out, err := ask.Matcher(regexp.MustCompile(".*")).String("matcher")
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	{
		out, err := ask.Enum([]string{"foo", "bar"}).String("enum")
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	{
		out, err := ask.Column(42).String("column")
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	{
		out, err := ask.Optional(true).String("optional")
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
	// if the answer is invalid, it'll return error except for retrying just twice.
	{
		out, err := ask.Limit(2).String("limit(2)")
		if err != nil {
			panic(err)
		}
		fmt.Printf("echo: %s\n", *out)
	}
}
