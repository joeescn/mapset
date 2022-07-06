package mapset_test

import (
	"sort"
	"testing"

	"github.com/joeescn/mapset"
	"github.com/stretchr/testify/assert"
)

func TestSet_Clear(t *testing.T) {
	s := mapset.NewSet(1, 2, 3)
	assert.Equal(t, 3, s.Size())
	s.Clear()
	assert.Equal(t, 0, s.Size())
}

func TestSet_Add(t *testing.T) {
	s := mapset.NewSet[int]()
	s.Add(1)
	s.Add(2)
	elems := s.Elements()
	sort.IntSlice(elems).Sort()
	assert.Equal(t, []int{1, 2}, elems)
}

func TestSet_Adds(t *testing.T) {
	s := mapset.NewSet[int]()
	s.Adds(1, 2)
	elems := s.Elements()
	sort.IntSlice(elems).Sort()
	assert.Equal(t, []int{1, 2}, elems)
}

func TestSet_Has(t *testing.T) {
	s := mapset.NewSet[int]()
	s.Adds(1)
	assert.Equal(t, true, s.Has(1))
	assert.Equal(t, false, s.Has(2))
}

func TestSet_Contains(t *testing.T) {
	s := mapset.NewSet[int]()
	s.Adds(1)
	assert.Equal(t, true, s.Contains(1))
	assert.Equal(t, false, s.Contains(2))
}

func TestSet_Size(t *testing.T) {
	s := mapset.NewSet[int]()
	s.Add(1)
	assert.Equal(t, 1, s.Size())
	s.Add(2)
	assert.Equal(t, 2, s.Size())
	s.Add(2)
	assert.Equal(t, 2, s.Size())
}

func TestSet_Clone(t *testing.T) {
	s := mapset.NewSet[int]()
	s.Add(1)
	clone := s.Clone()
	assert.Equal(t, true, s.IsSubset(clone) && s.Size() == clone.Size())
	clone.Add(2)
	assert.Equal(t, false, s.IsSubset(clone) && s.Size() == clone.Size())
}

func TestSet_Delete(t *testing.T) {
	s := mapset.NewSet[int]()
	s.Add(1)
	assert.Equal(t, true, s.Has(1))
	s.Delete(1)
	assert.Equal(t, false, s.Has(1))
}

func TestSet_Pop(t *testing.T) {
	s := mapset.NewSet[int]()
	s.Add(1)
	val, find := s.Pop()
	assert.Equal(t, true, find, "集合不为空时，应该返回 true")
	assert.Equal(t, 1, val, "抛出的元素不是期望值")
	s.Add(0)
	val, find = s.Pop()
	assert.Equal(t, true, find, "集合不为空时，应该返回 true")
	assert.Equal(t, 0, val, "抛出的元素不是期望值")
	val, find = s.Pop()
	var zero int
	assert.Equal(t, false, find, "集合为空时，应该返回 false")
	assert.Equal(t, zero, val, "没有元素是应该返回类型的默认值")
}

func TestSet_ElementsAndRangeAndEqual(t *testing.T) {
	s := mapset.NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	comp := mapset.NewSet[int]()
	s.Range(func(i int) bool {
		// 添加两个元素
		comp.Add(i)
		return comp.Size() < 2
	})
	assert.Equal(t, false, s.Equal(comp))
	comp.Adds(s.Elements()...)
	assert.Equal(t, true, s.Equal(comp))
}

func TestSet_Difference(t *testing.T) {
	a := mapset.NewSet(1, 2, 3)
	b := mapset.NewSet[int]()
	b.Adds(2, 3)
	b.Add(4)
	difference := a.Difference(b).Elements()
	sort.IntSlice(difference).Sort()
	assert.Equal(t, []int{1}, difference)
}
func TestSet_Intersection(t *testing.T) {
	a := mapset.NewSet(1, 2, 3)
	b := mapset.NewSet[int]()
	b.Adds(2, 3)
	b.Add(4)
	intersection := a.Intersection(b).Elements()
	sort.IntSlice(intersection).Sort()
	assert.Equal(t, []int{2, 3}, intersection)
}

func TestSet_IsProperSubset(t *testing.T) {
	a := mapset.NewSet(1, 2, 3)
	b := mapset.NewSet[int]()
	b.Adds(2, 3)
	b.Add(4)
	assert.Equal(t, false, a.IsProperSubset(b))
	b.Add(1)
	assert.Equal(t, true, a.IsProperSubset(b))
}

func TestSet_IsProperSuperset(t *testing.T) {
	a := mapset.NewSet(1, 2, 3)
	b := mapset.NewSet[int]()
	b.Adds(2, 3)
	b.Add(4)
	assert.Equal(t, false, a.IsProperSuperset(b))
	b.Delete(4)
	assert.Equal(t, true, a.IsProperSuperset(b))
}

func TestSet_IsSubset(t *testing.T) {
	a := mapset.NewSet(1, 2, 3)
	b := mapset.NewSet[int]()
	b.Adds(2, 3)
	b.Add(4)
	assert.Equal(t, false, a.IsSubset(b))
	a.Delete(1)
	assert.Equal(t, true, a.IsSubset(b))
}

func TestSet_IsSuperset(t *testing.T) {
	a := mapset.NewSet(1, 2, 3)
	b := mapset.NewSet[int]()
	b.Adds(2, 3)
	b.Add(4)
	assert.Equal(t, false, a.IsSuperset(b))
	a.Add(4)
	assert.Equal(t, true, a.IsSuperset(b))
}

func TestSet_Union(t *testing.T) {
	a := mapset.NewSet(1, 2, 3)
	b := mapset.NewSet[int]()
	b.Adds(2, 3)
	b.Add(4)
	union := a.Union(b).Elements()
	sort.IntSlice(union).Sort()
	assert.Equal(t, []int{1, 2, 3, 4}, union)
}
