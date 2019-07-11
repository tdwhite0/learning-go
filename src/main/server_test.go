package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hola, ' when language is Spanish", func(t *testing.T) {
		got := Hello("Tomas", "Spanish")
		want := "Hola, Tomas"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Bonjour, ' when language is French", func(t *testing.T) {
		got := Hello("Tom", "French")
		want := "Bonjour, Tom"

		assertCorrectMessage(t, got, want)
	})

	t.Run("the greeting prefix of Spanish should be Hola", func(t *testing.T) {
		got := greetingPrefix("Spanish")
		want := "Hola, "

		assertCorrectMessage(t, got, want)
	})
}
