package best_time_to_buy_and_sell_stock_121

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Test Case 1": {input: []int{7, 1, 5, 3, 6, 4}, want: 5},
		"Test Case 2": {input: []int{7, 6, 4, 3, 1}, want: 0},
		"Test Case 3": {input: []int{4}, want: 0},
		"Test Case 4": {input: []int{0, 0}, want: 0},
		"Test Case 5": {input: []int{42, 7, 15, 8, 36, 4, 22}, want: 29},
	}

	for name, tc := range tests {
		got := maxProfit(tc.input)

		if got != tc.want {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}
