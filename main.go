package main

func main() {
}
func Bonus(salary []int) int {
	generalBonuses :=0
	const MinBonuses  = 10_000
	const perсent = 5
	const MaxPercent = 100
	for _, value := range salary{
		if value >= MinBonuses {
			generalBonuses += (value - MinBonuses) * perсent / MaxPercent
		}
	}
	return generalBonuses
}
