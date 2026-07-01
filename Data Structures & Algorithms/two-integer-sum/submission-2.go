
// Define a struct to bind the number's value to its original index position
type NumWithIndex struct {
	Value int
	Index int
}

func twoSum(nums []int, target int) []int {
	// STEP 1: Create a copy holding original indices
	indexedNums := make([]NumWithIndex, len(nums))
	for i, v := range nums {
		indexedNums[i] = NumWithIndex{Value: v, Index: i}
	}

	// STEP 2: Sort based on the Value (Handles negative numbers perfectly)
	sort.Slice(indexedNums, func(i, j int) bool {
		return indexedNums[i].Value < indexedNums[j].Value
	})

	// STEP 3: Initialize pointers at the absolute beginning and end
	left := 0
	right := len(indexedNums) - 1

	// STEP 4: Start the Target Loop
	for left < right {
		currentSum := indexedNums[left].Value + indexedNums[right].Value

		// Found the target!
		if currentSum == target {
			idx1 := indexedNums[left].Index
			idx2 := indexedNums[right].Index

			// TRICK TO FIX: Always return the smaller original index first
			if idx1 > idx2 {
				return []int{idx2, idx1}
			}
			return []int{idx1, idx2}
		}

		// Adjust pointers based on the sum
		if currentSum < target {
			left++ // Sum is too small, move to a larger number
		} else {
			right-- // Sum is too big, move to a smaller number
		}
	}

	return []int{}
}