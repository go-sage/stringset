

# stringset
`import "github.com/go-sage/stringset"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package stringset provides the Set type as a container for a set of strings.


## <a name="pkg-index">Index</a>
* [type Map](#Map)
  * [func (m Map) Add(key string, set Set) Map](#Map.Add)
  * [func (m Map) Equal(o Map) bool](#Map.Equal)
  * [func (m Map) Expunge(value string)](#Map.Expunge)
  * [func (m Map) Insert(key, value string)](#Map.Insert)
* [type Set](#Set)
  * [func New(strings ...string) Set](#New)
  * [func (ss Set) Add(strings ...string)](#Set.Add)
  * [func (ss Set) Clear()](#Set.Clear)
  * [func (ss Set) Clone() Set](#Set.Clone)
  * [func (ss Set) Empty() bool](#Set.Empty)
  * [func (ss Set) Equal(oss Set) bool](#Set.Equal)
  * [func (ss Set) Has(s string) bool](#Set.Has)
  * [func (ss Set) Intersection(oss Set) Set](#Set.Intersection)
  * [func (ss Set) IsSubsetOf(oss Set) bool](#Set.IsSubsetOf)
  * [func (ss Set) IsSupersetOf(oss Set) bool](#Set.IsSupersetOf)
  * [func (ss Set) Minus(oss Set) Set](#Set.Minus)
  * [func (ss Set) Ordered() []string](#Set.Ordered)
  * [func (ss Set) Remove(strings ...string)](#Set.Remove)
  * [func (ss Set) Strings() []string](#Set.Strings)
  * [func (ss Set) Union(oss Set) Set](#Set.Union)
  * [func (ss Set) Update(oss Set)](#Set.Update)


#### <a name="pkg-files">Package files</a>
[map.go](/src/target/map.go) [stringset.go](/src/target/stringset.go) 






## <a name="Map">type</a> [Map](/src/target/map.go?s=19:42#L1)
``` go
type Map map[string]Set
```









### <a name="Map.Add">func</a> (Map) [Add](/src/target/map.go?s=44:85#L1)
``` go
func (m Map) Add(key string, set Set) Map
```



### <a name="Map.Equal">func</a> (Map) [Equal](/src/target/map.go?s=321:351#L14)
``` go
func (m Map) Equal(o Map) bool
```



### <a name="Map.Expunge">func</a> (Map) [Expunge](/src/target/map.go?s=234:268#L8)
``` go
func (m Map) Expunge(value string)
```



### <a name="Map.Insert">func</a> (Map) [Insert](/src/target/map.go?s=115:153#L1)
``` go
func (m Map) Insert(key, value string)
```



## <a name="Set">type</a> [Set](/src/target/stringset.go?s=244:268#L1)
``` go
type Set map[string]bool
```
Set is a container for a set of strings providing most commonly used set
operations (e.g. union, intersection, etc...)







### <a name="New">func</a> [New](/src/target/stringset.go?s=338:369#L3)
``` go
func New(strings ...string) Set
```
New returns a new Set initialized with the given list of string.





### <a name="Set.Add">func</a> (Set) [Add](/src/target/stringset.go?s=495:531#L10)
``` go
func (ss Set) Add(strings ...string)
```
Add incorporates each of the given strings into the receiver Set.




### <a name="Set.Clear">func</a> (Set) [Clear](/src/target/stringset.go?s=933:954#L31)
``` go
func (ss Set) Clear()
```
Clear removes all elements from the receiver Set.




### <a name="Set.Clone">func</a> (Set) [Clone](/src/target/stringset.go?s=1233:1258#L47)
``` go
func (ss Set) Clone() Set
```
Clone returns a copy of the receiver Set.




### <a name="Set.Empty">func</a> (Set) [Empty](/src/target/stringset.go?s=1093:1119#L39)
``` go
func (ss Set) Empty() bool
```
Empty returns true if the receiver Set is nil or contains no strings,
otherwise false.




### <a name="Set.Equal">func</a> (Set) [Equal](/src/target/stringset.go?s=1779:1812#L80)
``` go
func (ss Set) Equal(oss Set) bool
```
Equal returns true if both oss and the receiver Set contain the same
members.




### <a name="Set.Has">func</a> (Set) [Has](/src/target/stringset.go?s=1409:1441#L56)
``` go
func (ss Set) Has(s string) bool
```
Has returns true if s is a member of the receiver Set, false otherwise.




### <a name="Set.Intersection">func</a> (Set) [Intersection](/src/target/stringset.go?s=2341:2380#L116)
``` go
func (ss Set) Intersection(oss Set) Set
```
Intersection returns a new Set containing members that are common to both
oss and the receiver Set.




### <a name="Set.IsSubsetOf">func</a> (Set) [IsSubsetOf](/src/target/stringset.go?s=1461:1499#L60)
``` go
func (ss Set) IsSubsetOf(oss Set) bool
```



### <a name="Set.IsSupersetOf">func</a> (Set) [IsSupersetOf](/src/target/stringset.go?s=1577:1617#L69)
``` go
func (ss Set) IsSupersetOf(oss Set) bool
```



### <a name="Set.Minus">func</a> (Set) [Minus](/src/target/stringset.go?s=2567:2599#L128)
``` go
func (ss Set) Minus(oss Set) Set
```
Minus returns a new Set containing members from the receiver Set that are
not in oss.




### <a name="Set.Ordered">func</a> (Set) [Ordered](/src/target/stringset.go?s=2997:3029#L151)
``` go
func (ss Set) Ordered() []string
```
Ordered returns the members of the receiver Set sorted lexically ascending.




### <a name="Set.Remove">func</a> (Set) [Remove](/src/target/stringset.go?s=652:691#L17)
``` go
func (ss Set) Remove(strings ...string)
```
Remove discards each of the given strings from the receiver Set.




### <a name="Set.Strings">func</a> (Set) [Strings](/src/target/stringset.go?s=2769:2801#L139)
``` go
func (ss Set) Strings() []string
```
Strings returns the members of the receiver Set in an undefined order.




### <a name="Set.Union">func</a> (Set) [Union](/src/target/stringset.go?s=2132:2164#L105)
``` go
func (ss Set) Union(oss Set) Set
```
Union returns a new Set containing members from both oss and the receiver Set.




### <a name="Set.Update">func</a> (Set) [Update](/src/target/stringset.go?s=805:834#L24)
``` go
func (ss Set) Update(oss Set)
```
Update adds each element from oss into the receiver Set.


