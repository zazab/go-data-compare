package comparator

import "reflect"

func Compare(a, b interface{}, flags int) *Diff {
	var (
		valueA = reflect.TypeOf(a)
		valueB = reflect.TypeOf(b)

		ignoreSliceOrder  = flags == 1
	)

	return compare(valueA, valueB, ignoreSliceOrder)
}