package sort

import (

)

type DataRange []int64

func (a DataRange) Len() int           { return len(a) }
func (a DataRange) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a DataRange) Less(i, j int) bool { return a[i] < a[j] }