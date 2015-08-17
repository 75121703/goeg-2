package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := Minimum(4, 3, 8, 2, 9).(int)
	fmt.Printf("%T %v\n", i, i)
	f := Minimum(9.4, -5.4, 3.8, 17.0, -3.1, 0.0).(float64)
	fmt.Printf("%T %v\n", f, f)
	s := Minimum("K", "X", "B", "C", "CC", "CA", "D", "M").(string)
	fmt.Printf("%T %q\n", s, s)

	xs := []int{2, 4, 6, 8}
	ys := []string{"C", "B", "K", "A"}
	fmt.Println("5 @", Index(xs, 5), "  6 @", Index(xs, 6))
	fmt.Println("Z @", Index(ys, "Z"), "  A @", Index(ys, "A"))
	fmt.Println("5 @", IndexReflectX(xs, 5), "  6 @", IndexReflectX(xs, 6))
	fmt.Println("Z @", IndexReflectX(ys, "Z"), "  A @", IndexReflectX(ys, "A"))
	fmt.Println("5 @", IndexReflect(xs, 5), "  6 @", IndexReflect(xs, 6))
	fmt.Println("Z @", IndexReflect(ys, "Z"), "  A @", IndexReflect(ys, "A"))
	fmt.Println("5 @", IntIndexSlicer(xs, 5), "  6 @", IntIndexSlicer(xs, 6))
	fmt.Println("Z @", StringIndexSlicer(ys, "Z"), "  A @", StringIndexSlicer(ys, "A"))

	fmt.Print("5 @ ", SliceIndex(len(xs), func(i int) bool { return xs[i] == 5 }))
	fmt.Println("   6 @", SliceIndex(len(xs), func(i int) bool { return xs[i] == 6 }))
	fmt.Print("Z @ ", SliceIndex(len(ys), func(i int) bool { return ys[i] == "Z" }))
	fmt.Println("   A @", SliceIndex(len(ys), func(i int) bool { return ys[i] == "A" }))
}

// Minimum ...
func Minimum(first interface{}, rest ...interface{}) interface{} {
	minimum := first

	for _, x := range rest {
		switch x := x.(type) {
		case int:
			if x < minimum.(int) {
				minimum = x
			}
		case float64:
			if x < minimum.(float64) {
				minimum = x
			}
		case string:
			if x < minimum.(string) {
				minimum = x
			}
		}
	}

	return minimum
}

// Index ...
func Index(xs interface{}, x interface{}) int {
	switch slice := xs.(type) {
	case []int:
		for i, y := range slice {
			if y == x.(int) {
				return i
			}
		}
	case []string:
		for i, y := range slice {
			if y == x.(string) {
				return i
			}
		}
	}

	return -1
}

// IndexReflectX ...
func IndexReflectX(xs interface{}, x interface{}) int { // Long-winded way
	if slice := reflect.ValueOf(xs); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			switch y := slice.Index(i).Interface().(type) {
			case int:
				if y == x.(int) {
					return i
				}
			case string:
				if y == x.(string) {
					return i
				}
			}
		}
	}

	return -1
}

// IndexReflect ...
func IndexReflect(xs interface{}, x interface{}) int {
	if slice := reflect.ValueOf(xs); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			if reflect.DeepEqual(x, slice.Index(i).Interface()) {
				return i
			}
		}
	}

	return -1
}

// Slicer ...
type Slicer interface {
	EqualTo(i int, x interface{}) bool
	Len() int
}

// IntSlice ...
type IntSlice []int

// EqualTo ...
func (slice IntSlice) EqualTo(i int, x interface{}) bool {
	return slice[i] == x.(int)
}

// Len ...
func (slice IntSlice) Len() int {
	return len(slice)
}

// IntIndexSlicer ...
func IntIndexSlicer(ints []int, x int) int {
	return IndexSlicer(IntSlice(ints), x)
}

// StringSlice ...
type StringSlice []string

// EqualTo ...
func (slice StringSlice) EqualTo(i int, x interface{}) bool {
	return slice[i] == x.(string)
}

// Len ...
func (slice StringSlice) Len() int {
	return len(slice)
}

// StringIndexSlicer ...
func StringIndexSlicer(strs []string, x string) int {
	return IndexSlicer(StringSlice(strs), x)
}

// IndexSlicer ...
func IndexSlicer(slice Slicer, x interface{}) int {
	for i := 0; i < slice.Len(); i++ {
		if slice.EqualTo(i, x) {
			return i
		}
	}
	return -1
}

// SliceIndex ...
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}

	return -1
}
