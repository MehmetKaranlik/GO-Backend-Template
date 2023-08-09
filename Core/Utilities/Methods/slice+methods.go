package Methods

// Append appends given item to the slice
func Append[T any](slice []T, item T) []T {
	return append(slice, item)
}

// RemoveAt removes the item with given index from the slice
func RemoveAt[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

// Contains Checks if the given slice contains the given element returns bool
func Contains[T comparable](slice []T, element T) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

// ContainsIndexed Checks if the given slice contains the given element returns bool and index of that item
func ContainsIndexed[T comparable](slice []T, element T) (int, bool) {
	for i, e := range slice {
		if e == element {
			return i, true
		}
	}
	return -1, false
}

// Map Transforms the given slice with the given function
func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

// RemoveDuplicate Removes the duplicate items from the given slice
func RemoveDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// FirstIndexOf Returns the first index of the given element in the given slice
func FirstIndexOf[T comparable](slice []T, element T) int {
	for i, e := range slice {
		if e == element {
			return i
		}
	}
	return -1
}

// Filter Filters the slice with given function
func Filter[T any](slice *[]T, f func(T) bool) []T {
	var result []T
	for _, v := range *slice {
		if !f(v) {
			result = append(result, v)
		}
	}
	return result
}
