package main

import "fmt"

func main() {
	salary := []int{12_000, 8_000, 15_000, 8_000}
	fmt.Println(Bonus(salary))
}
func Bonus(salary []int) int {
	generalBonuses :=0
	const MinBonuses  = 10_000
	for _, value := range salary{
		if value >= MinBonuses {
			generalBonuses += (value - MinBonuses) * 5 / 100
		}
	}
	return generalBonuses
}
