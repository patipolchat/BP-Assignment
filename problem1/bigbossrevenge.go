package problem1

func BossBabyRevenge(input string) string {
	bullet := 0
	killSteak := 0
	for _, char := range input {
		switch char {
		case 'S':
			bullet++
			killSteak = 0
		case 'R':
			if bullet <= 0 && killSteak <= 0 {
				return "Bad Boy"
			}
			bullet--
			killSteak++
			if bullet < 0 {
				bullet = 0
			}
		default:
			return "Invalid input"
		}
	}

	if bullet == 0 {
		return "Good Boy"
	}
	return "Bad Boy"
}
