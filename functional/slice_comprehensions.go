package functional

// SliceMap applies a function to all elements of a slice returning a slice  of results
func SliceMap[InputType interface{}, OutputType interface{}](f func(InputType) OutputType, is []InputType) []OutputType {
	rs := make([]OutputType, len(is))
	for i, v := range is {
		rs[i] = f(v)
	}
	return rs
}

// SliceFlatMap flattens a slice of slices of input values, mapping a function across them all returning a simple slice of the results
func SliceFlatMap[InputType interface{}, OutputType interface{}](f func(InputType) OutputType, is [][]InputType) []OutputType {
	rs := make([]OutputType, 0, len(is))
	for _, a := range is {
		for _, v := range a {
			rs = append(rs, f(v))
		}
	}
	return rs
}

func SliceFilter[T interface{}](f func(T) bool, ts []T) []T {
	rs := make([]T, 0, len(ts))
	for _, v := range ts {
		if f(v) {
			rs = append(rs, v)
		}
	}
	return rs
}

// SliceFoldl folds the slice starting from the left most item and an initializer, applying the functor to combine items.
func SliceFoldl[T interface{}](f func(T, T) T, ts []T, init T) T {
	val := init
	for _, v := range ts {
		val = f(val, v)
	}
	return val
}

// SliceFoldr folds the slice starting from the right most item and an initializer, applying the functor to combine items.
func SliceFoldr[T interface{}](f func(T, T) T, ts []T, init T) T {
	val := init
	for i := len(ts) - 1; i >= 0; i-- {
		val = f(val, ts[i])
	}
	return val
}

// SliceReduce reduces the slice into a single instance using a functor to combine items.
func SliceReduce[T interface{}](f func(T, T) T, ts []T) T {
	var val T
	if len(ts) == 0 {
		return val
	}

	if len(ts) == 1 {
		return ts[0]
	}

	init := f(ts[0], ts[1])
	return SliceFoldl(f, ts[2:], init)
}

// SliceToMap transforms a slice into a map using a key derived from the passed in functor
func SliceToMap[TKey comparable, TVal interface{}](f func(TVal) TKey, ts []TVal) map[TKey]TVal {
	returnSet := make(map[TKey]TVal)

	for _, t := range ts {
		returnSet[f(t)] = t
	}

	return returnSet
}

// SliceFirst attempts to get the first entry in a slice to match a functor returning a bool for
// whether any matches were found.
func SliceFirst[T interface{}](f func(T) bool, ts []T) (T, bool) {
	var defaultValue T
	for _, t := range ts {
		if f(t) {
			return t, true
		}
	}
	return defaultValue, false
}
