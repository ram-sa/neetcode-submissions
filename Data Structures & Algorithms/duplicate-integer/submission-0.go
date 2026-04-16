func hasDuplicate(nums []int) bool {
    m := make(map[int]struct{})

	for _, num := range nums {
		_, ok := m[num]
		if ok {
			return true
		} else {
			m[num] = struct{}{}
		}
	}

	return false
}
