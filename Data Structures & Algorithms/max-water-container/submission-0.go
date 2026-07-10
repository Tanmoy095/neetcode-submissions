func maxArea(heights []int) int {
	// Stores the maximum water found so far.
	maxContainer := 0

	// Left pointer starts from the beginning.
	leftP := 0

	// Right pointer starts from the end.
	rightP := len(heights) - 1

	// Continue until both pointers meet.
	for leftP < rightP {

		// ---------------------------------------------------
		// Height of the container
		// ---------------------------------------------------
		// Water level is limited by the shorter wall.
		height := min(heights[leftP], heights[rightP])

		// ---------------------------------------------------
		// Width of the container
		// ---------------------------------------------------
		width := rightP - leftP

		// ---------------------------------------------------
		// Current container area
		// ---------------------------------------------------
		currentContainer := height * width

		// Update maximum area if current one is larger.
		maxContainer = max(maxContainer, currentContainer)

		// ---------------------------------------------------
		// Move the pointer with the smaller height.
		// ---------------------------------------------------
		//
		// Why?
		//
		// Area = min(heightLeft, heightRight) * width
		//
		// Width always decreases when we move pointers.
		// Therefore, the only hope of getting a larger area
		// is finding a taller shorter wall.
		//
		// Moving the taller wall cannot increase the height,
		// because the shorter wall still limits the water.
		//
		if heights[leftP] < heights[rightP] {
			// Left wall is shorter.
			// Try to find a taller left wall.
			leftP++
		} else {
			// Right wall is shorter (or equal).
			// Try to find a taller right wall.
			rightP--
		}
	}

	return maxContainer
}