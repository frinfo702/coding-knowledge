package main

import "testing"

func TestWithOnError(t *testing.T) {
    var called bool
    errHook := func(error) { called = true }
    s := NewServer(WithOnError(errHook))
    s.onError(nil)
    if !called {
        t.Fatalf("onError hook not executed")
    }
}
