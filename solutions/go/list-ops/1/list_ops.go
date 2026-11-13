package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for i := range s {
		initial = fn(initial, s[i])
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := range s {
		initial = fn(s[len(s)-1-i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	out := make(IntList, len(s))
	pNewLst := 0

	for i := range s {
		if fn(s[i]) {
			out[pNewLst] = s[i]
			pNewLst++
		}
	}

	return out[:pNewLst]
}

func (s IntList) Length() int {
	n := 0
	for range s {
		n++
	}
	return n
}

func (s IntList) Map(fn func(int) int) IntList {
	out := make(IntList, len(s))
	for i := range s {
		out[i] = fn(s[i])
	}
	return out
}

func (s IntList) Reverse() IntList {
	out := make(IntList, len(s))
	for i := range s {
		out[i] = s[len(s)-1-i]
	}
	return out
}

func (s IntList) Append(lst IntList) IntList {
	out := make(IntList, len(s)+len(lst))

	for i := range s {
		out[i] = s[i]
	}
	for j := range lst {
		out[len(s)+j] = lst[j]
	}

	return out
}

func (s IntList) Concat(lists []IntList) IntList {
	totalLen := len(s)
	for i := range lists {
		totalLen += len(lists[i])
	}

	out := make(IntList, totalLen)

	for i := range s {
		out[i] = s[i]
	}
	n := len(s)
	for j := range lists {
		for k := range lists[j] {
			out[n] = lists[j][k]
			n++
		}
	}

	return out
}
