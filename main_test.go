package main

import "testing"

func TestBonus(t *testing.T) {
	testCases := []struct {
		name string
		sale []int
		want int
	}{
		{"if all sales receive a bonus", []int{12_000, 18_000, 15_000, 13_000},900},
		{"if some sales get a bonus", []int{12_000, 10_000, 15_000, 8_000},350},
		{"if all sales are not worth the bonus", []int{7_000, 8_000, 9_000, 9_500},0},
		{"if all sales are worth the bonus but do not receive", []int{10_000, 10_000, 10_000, 10_000},0},
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got := Bonus(testCase.sale)
		if got != testCase.want {
			t.Error("nps with test:", testCase.name, "got:", got, "want:", testCase.want)
		}
	}
}