func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		runeStr := []rune(str)
		sortRunes(runeStr)
		sortedStr := string(runeStr)

		slc, ok := m[sortedStr]
		if ok {
			m[sortedStr] = append(slc, str)
		} else {
			m[sortedStr] = []string{str}
		}
	}

	values := make([][]string, 0, len(m))
	for _, v := range m {
    	values = append(values, v)
	}

	return values

}

func sortRunes(r []rune){
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
}
