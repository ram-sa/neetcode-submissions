func productExceptSelf(nums []int) []int {
	//[ 1, 2, 4, 6]  Input
	
	//[  , 1, 2, 8]  Prefix
	//[48,24, 6,  ]  Suffix
	
	//[48,24,12, 8]  Products

	pfx := make([]int, len(nums))
	pfx[0] = 1
	
	for i := 1; i < len(nums); i++ {
		pfx[i] = pfx[i-1] * nums[i-1]
	}

	sfx := make([]int, len(nums))
	sfx[len(nums) - 1] = 1

	for j := len(nums) - 2; j > -1; j-- {
		sfx[j] = sfx[j+1] * nums[j+1]
	}

	prds := make([]int, len(nums))
	
	for k := 0; k < len(prds); k++ {
		prds[k] = pfx[k] * sfx[k]
	}

	return prds
}
