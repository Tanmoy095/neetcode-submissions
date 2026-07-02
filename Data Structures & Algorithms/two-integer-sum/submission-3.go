func twoSum(nums []int, target int) []int {

		//with Hashmap

		seen := make(map[int]int)

	for currIdx, currVal := range nums {
		complement := target - currVal
		//Check if this complement is already in  map
		if savedIdx, found := seen[complement]; found {
			//found it
			return []int{savedIdx, currIdx}
		}

		//If not found log or save the current idx or val in map

		seen[currVal] = currIdx

	}
	return []int{}
}





	