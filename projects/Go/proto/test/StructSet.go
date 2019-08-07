// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.3.0.0

package test

import "fmt"
import "strconv"
import "strings"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// Workaround for Go unused imports issue
var _ = fmt.Print
var _ = strconv.FormatInt

// StructSet key
type StructSetKey struct {
}

// Convert StructSet flags key to string
func (k *StructSetKey) String() string {
    var sb strings.Builder
    sb.WriteString("StructSetKey(")
    sb.WriteString(")")
    return sb.String()
}

// StructSet struct
type StructSet struct {
    F1 setF1 `json:"f1"`
    F2 setF2 `json:"f2"`
    F3 setF3 `json:"f3"`
    F4 setF4 `json:"f4"`
}

// Create a new StructSet struct
func NewStructSet() *StructSet {
    return &StructSet{
        F1: make(setF1),
        F2: make(setF2),
        F3: make(setF3),
        F4: make(setF4),
    }
}

// Create a new StructSet struct from the given field values
func NewStructSetFromFieldValues(F1 setF1, F2 setF2, F3 setF3, F4 setF4) *StructSet {
    return &StructSet{F1, F2, F3, F4}
}

// Create a new StructSet struct from JSON
func NewStructSetFromJSON(buffer []byte) (*StructSet, error) {
    result := *NewStructSet()
    err := fbe.Json.Unmarshal(buffer, &result)
    if err != nil {
        return nil, err
    }
    return &result, nil
}

// Struct shallow copy
func (s *StructSet) Copy() *StructSet {
    var result = *s
    return &result
}

// Struct deep clone
func (s *StructSet) Clone() *StructSet {
    // Serialize the struct to the FBE stream
    writer := NewStructSetModel(fbe.NewEmptyBuffer())
    _, _ = writer.Serialize(s)

    // Deserialize the struct from the FBE stream
    reader := NewStructSetModel(writer.Buffer())
    result, _, _ := reader.Deserialize()
    return result
}

// Get the struct key
func (s *StructSet) Key() StructSetKey {
    return StructSetKey{
    }
}

// Convert struct to optional
func (s *StructSet) Optional() *StructSet {
    return s
}

// Get the FBE type
func (s *StructSet) FBEType() int { return 132 }

// Convert struct to string
func (s *StructSet) String() string {
    var sb strings.Builder
    sb.WriteString("StructSet(")
    sb.WriteString("f1=")
    if s.F1 != nil {
        first := true
        sb.WriteString("[" + strconv.FormatInt(int64(s.F1.Size()), 10) + "]{")
        for _, v := range s.F1 {
            if first { sb.WriteString("") } else { sb.WriteString(",") }
            sb.WriteString(strconv.FormatUint(uint64(v), 10))
            first = false
        }
        sb.WriteString("}")
    } else {
        sb.WriteString("f1=[0]{}")
    }
    sb.WriteString(",f2=")
    if s.F2 != nil {
        first := true
        sb.WriteString("[" + strconv.FormatInt(int64(s.F2.Size()), 10) + "]{")
        for _, v := range s.F2 {
            if first { sb.WriteString("") } else { sb.WriteString(",") }
            sb.WriteString(v.String())
            first = false
        }
        sb.WriteString("}")
    } else {
        sb.WriteString(",f2=[0]{}")
    }
    sb.WriteString(",f3=")
    if s.F3 != nil {
        first := true
        sb.WriteString("[" + strconv.FormatInt(int64(s.F3.Size()), 10) + "]{")
        for _, v := range s.F3 {
            if first { sb.WriteString("") } else { sb.WriteString(",") }
            sb.WriteString(v.String())
            first = false
        }
        sb.WriteString("}")
    } else {
        sb.WriteString(",f3=[0]{}")
    }
    sb.WriteString(",f4=")
    if s.F4 != nil {
        first := true
        sb.WriteString("[" + strconv.FormatInt(int64(s.F4.Size()), 10) + "]{")
        for _, v := range s.F4 {
            if first { sb.WriteString("") } else { sb.WriteString(",") }
            sb.WriteString(v.String())
            first = false
        }
        sb.WriteString("}")
    } else {
        sb.WriteString(",f4=[0]{}")
    }
    sb.WriteString(")")
    return sb.String()
}

// Convert struct to JSON
func (s *StructSet) JSON() ([]byte, error) {
    return fbe.Json.Marshal(s)
}

// Set wrapper for the field F1
type setF1 map[byte]byte

// Is the set empty?
func (s setF1) IsEmpty() bool {
    return s.Size() > 0
}

// Get the set size
func (s setF1) Size() int {
    return len(s)
}

// Add the given item to the set
func (s setF1) Add(item byte) {
    s[item] = item
}

// Contains the given item in the set?
func (s setF1) Contains(item byte) bool {
    _, exists := s[item]
return exists
}

// Remove the given item from the set
func (s setF1) Remove(item byte) {
    delete(s, item)
}

// Clear the set
func (s setF1) Clear() {
    for i := range s {
        delete(s, i)
    }
}

// Enumerate the set
func (s setF1) Enumerate() []byte {
    array := make([]byte, 0)
    for _, v := range s {
        array = append(array, v)
    }
    return array
}

// Is the current set a subset of the given set?
func (s setF1) Subset(set setF1) bool {
    result := true
    for _, v := range s {
        if !set.Contains(v) {
            result = false
            break
        }
    }
    return result
}

// Union of the current set and the given set
func (s setF1) Union(set setF1) setF1 {
    result := make(setF1)
    for _, v := range s {
        result.Add(v)
    }
    for _, v := range set {
        result.Add(v)
    }
    return result
}

// Intersection of the current set and the given set
func (s setF1) Intersection(set setF1) setF1 {
    result := make(setF1)
    for _, v := range set {
        if s.Contains(v) {
            result.Add(v)
        }
    }
    return result
}

// Difference between the current set and the given set
func (s setF1) Difference(set setF1) setF1 {
    result := make(setF1)
    for _, v := range set {
        if !s.Contains(v) {
            result.Add(v)
        }
    }
    for _, v := range s {
        if !set.Contains(v) {
            result.Add(v)
        }
    }
    return result
}

// Convert set to JSON
func (s setF1) MarshalJSON() ([]byte, error) {
    array := make([]byte, 0)
    for _, v := range s {
        array = append(array, v)
    }
    return fbe.Json.Marshal(&array)
}

// Convert JSON to set
func (s *setF1) UnmarshalJSON(body []byte) error {
    var array []byte
    err := fbe.Json.Unmarshal(body, &array)
    if err != nil {
        return err
    } else {
        for _, v := range array {
            (*s)[v] = v
        }
    }
    return nil
}

// Set wrapper for the field F2
type setF2 map[EnumSimpleKey]EnumSimple

// Is the set empty?
func (s setF2) IsEmpty() bool {
    return s.Size() > 0
}

// Get the set size
func (s setF2) Size() int {
    return len(s)
}

// Add the given item to the set
func (s setF2) Add(item EnumSimple) {
    s[item.Key()] = item
}

// Contains the given item in the set?
func (s setF2) Contains(item EnumSimple) bool {
    _, exists := s[item.Key()]
return exists
}

// Remove the given item from the set
func (s setF2) Remove(item EnumSimple) {
    delete(s, item.Key())
}

// Clear the set
func (s setF2) Clear() {
    for i := range s {
        delete(s, i)
    }
}

// Enumerate the set
func (s setF2) Enumerate() []EnumSimple {
    array := make([]EnumSimple, 0)
    for _, v := range s {
        array = append(array, v)
    }
    return array
}

// Is the current set a subset of the given set?
func (s setF2) Subset(set setF2) bool {
    result := true
    for _, v := range s {
        if !set.Contains(v) {
            result = false
            break
        }
    }
    return result
}

// Union of the current set and the given set
func (s setF2) Union(set setF2) setF2 {
    result := make(setF2)
    for _, v := range s {
        result.Add(v)
    }
    for _, v := range set {
        result.Add(v)
    }
    return result
}

// Intersection of the current set and the given set
func (s setF2) Intersection(set setF2) setF2 {
    result := make(setF2)
    for _, v := range set {
        if s.Contains(v) {
            result.Add(v)
        }
    }
    return result
}

// Difference between the current set and the given set
func (s setF2) Difference(set setF2) setF2 {
    result := make(setF2)
    for _, v := range set {
        if !s.Contains(v) {
            result.Add(v)
        }
    }
    for _, v := range s {
        if !set.Contains(v) {
            result.Add(v)
        }
    }
    return result
}

// Convert set to JSON
func (s setF2) MarshalJSON() ([]byte, error) {
    array := make([]EnumSimple, 0)
    for _, v := range s {
        array = append(array, v)
    }
    return fbe.Json.Marshal(&array)
}

// Convert JSON to set
func (s *setF2) UnmarshalJSON(body []byte) error {
    var array []EnumSimple
    err := fbe.Json.Unmarshal(body, &array)
    if err != nil {
        return err
    } else {
        for _, v := range array {
            (*s)[v.Key()] = v
        }
    }
    return nil
}

// Set wrapper for the field F3
type setF3 map[FlagsSimpleKey]FlagsSimple

// Is the set empty?
func (s setF3) IsEmpty() bool {
    return s.Size() > 0
}

// Get the set size
func (s setF3) Size() int {
    return len(s)
}

// Add the given item to the set
func (s setF3) Add(item FlagsSimple) {
    s[item.Key()] = item
}

// Contains the given item in the set?
func (s setF3) Contains(item FlagsSimple) bool {
    _, exists := s[item.Key()]
return exists
}

// Remove the given item from the set
func (s setF3) Remove(item FlagsSimple) {
    delete(s, item.Key())
}

// Clear the set
func (s setF3) Clear() {
    for i := range s {
        delete(s, i)
    }
}

// Enumerate the set
func (s setF3) Enumerate() []FlagsSimple {
    array := make([]FlagsSimple, 0)
    for _, v := range s {
        array = append(array, v)
    }
    return array
}

// Is the current set a subset of the given set?
func (s setF3) Subset(set setF3) bool {
    result := true
    for _, v := range s {
        if !set.Contains(v) {
            result = false
            break
        }
    }
    return result
}

// Union of the current set and the given set
func (s setF3) Union(set setF3) setF3 {
    result := make(setF3)
    for _, v := range s {
        result.Add(v)
    }
    for _, v := range set {
        result.Add(v)
    }
    return result
}

// Intersection of the current set and the given set
func (s setF3) Intersection(set setF3) setF3 {
    result := make(setF3)
    for _, v := range set {
        if s.Contains(v) {
            result.Add(v)
        }
    }
    return result
}

// Difference between the current set and the given set
func (s setF3) Difference(set setF3) setF3 {
    result := make(setF3)
    for _, v := range set {
        if !s.Contains(v) {
            result.Add(v)
        }
    }
    for _, v := range s {
        if !set.Contains(v) {
            result.Add(v)
        }
    }
    return result
}

// Convert set to JSON
func (s setF3) MarshalJSON() ([]byte, error) {
    array := make([]FlagsSimple, 0)
    for _, v := range s {
        array = append(array, v)
    }
    return fbe.Json.Marshal(&array)
}

// Convert JSON to set
func (s *setF3) UnmarshalJSON(body []byte) error {
    var array []FlagsSimple
    err := fbe.Json.Unmarshal(body, &array)
    if err != nil {
        return err
    } else {
        for _, v := range array {
            (*s)[v.Key()] = v
        }
    }
    return nil
}

// Set wrapper for the field F4
type setF4 map[StructSimpleKey]StructSimple

// Is the set empty?
func (s setF4) IsEmpty() bool {
    return s.Size() > 0
}

// Get the set size
func (s setF4) Size() int {
    return len(s)
}

// Add the given item to the set
func (s setF4) Add(item StructSimple) {
    s[item.Key()] = item
}

// Contains the given item in the set?
func (s setF4) Contains(item StructSimple) bool {
    _, exists := s[item.Key()]
return exists
}

// Remove the given item from the set
func (s setF4) Remove(item StructSimple) {
    delete(s, item.Key())
}

// Clear the set
func (s setF4) Clear() {
    for i := range s {
        delete(s, i)
    }
}

// Enumerate the set
func (s setF4) Enumerate() []StructSimple {
    array := make([]StructSimple, 0)
    for _, v := range s {
        array = append(array, v)
    }
    return array
}

// Is the current set a subset of the given set?
func (s setF4) Subset(set setF4) bool {
    result := true
    for _, v := range s {
        if !set.Contains(v) {
            result = false
            break
        }
    }
    return result
}

// Union of the current set and the given set
func (s setF4) Union(set setF4) setF4 {
    result := make(setF4)
    for _, v := range s {
        result.Add(v)
    }
    for _, v := range set {
        result.Add(v)
    }
    return result
}

// Intersection of the current set and the given set
func (s setF4) Intersection(set setF4) setF4 {
    result := make(setF4)
    for _, v := range set {
        if s.Contains(v) {
            result.Add(v)
        }
    }
    return result
}

// Difference between the current set and the given set
func (s setF4) Difference(set setF4) setF4 {
    result := make(setF4)
    for _, v := range set {
        if !s.Contains(v) {
            result.Add(v)
        }
    }
    for _, v := range s {
        if !set.Contains(v) {
            result.Add(v)
        }
    }
    return result
}

// Convert set to JSON
func (s setF4) MarshalJSON() ([]byte, error) {
    array := make([]StructSimple, 0)
    for _, v := range s {
        array = append(array, v)
    }
    return fbe.Json.Marshal(&array)
}

// Convert JSON to set
func (s *setF4) UnmarshalJSON(body []byte) error {
    var array []StructSimple
    err := fbe.Json.Unmarshal(body, &array)
    if err != nil {
        return err
    } else {
        for _, v := range array {
            (*s)[v.Key()] = v
        }
    }
    return nil
}
