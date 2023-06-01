package main

import "testing"

func helloWorldTest(t *testing.T) {

	got := Hello()
	want := "Hello, World"
	t.Errorf("got %q want %q .", got, want)
}
