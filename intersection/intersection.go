package intersection

func Intersection(a, b []int) []int {
	result := []int{}
	counter := make(map[int]int)

	for _, elem := range a {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1
		} else {
			counter[elem] += 1
		}
	}

	for _, elem := range b {
		if count, ok := counter[elem]; ok && count > 0 {
			counter[elem] -= 1
			result = append(result, elem)
		}
	}

	return result
}
