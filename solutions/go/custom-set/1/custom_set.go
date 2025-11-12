package stringset

import (
	"fmt"
	"sort"
	"strings"
)

// Set is a collection of unique string values.
type Set map[string]struct{}

// New creates an empty Set.
func New() Set {
	return make(Set)
}

// NewFromSlice creates a Set from a slice of strings.
func NewFromSlice(l []string) Set {
	s := New()
	for _, v := range l {
		s[v] = struct{}{}
	}
	return s
}

// String returns a string representation of the Set in Go syntax style.
// Example: {"a", "b"} or {} for an empty set.
func (s Set) String() string {
	if len(s) == 0 {
		return "{}"
	}
	elems := make([]string, 0, len(s))
	for k := range s {
		elems = append(elems, fmt.Sprintf("%q", k)) // Go-safe quoting
	}
	sort.Strings(elems)
	return fmt.Sprintf("{%s}", strings.Join(elems, ", "))
}

// IsEmpty returns true if the set has no elements.
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has returns true if the element exists in the set.
func (s Set) Has(elem string) bool {
	_, exists := s[elem]
	return exists
}

// Add inserts an element into the set.
func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

// Subset returns true if s1 is a subset of s2.
func Subset(s1, s2 Set) bool {
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

// Disjoint returns true if s1 and s2 have no elements in common.
func Disjoint(s1, s2 Set) bool {
	for k := range s1 {
		if s2.Has(k) {
			return false
		}
	}
	return true
}

// Equal returns true if s1 and s2 contain exactly the same elements.
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

// Intersection returns a new set containing elements common to both s1 and s2.
func Intersection(s1, s2 Set) Set {
	res := New()
	for k := range s1 {
		if s2.Has(k) {
			res.Add(k)
		}
	}
	return res
}

// Difference returns a new set containing elements in s1 but not in s2.
func Difference(s1, s2 Set) Set {
	res := New()
	for k := range s1 {
		if !s2.Has(k) {
			res.Add(k)
		}
	}
	return res
}

// Union returns a new set containing all elements from both s1 and s2.
func Union(s1, s2 Set) Set {
	res := New()
	for k := range s1 {
		res.Add(k)
	}
	for k := range s2 {
		res.Add(k)
	}
	return res
}
