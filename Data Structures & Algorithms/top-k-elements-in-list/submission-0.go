func topKFrequent(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}
	
	m := make(map[int]int)

	for _, num := range nums {
		m[num] += 1
	}

	// Convert map to slice

	type kv struct {
		Key   int
		Value int
	}
	var sorted []kv
	for k, v := range m {
		sorted = append(sorted, kv{k, v})
	}
	// Sort by value
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	result := make([]int, k)

	for i := 0; i < k; i++ {
		result[i] = sorted[i].Key
	}

	return result
}
