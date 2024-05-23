// Copyright © 2024 Timothy E. Peoples
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.

// Package stringset provides the Set type as a container for a set of strings.
package stringset // import "github.com/go-sage/stringset"

import (
	"sort"
)

// Set is a container for a set of strings providing most commonly used set
// operations (e.g. union, intersection, etc...)
type Set map[string]bool

// New returns a new Set initialized with the given list of string.
func New(strings ...string) Set {
	nss := make(Set)
	nss.Add(strings...)
	return nss
}

// Add incorporates each of the given strings into the receiver Set.
func (ss Set) Add(strings ...string) {
	for _, s := range strings {
		ss[s] = true
	}
}

// Remove discards each of the given strings from the receiver Set.
func (ss Set) Remove(strings ...string) {
	for _, s := range strings {
		delete(ss, s)
	}
}

// Update adds each element from oss into the receiver Set.
func (ss Set) Update(oss Set) {
	for s := range oss {
		ss[s] = true
	}
}

// Clear removes all elements from the receiver Set.
func (ss Set) Clear() {
	for s := range ss {
		delete(ss, s)
	}
}

// Empty returns true if the receiver Set is nil or contains no strings,
// otherwise false.
func (ss Set) Empty() bool {
	if ss == nil || len(ss) == 0 {
		return true
	}
	return false
}

// Clone returns a copy of the receiver Set.
func (ss Set) Clone() Set {
	nss := make(Set)
	for s := range ss {
		nss[s] = true
	}
	return nss
}

// Has returns true if s is a member of the receiver Set, false otherwise.
func (ss Set) Has(s string) bool {
	return ss[s]
}

func (ss Set) IsSubsetOf(oss Set) bool {
	for s := range ss {
		if !oss[s] {
			return false
		}
	}
	return true
}

func (ss Set) IsSupersetOf(oss Set) bool {
	for s := range oss {
		if !ss[s] {
			return false
		}
	}
	return true
}

// Equal returns true if both oss and the receiver Set contain the same
// members.
func (ss Set) Equal(oss Set) bool {
	if ss == nil && oss == nil {
		return true
	}

	if ss == nil || oss == nil {
		return false
	}

	for s := range ss {
		if !oss[s] {
			return false
		}
	}

	for os := range oss {
		if !ss[os] {
			return false
		}
	}

	return true
}

// Union returns a new Set containing members from both oss and the receiver Set.
func (ss Set) Union(oss Set) Set {
	nss := make(Set)

	nss.Update(ss)
	nss.Update(oss)

	return nss
}

// Intersection returns a new Set containing members that are common to both
// oss and the receiver Set.
func (ss Set) Intersection(oss Set) Set {
	nss := make(Set)
	for s := range ss {
		if oss[s] {
			nss[s] = true
		}
	}
	return nss
}

// Minus returns a new Set containing members from the receiver Set that are
// not in oss.
func (ss Set) Minus(oss Set) Set {
	nss := make(Set)
	for s := range ss {
		if !oss[s] {
			nss[s] = true
		}
	}
	return nss
}

// Strings returns the members of the receiver Set in an undefined order.
func (ss Set) Strings() []string {
	strings := make([]string, len(ss))
	var i int
	for k := range ss {
		strings[i] = k
		i++
	}

	return strings
}

// Ordered returns the members of the receiver Set sorted lexically ascending.
func (ss Set) Ordered() []string {
	strings := ss.Strings()
	sort.Strings(strings)
	return strings
}
