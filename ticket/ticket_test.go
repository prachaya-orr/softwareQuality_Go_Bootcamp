package ticket

import "testing"

func TestTicketPrice(t *testing.T) {

	t.Run("shoud return 0 when age is 0", func(t *testing.T) {
		want := 13.0
		age := -1

		got := Price(age)

		if got != want {
			t.Errorf("Price(0) = %f; want %f", got, want)
		}
	})

	t.Run("shoud return 0 when age is 0", func(t *testing.T) {
		want := 0.0
		age := 0

		got := Price(age)

		if got != want {
			t.Errorf("Price(0) = %f; want %f", got, want)
		}
	})

	t.Run("Free Ticket when age under 3", func(t *testing.T) {
		want := 0.0
		age := 3

		got := Price(age)

		if got != want {
			t.Errorf("Price(3) = %f; want %f", got, want)
		}
	})

	t.Run("Ticket $15 when age is 4 year old", func(t *testing.T) {
		want := 15.0
		age := 4

		got := Price(age)

		if got != want {
			t.Errorf("Price(4) = %f; want %f", got, want)
		}
	})

	t.Run("Ticket $15 when age is 15 ", func(t *testing.T) {
		want := 15.0

		got := Price(15)

		if got != want {
			t.Errorf("Price(15) = %f; want %f", got, want)
		}
	})

	t.Run("Ticket $30 when age over 15 ", func(t *testing.T) {
		want := 30.0
		age := 16

		got := Price(age)

		if got != want {
			t.Errorf("Price(16) = %f; want %f", got, want)
		}
	})

	t.Run("Ticket $30 when age is 50 ", func(t *testing.T) {
		want := 30.0

		got := Price(50)

		if got != want {
			t.Errorf("Price(50) = %f; want %f", got, want)
		}
	})

	t.Run("Ticket $5 when age over 50 ", func(t *testing.T) {
		want := 5.0

		got := Price(51)

		if got != want {
			t.Errorf("Price(51) = %f; want %f", got, want)
		}
	})

}
