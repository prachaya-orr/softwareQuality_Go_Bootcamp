package sum

import "testing"

// Step 3
func TestSum(t *testing.T) {
	t.Run("should multi parameters", func(t *testing.T) {
		want := 0

		got := sum([]int{}...)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("including sign integer", func(t *testing.T) {
		want := 7
		xs := []int{2, 3, 3, -1}
		// when xs in parameter we spread with xs...
		got := sum(xs...)

		if got != want {
			t.Error("Expected", 7, "Got", got)
		}
	})

	t.Run("Should return 3 when 1 and 2", func(t *testing.T) {
		// Arrange
		want := 3

		// Act
		got := sum(1, 2)

		// Assertion
		if got != want {
			t.Errorf("sum(1,2) = %d; want 3", got)
		}

	})

	t.Run("Should return -2 when -1 and -1", func(t *testing.T) {
		want := -2

		got := sum(-1, -1)

		if got != want {
			t.Errorf("sum(-1,-1) = %d; want -2", got)
		}

	})

}

// Step 2

// func TestSum(t *testing.T) {
// 	t.Run("Should return 3 when 1 and 2", func(t *testing.T) {

// 		got := sum(1, 2)

// 		if got != 3 {
// 			t.Errorf("sum(1,2) = %d; want 3", got)
// 		}

// 	})

// 	t.Run("Should return -2 when -1 and -1", func(t *testing.T) {

// 		got := sum(-1, -1)

// 		if got != -2 {
// 			t.Errorf("sum(-1,-1) = %d; want -2", got)
// 		}

// 	})

// }

// Step 1

// func TestSumShouldReturn1wheninput1and0(t *testing.T) {
// 	got := sum(1, 0)

// 	if got != 1 {
// 		t.Errorf("sum(1,0) = %d; want 1 ", got)
// 	}
// }

// func TestSumShouldReturnMinus2wheninputMinus1andMinus1(t *testing.T) {
// 	got := sum(-1, -1)

// 	if got != -2 {
// 		t.Errorf("sum(-1,-1) = %d; want = -2 ", got)
// 	}
// }
