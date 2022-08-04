package main

import (
	"testing"
)

func TestBasicAuth(t *testing.T) {
	expected := "Zm9vYmFyOmZvb2Jhcg=="
	actual := BasicAuth("foobar", "foobar")

	if expected != actual {
		t.Fatalf(`actual == %s ; expected == %s`, actual, expected)
	}
}
