package main

// LargestRectangleArea returns the area of the largest rectangle in the histogram.
func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	// Stack will store indices of bars.
	// We'll maintain it as a monotonic increasing stack by height.
	stack := make([]int, 0, n)

	maxArea := 0

	for i := 0; i <= n; i++ {
		// Treat i == n as a "virtual" bar of height 0 to flush the stack.
		var currHeight int
		if i == n {
			currHeight = 0
		} else {
			currHeight = heights[i]
		}

		// While current bar breaks the increasing sequence, pop and compute area
		for len(stack) > 0 && currHeight < heights[stack[len(stack)-1]] {
			// The bar at top is the height of the rectangle
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			// Determine width:
			// right boundary is i-1 (current i is first smaller to the right),
			// left boundary depends on new top of stack.
			var width int
			if len(stack) == 0 {
				// No smaller to the left -> rectangle spans [0..i-1]
				width = i
			} else {
				// Smaller to the left at stack top -> spans (top..i-1)
				leftIndex := stack[len(stack)-1]
				width = i - leftIndex - 1
			}

			area := h * width
			if area > maxArea {
				maxArea = area
			}
		}

		// Push current index to stack (maintains increasing heights)
		stack = append(stack, i)
	}

	return maxArea
}