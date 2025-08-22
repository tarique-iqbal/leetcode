package solution

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	left, right := 0, len(people)-1
	boats := 0

	for left <= right {
		if people[left]+people[right] <= limit {
			left++
			right--
		} else {
			right--
		}
		boats++
	}

	return boats
}
