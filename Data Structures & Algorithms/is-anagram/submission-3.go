func isAnagram(s string, t string) bool {
	
	if len(s) != len(t){
		return false
	}
	
	runeS, runeT := []rune(s), []rune(t)

	sortRunes(runeS)
	sortRunes(runeT)

	for i, rS := range runeS {
		if rS != runeT[i] {
			return false
		}
	}

	return true

}

func sortRunes(r []rune){
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
}