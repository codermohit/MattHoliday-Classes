package hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Jonas"},
			result: "Hello, Jonas!",
		}, {
			items:  []string{"Jonas", "John", "Jane"},
			result: "Hello, Jonas, John, Jane!",
		},
	}

	for _, st := range subtests {
		want := st.result
		got := SayHello(st.items)

		if got != want {
			t.Errorf("want %s , got %s for %v", want, got, st.items)
		}
	}
}
