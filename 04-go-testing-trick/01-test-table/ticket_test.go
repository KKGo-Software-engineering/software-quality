package ticket

import "testing"

func TestTicketPrice(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want float64
	}{
		{"Free Ticket when age is 0", 0, 0.0},
		{"Free Ticket when age under 3", 3, 0.0},
		{"Ticket $15 when age at 4 year old", 4, 15.0},
		{"Ticket $15 when age is 15", 15, 15.0},
		{"Ticket $30 when age over 15", 16, 30.0},
		{"Ticket $30 when age is 50", 50, 30.0},
		{"Ticket $5 when age over 50", 51, 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Price(tt.age)
			if got != tt.want {
				t.Errorf("Price(%d) = %f; want %f", tt.age, got, tt.want)
			}
		})
	}
}
