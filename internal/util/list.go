package util

import "sort"

type ReduceList struct {
	Values      []string
	FoundValues map[string]bool
}

func (r *ReduceList) Init() {
	r.Values = []string{}
	r.FoundValues = make(map[string]bool)
}

func (r *ReduceList) Add(t string) {
	if !r.FoundValues[t] {
		r.FoundValues[t] = true
		r.Values = append(r.Values, t)
	}
}
func (r *ReduceList) Sorted() []string {
	sort.Strings(r.Values)
	return r.Values
}
