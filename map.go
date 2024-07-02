// Copyright © 2024 Timothy E. Peoples

package stringset

type Map map[string]Set

func (m Map) Add(key string, set Set) Map {
	m[key] = set
	return m
}

func (m Map) Insert(key, value string) {
	if m[key] == nil {
		m[key] = New(value)
	} else {
		m[key].Add(value)
	}
}

func (m Map) Expunge(value string) {
	for _, ss := range m {
		delete(ss, value)
	}
}

func (m Map) Equal(o Map) bool {
	if m == nil && o == nil {
		return true
	}

	if m == nil || o == nil {
		return false
	}

	if len(m) != len(o) {
		return false
	}

	for mkey, mset := range m {
		if oset, ok := o[mkey]; !ok || !oset.Equal(mset) {
			return false
		}
	}

	for okey := range o {
		if _, ok := m[okey]; !ok {
			return false
		}
	}

	return true
}

func (m Map) Keys() Set {
	keys := make([]string, len(m))
	var i int
	for k := range m {
		keys[i] = k
		i++
	}

	return New(keys...)
}
