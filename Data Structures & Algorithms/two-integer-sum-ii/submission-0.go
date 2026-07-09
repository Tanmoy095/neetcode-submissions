func twoSum(numbers []int, target int) []int {
	targets := make([]int, 0, 3)

	// 	numbers = [1,2,4,6]
	//target = 8
	
		poin1 := 0
		poin2 := len(numbers) - 1

		for poin1 < poin2 {
			sum := numbers[poin1] + numbers[poin2]
			if sum == target {
			  targets = append(targets, poin1+1, poin2+1)
				return  targets
			} 
			if sum < target {
				poin1++
			} else {
				poin2--
			}
		}
	return  targets
}
