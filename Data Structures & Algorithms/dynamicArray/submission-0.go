type DynamicArray struct {
	array []int
}

func NewDynamicArray(capacity int) *DynamicArray {
	da := DynamicArray{make([]int, 0, capacity)}
	return &da
}

func (da *DynamicArray) Get(i int) int {
	// Assume that index i is valid.
	return da.array[i]
}

func (da *DynamicArray) Set(i int, n int) {
	// Assume that index i is valid.
	da.array[i] = n
}

func (da *DynamicArray) Pushback(n int) {
	if len(da.array) == cap(da.array) {
		da.resize()
	}
	da.array = append(da.array, n)
}

func (da *DynamicArray) Popback() int {
	// Assume that the array is non-empty.
	popped := da.array[len(da.array)-1]
	da.array = da.array[:len(da.array)-1]
	return popped
}

func (da *DynamicArray) resize() {
	resized := make([]int, 0, 2 * cap(da.array))
	for _, n := range da.array {
		resized = append(resized, n)
	}
}

func (da *DynamicArray) GetSize() int {
	return len(da.array)
}

func (da *DynamicArray) GetCapacity() int {
	return cap(da.array)
}

func (da *DynamicArray) findi(n int) int {
	// cannot use binary search since array is not sorted
	for i, elem := range da.array {
		if elem == n {
			return i
		}
	}
	return -1
}
