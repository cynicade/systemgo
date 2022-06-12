package processselector

import (
	"testing"
)

func TestParseLoad(t *testing.T) {
	arr := [...]string{"loaded", "not found", "bad setting", "error", "masked"}
	for _, load := range arr {
		got := parseLoad(load).String()
		want := load

		if want != got {
			t.Errorf(`ParseLoad("%s").String() = %q, want match for %#q`, load, got, want)
		}
	}
}

func TestParseActive(t *testing.T) {
	arr := [...]string{"active", "reloading", "inactive", "failed", "activating", "deactivating"}
	for _, load := range arr {
		got := parseActive(load).String()
		want := load

		if want != got {
			t.Errorf(`ParseActive("%s").String() = %q, want match for %#q`, load, got, want)
		}
	}
}
