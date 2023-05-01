package funciones

func FindMaxAndMin(arr []int) (int, int) {
	arrLen := len(arr)
	if arrLen == 1 {
		return arr[0], arr[0]
	}
	if arrLen == 2 {
		if arr[0] < arr[1] {
			return arr[1], arr[0]
		}
		return arr[0], arr[1]
	}

	mid := arrLen / 2
	max1, min1 := FindMaxAndMin(arr[:mid])
	max2, min2 := FindMaxAndMin(arr[mid:])
	maxVal := max1

	if max2 > maxVal {
		maxVal = max2
	}

	minVal := min1
	if min2 < minVal {
		minVal = min2
	}

	return maxVal, minVal
}
