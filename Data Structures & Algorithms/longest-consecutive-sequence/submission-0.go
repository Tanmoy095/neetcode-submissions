func longestConsecutive(nums []int) int {
  	//suppose nums = [2, 20, 4, 10, 3, 4, 5]
	numberExists := make (map [int]bool)
	//Put all numbers into a map for O(1) lookups
	for _,num := range nums {
		numberExists[num] = true
	}
	//now map become {2: true, 20: true, 4: true, 10: true, 3: true, 5: true}

	//If num - 1 does not exist in our map, then num must be the start of a sequence!
	//If num - 1 does exist, we ignore it and move on, because we will catch this entire sequence later when we start from the true beginning.

	longestStreak:=0

	for _,num := range nums {

		if !numberExists[num -1]{ // not exist ,so starter 

			currentNum := num
            currentStreak := 1

		
		for numberExists[currentNum+1]{
			currentNum++
            currentStreak++
		}

		// Update longest streak.
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
		
	}
	
	return longestStreak

}

