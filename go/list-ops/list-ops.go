package listops

// IntList define a type for an array of ints
type IntList []int

type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// Length of a IntList type
func (il IntList) Length() int {

	len := 0
	emptyList := 1

	for k := range il {
		len = k
		emptyList = 0
	}

	if emptyList == 0 {
		len++
	}

	return len
}

// Reverse the order of a IntList
func (il IntList) Reverse() IntList {
	var result []int

	if il.Length() == 0 {
		return IntList{}
	}

	for i := il.Length() - 1; i >= 0; i-- {
		result = append(result, il[i])
	}
	return IntList(result)
}

// Append a given IntList in the end of a IntList
func (il IntList) Append(list IntList) IntList {
	for _, v := range list {
		il = append(il, v)
	}
	return il
}

// Concat an array of IntList into another one
func (il IntList) Concat(lists []IntList) IntList {
	for _, vl := range lists {
		for _, v := range vl {
			il = append(il, v)
		}
	}

	return il
}

// Filter InstList elements using a given function fn
func (il IntList) Filter(fn predFunc) IntList {
	var result []int

	if il.Length() == 0 {
		return IntList{}
	}

	for _, v := range il {
		if fn(v) {
			result = append(result, v)
		}
	}

	return IntList(result)
}

// Map call the given function for each IntList element
func (il IntList) Map(fn unaryFunc) IntList {
	if il.Length() == 0 {
		return IntList{}
	}

	for k := range il {
		il[k] = fn(il[k])
	}

	return il
}

// Foldl call the given function reading the IntList elements
// from right(end) to left(begin)
func (il IntList) Foldl(fn binFunc, n int) int {
	if il.Length() == 0 {
		return n
	}

	result := n
	for i := 0; i < il.Length(); i++ {
		result = fn(result, il[i])
	}

	return result
}

// Foldr call the given function reading the IntList elements
// from left(begin) to right(end)
func (il IntList) Foldr(fn binFunc, n int) int {
	if il.Length() == 0 {
		return n
	}

	result := n
	for i := il.Length() - 1; i >= 0; i-- {
		result = fn(il[i], result)
	}

	return result
}
