package arrays

import (
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