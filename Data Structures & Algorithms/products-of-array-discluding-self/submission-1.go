func productExceptSelf(nums []int) []int {
	//[ 1, 2, 4, 6]  Input
	
	//[  , 1, 2, 8]  Prefix
	//[48,24, 6,  ]  Suffix
	
	//[48,24,12, 8]  Products

    prds := make([]int, len(nums))
    for i := range prds {
        prds[i] = 1
    }

    prefix := 1
    for i := 0; i < len(nums); i++ {
        prds[i] = prefix
        prefix *= nums[i]
    }

    postfix := 1
    for i := len(nums) - 1; i > -1; i-- {
        prds[i] *= postfix
        postfix *= nums[i]
    }

    return prds
}
