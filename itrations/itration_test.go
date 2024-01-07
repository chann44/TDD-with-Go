package itrations

import "testing"


func TestIntegers(t *testing.T) {
	t.Run("Test itrations", func(t *testing.T) {
	repeated := Repeat(3, "a")	
	expected := "aaa"

	if(expected != repeated) {
		t.Errorf("got: %q want: %q", repeated, expected)
	}
	})
}


// this code is for bench marking

// func BenchmarkRepeat(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Repeat("a")
// 	}
// }