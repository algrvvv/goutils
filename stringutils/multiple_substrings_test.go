package stringutils

import "testing"

func TestCheckSubstring(t *testing.T) {
	t.Run("should be return true if str contains all subs", func(t *testing.T) {
		actual := CheckSubstring("[server]", "[", "]")
		if actual != true {
			t.Errorf("got %v, want %v", actual, true)
		}
	})

	t.Run("should be return false if str dont contains all subs", func(t *testing.T) {
		actual := CheckSubstring("server", "[", "]")
		if actual != false {
			t.Errorf("got %v, want %v", actual, true)
		}
	})

	t.Run("should be return false if str dont contains at least one sub", func(t *testing.T) {
		actual := CheckSubstring("server]", "[")
		if actual != false {
			t.Errorf("got %v, want %v", actual, true)
		}
	})
}
