package main

import (
	"testing" 
	"github.com/ivymmurithi/learn-go-with-tests/helpers"
)

func TestHello(t *testing.T) {
	//subtests --> testing a thing and then describing different scenarios
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		helpers.AssertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		helpers.AssertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Ivy", "Spanish")
		want := "Hola, Ivy"
		helpers.AssertCorrectMessage(t, got, want)
	})
}

