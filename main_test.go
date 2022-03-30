package main

import "testing"

func Test_parseBase64EncodedEvent(t *testing.T) {
	parsed, err := parseBase64EncodedEvent(data)
	if err != nil {
		t.Fatal(err)
	}

	const expected = "page_view"
	val, err := parsed.GetValue("event_name")
	if err != nil {
		t.Fatal(err)
	}

	actual, ok := val.(string)
	if !ok {
		t.Errorf("expected string value, got %T", val)
	}

	if actual != expected {
		t.Errorf("expected: %s, got: %s", expected, actual)
	}
}
