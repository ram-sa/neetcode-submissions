type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encStr strings.Builder
	encStr.WriteString(fmt.Sprintf("%02d", len(strs)))

	for _, str := range strs {
		encStr.WriteString(fmt.Sprintf("%03d", len(str)))
		encStr.WriteString(str)
	}
	fmt.Println(encStr.String())
	return encStr.String()
}

func (s *Solution) Decode(encoded string) []string {
	arrLen, _ := strconv.Atoi(string(encoded[0:2]))
	decArr := make([]string, arrLen)

	j := 2
	for i := 0; i < arrLen; i++ {
		strLen, _ := strconv.Atoi(string(encoded[j:j+3]))
		fmt.Println(strLen)
		j += 3
		decArr[i] = string(encoded[j:j+strLen])
		fmt.Println(decArr[i])
		j += strLen
	}

	return decArr
}
