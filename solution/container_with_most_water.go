package solution

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		width := right - left
		minHeight := 0
		if height[left] < height[right] {
			minHeight = height[left]
		} else {
			minHeight = height[right]
		}
		area := width * minHeight

		if area > maxArea {
			maxArea = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}
