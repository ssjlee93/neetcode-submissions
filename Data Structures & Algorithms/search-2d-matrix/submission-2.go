func searchMatrix(matrix [][]int, target int) bool {
	// edge : matrix empty
	if len(matrix) == 0 {
		return false
	}

	// edge : matrix only has 1 arr 
	if len(matrix) == 1 {
		// process binary search
		return binarySearch(matrix[0], target)
	}
	
	// case : matrix has many rows
	// treat each head of row as binary search element

	l, r := 0, len(matrix)-1
	targetRow := -1

	for l <= r {
		mid := l + (r-l)/2
		
		if target > matrix[mid][len(matrix[0])-1] {
			l = mid+1
		} else if target < matrix[mid][0] {
			r = mid-1
		} else {
			targetRow = mid
			break;
		}
	}
	
	// edge : target out of range
	if targetRow == -1 {
		return false
	}
	
	return binarySearch(matrix[targetRow], target)
}

// binarySearch
func binarySearch(arr []int, target int) bool {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l+r) / 2
		
		if target > arr[mid] {
			l = mid+1
		} else if target < arr[mid] {
			r = mid-1
		} else {
			return true
		}
	}
	return false
}
