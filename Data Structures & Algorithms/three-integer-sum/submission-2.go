import "slices"
func threeSum(nums []int) [][]int {
	n := len(nums)
	
	// MISTAKE WATCH: Start with length 0. 
	// We do NOT use a map/set anymore because our two optimizations handle all duplicates!
	answer := make([][]int, 0)
	
	// CRITICAL STEP: You must sort the array first.
	// Sorting brings all duplicate "twin" numbers next to each other.
	slices.Sort(nums)

	for i := 0; i < n-2; i++ {
		
		// -------------------------------------------------------------------------
		// OPTIMIZATION 1: Skip twin anchor values for 'i'
		// -------------------------------------------------------------------------
		// MISTAKE WATCH:  must check 'i > 0' first so you don't look out of bounds.
		// We look BACKWARD (nums[i] == nums[i-1]). If true, we skip this turn.
		if i > 0 && nums[i] == nums[i-1] {
			continue // 'continue' means: skip this one turn and move to the next 'i'
		}

		// -------------------------------------------------------------------------
		// EARLY STOP CHECK: Mathematical optimization
		// -------------------------------------------------------------------------
		// Since the array is sorted, if our smallest current number is greater than 0,
		// adding it to two even larger numbers can NEVER equal 0. 
		if nums[i] > 0 {
			break // 'break' means: stop everything, destroy the loop, we are done!
		}

		// Set up  two pointers for the remaining window
		j := i + 1  // Left pointer (starts right after anchor)
		k := n - 1  // Right pointer (starts at the absolute end)

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum < 0 {
				// Sum is too small! We need a bigger value.
				// Moving 'j' to the right gives us a larger number.
				j++
			} else if sum > 0 {
				// Sum is too big! We need a smaller value.
				// Moving 'k' to the left gives us a smaller number.
				k--
			} else {
				// Perfect Match Found! Store this unique triplet combination.
				answer = append(answer, []int{nums[i], nums[j], nums[k]})

				// -------------------------------------------------------------------------
				// OPTIMIZATION 2: Skip twin values for 'j' and 'k'
				// -------------------------------------------------------------------------
				// MISTAKE WATCH: Why a 'for' loop and not an 'if'? 
				// An 'if' only skips once. If there are 3 or 4 twins in a row (like 1, 1, 1), 
				// a 'for' loop acts like a bulldozer to push past ALL of them.
				// This does NOT make the code O(n³); it just slides the same pointer forward.
				
				// Bulldoze 'j' past identical numbers to its right
				for j < k && nums[j] == nums[j+1] {
					j++
				}

				// Bulldoze 'k' past identical numbers to its left
				for j < k && nums[k] == nums[k-1] {
					k--
				}

				// MISTAKE WATCH: After the loops stop, 'j' and 'k' are still sitting on the 
				// LAST twin they evaluated. We must take one more step onto a BRAND NEW number.
				j++
				k--
			}
		}
	}

	return answer
}