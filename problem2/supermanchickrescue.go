package problem2

import "fmt"

func SupermanChickenRescue(chickenCount int, lengthOfRoof int, chickenPos []int) (int, error) {
	if chickenCount == 0 {
		return 0, nil
	}
	if chickenCount != len(chickenPos) {
		return 0, fmt.Errorf("chickenCount and length of chickenPos are not equal")
	}

	maximumChicken := 0
	for currentIndex, pos := range chickenPos {
		stopPos := pos + lengthOfRoof - 1
		rescueCount := 0
		for i := currentIndex; i < chickenCount; i++ {
			if chickenPos[i] > stopPos {
				break
			}
			rescueCount++
		}
		maximumChicken = max(maximumChicken, rescueCount)
	}

	return maximumChicken, nil
}
