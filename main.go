package main

func main() {
}
func Bonus(sale []int) int {
	generalBonuses :=0
	const minbonuses  = 10_000
	const perсent = 5
	for _, value := range sale{
		if value >= minbonuses{
			generalBonuses += (value - minbonuses) * perсent / 100
		}
	}
	return generalBonuses
}
