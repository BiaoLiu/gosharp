package utils

import (
	"sort"
)

type MapSorter struct {
	Keys []string
	Vals []string
}

func NewMapSorter(m map[string]string) *MapSorter {
	ms := &MapSorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]string, 0, len(m)),
	}
	for k, v := range m {
		ms.Keys = append(ms.Keys, k)
		ms.Vals = append(ms.Vals, v)
	}
	return ms
}

func (ms *MapSorter) Sort() {
	sort.Sort(ms)
}

func (ms *MapSorter) Len() int {
	return len(ms.Keys)
}

func (ms *MapSorter) Less(i, j int) bool {
	return ms.Keys[i] < ms.Keys[j]
}

func (ms *MapSorter) Swap(i, j int) {
	ms.Vals[i], ms.Vals[j] = ms.Vals[j], ms.Vals[i]
	ms.Keys[i], ms.Keys[j] = ms.Keys[j], ms.Keys[i]
}

func CastMap(m map[interface{}]interface{}) map[string]interface{} {
	m2 := make(map[string]interface{})
	for key, value := range m {
		switch key := key.(type) {
		case string:
			m2[key] = value
		}
	}
	return m2
}
