func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})

	for _, n := range nums {
		m[n] = struct{}{}
	}

	longest := 0

	for n := range m {
		if _, ok := m[n-1]; ok {
			continue
		}

		length := 1

		for {
			if _, ok := m[n+length]; !ok {
				break
			}
			length++
		}

		if length > longest {
			longest = length
		}
	}

	return longest
}
