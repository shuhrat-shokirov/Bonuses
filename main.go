package main

func main() {
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
