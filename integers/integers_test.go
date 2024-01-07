package integers

import "testing"


func TestIntegers(t *testing.T) {
	t.Run("is number zero", func(t *testing.T) {
		got := Integers()
		want := 0

		if(got != want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}