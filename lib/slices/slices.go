package slices

// InsertStringSliceCopy inserts items at the given index of a slice.
// Returns a new slice, does not modifies the original slice.
func InsertStringSliceCopy(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

// InsertStringSlice inserts items at the given index of a slice.
// Changes the original slice (and possibly the inserted slice).
func InsertStringSlice(slice, insertion []string, index int) []string {
	return append(slice[:index], append(insertion, slice[index:]...)...)
}

// RemoveStringSliceCopy returns a copy of the slice it is given but with the items
// from the start and end index positions removed
func RemoveStringSliceCopy(slice []string, start, end int) []string {
	result := make([]string, len(slice)-(end-start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}

// RemoveStringSlice returns the original slice it is given but with the items
// from the start and end index positions removed
func RemoveStringSlice(slice []string, start, end int) []string {
	return append(slice[:start], slice[end:]...)
}
