package main

import "testing"

func TestBonus(t *testing.T) {
	testCases := []struct {
		name string
		salery []int
		want int
	}{
		{"First Employees", []int{12_000, 8_000, 15_000, 8_000},350},
		{"First Employees", []int{10_000, 8_000, 10_000, 8_000},0},
		{"First Employees", []int{9_000, 8_000, 9_000, 8_000},0},
		// TODO: Add test cases.
	}
	for _, testCase := range testCases {
		got := Bonus(testCase.salery)
		if got != testCase.want {
			t.Error("nps with test:", testCase.name, "got:", got, "want:", testCase.want)
		}
	}
}