package main

func main() {
}
func Bonus(sale []int) int {
	generalBonuses :=0
	const minBonuses  = 10_000
	const percent = 5
	for _, value := range sale{
		if value >= minBonuses{
			generalBonuses += (value - minBonuses) * percent / 100
		}
	}
	return generalBonuses
}
