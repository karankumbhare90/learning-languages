package listops

// IntList is an abstraction of a list of integers which we can define methods on.
type IntList []int

// Foldl applies fn to each element and accumulator, starting from the left.
func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for _, val := range s {
		acc = fn(acc, val)
	}
	return acc
}

// Foldr applies fn to each element and accumulator, starting from the right.
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

// Filter returns a new IntList containing only the elements that satisfy fn.
func (s IntList) Filter(fn func(int) bool) IntList {
	result := IntList{}
	for _, val := range s {
		if fn(val) {
			result = append(result, val)
		}
	}
	return result
}

// Length returns the number of elements in the list.
func (s IntList) Length() int {
	return len(s)
}

// Map applies fn to each element and returns a new IntList with the results.
func (s IntList) Map(fn func(int) int) IntList {
	result := make(IntList, len(s))
	for i, val := range s {
		result[i] = fn(val)
	}
	return result
}

// Reverse returns a new IntList with elements in reverse order.
func (s IntList) Reverse() IntList {
	result := make(IntList, len(s))
	for i, val := range s {
		result[len(s)-1-i] = val
	}
	return result
}

// Append returns a new IntList with lst appended to the end of s.
func (s IntList) Append(lst IntList) IntList {
	result := make(IntList, len(s)+len(lst))
	copy(result, s)
	copy(result[len(s):], lst)
	return result
}

// Concat flattens a list of IntLists into a single IntList.
func (s IntList) Concat(lists []IntList) IntList {
	result := make(IntList, len(s))
	copy(result, s)
	for _, lst := range lists {
		result = append(result, lst...)
	}
	return result
}
