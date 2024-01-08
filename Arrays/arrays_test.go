package arrays

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T)  {

	t.Run("test sum any of numbers", func(t *testing.T) {
			numbers := []int{1, 2, 3}
			
			got := sum(numbers)
			want := 6

			if got != want {
				t.Errorf("got %d want %d ", got, want)	
			}
		})	
}


func TestAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{5, 3})
	want := []int{3, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}	

}


func TestSumAllTais(t *testing.T) {

	checkSum := func (t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}	
	}

	

	t.Run("some slices tals some", func(t *testing.T) {
	got := SumALLTails([]int{1, 2}, []int{5, 3})
	want := []int{2, 3}
	checkSum(t, got, want)

	})

	t.Run("one empty slices", func(t *testing.T) {
	got := SumALLTails([]int{}, []int{5, 3})
	want := []int{0, 3}
	checkSum(t, got, want)
	})

}