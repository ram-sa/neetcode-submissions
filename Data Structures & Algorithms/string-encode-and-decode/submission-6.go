type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encStr := fmt.Sprintf("%02d", len(strs))

	for _, str := range strs {
		encStr += fmt.Sprintf("%03d", len(str)) + str
	}

	return encStr
	//var encStr strings.Builder
	
	//If you pad the length, you don't need a demilitator at all.
	//You trade a miniscule amount of memory for some if operations.
	//The constraints keep the array size at two digits (<100) and 
	//String sizes at three (<200)
	//encStr.WriteString(fmt.Sprintf("%02d", len(strs)))

	//for _, str := range strs {
	//	encStr.WriteString(fmt.Sprintf("%03d", len(str)))
	//	encStr.WriteString(str)
	//}
	//fmt.Println(encStr.String())
	//return encStr.String()
}

func (s *Solution) Decode(encoded string) []string {
	arrLen, _ := strconv.Atoi(string(encoded[0:2]))
	decArr := make([]string, arrLen)

	j := 2
	for i := 0; i < arrLen; i++ {
		strLen, _ := strconv.Atoi(string(encoded[j:j+3]))
		j += 3
		decArr[i] = string(encoded[j:j+strLen])
		j += strLen
	}

	return decArr
}
