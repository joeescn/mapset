package mapset

import "fmt"

type Set[E comparable] interface {
	// 返回两个 集合 是否相等（内容相同）
	Equal(comp Set[E]) bool
	// 清空集合
	Clear()
	// 向集合添加元素
	Add(e E)
	// 向集合添加多个元素
	Adds(es ...E)
	// 返回元素是否存在
	Has(E) bool
	// 返回元素是否存在，同 Has(E)bool
	Contains(E) bool
	// 返回集合大小
	Size() int
	// 返回 集合 的拷贝
	Clone() Set[E]
	// 删除一个集合内的元素
	Delete(e E)
	// 删除并返回一个集合内的元素，如果不存在则返回 类型零值和false
	Pop() (E, bool)
	// 返回一个包含所有集合元素的数组
	Elements() []E
	// 迭代 Map 对象键/值对，返回 false 停止迭代
	Range(fn func(E) bool)
	// 差集 返回一个集合，元素包含在集合 当前集合 ，但不在集合 传入集合
	Difference(other Set[E]) Set[E]
	// 交集 返回一个集合，元素包含在 当前集合 和 传入集合 中（两个集合都包含）
	Intersection(other Set[E]) Set[E]
	// 确定 当前集合 中的所有元素是否都在 传入集合 中
	IsProperSubset(other Set[E]) bool
	// 确定 传入集合 中的元素都在 当前集合 中
	IsProperSuperset(other Set[E]) bool
	// 确定 当前集合 是否是 传入集合 的子集
	IsSubset(other Set[E]) bool
	// 判断 传入集合 是否是 当前集合 的超集
	IsSuperset(other Set[E]) bool
	// 并集 返回一个集合，元素包含 当前集合 和 传入集合 的所有元素
	Union(other Set[E]) Set[E]
}

type setImpl[E comparable] struct {
	m Map[E, struct{}]
}

func NewSet[E comparable](es ...E) Set[E] {
	s := &setImpl[E]{m: NewMap[E, struct{}]()}
	if len(es) > 0 {
		s.Adds(es...)
	}
	return s
}

func (s *setImpl[E]) Equal(comp Set[E]) bool {
	return s.IsSubset(comp) && s.Size() == comp.Size()
}

func (s *setImpl[E]) Clear() {
	s.m.Clear()
}

// 向集合添加元素
func (s *setImpl[E]) Add(e E) {
	s.m.Set(e, struct{}{})
}

// 向集合添加多个元素
func (s *setImpl[E]) Adds(es ...E) {
	for i := range es {
		s.Add(es[i])
	}
}

// 返回元素是否存在
func (s *setImpl[E]) Has(e E) bool {
	return s.m.Has(e)
}

// 返回集合大小
func (s *setImpl[E]) Size() int {
	return s.m.Size()
}

// 返回 集合 的拷贝
func (s *setImpl[E]) Clone() Set[E] {
	clone := &setImpl[E]{m: NewMap[E, struct{}]()}
	clone.Adds(s.Elements()...)
	return clone
}

// 返回元素是否存在，同 Has(E)bool
func (s *setImpl[E]) Contains(e E) bool {
	return s.Has(e)
}

// 删除一个集合内的元素
func (s *setImpl[E]) Delete(e E) {
	s.m.Delete(e)
}

// 删除并返回一个集合内的元素，如果不存在则返回 类型零值和false
func (s *setImpl[E]) Pop() (E, bool) {
	var (
		elem E
		find bool
	)
	s.m.Range(func(e E, _ struct{}) bool {
		elem = e
		find = true
		s.m.Delete(elem)
		return false
	})
	return elem, find
}

// 返回一个包含所有集合元素的数组
func (s *setImpl[E]) Elements() []E {
	var es []E
	s.Range(func(e E) bool {
		es = append(es, e)
		return true
	})
	return es
}

// 迭代 Map 对象键/值对，返回 false 停止迭代
func (s *setImpl[E]) Range(fn func(E) bool) {
	s.m.Range(func(e E, _ struct{}) bool {
		return fn(e)
	})
}

// 差集 返回一个集合，元素包含在集合 当前集合 ，但不在集合 传入集合
func (s *setImpl[E]) Difference(other Set[E]) Set[E] {
	diff := &setImpl[E]{m: NewMap[E, struct{}]()}
	s.Range(func(e E) bool {
		if !other.Has(e) {
			diff.Add(e)
		}
		return true
	})
	return diff
}

// 差集 返回一个集合，元素包含在集合 当前集合 ，但不在集合 传入集合
func (s *setImpl[E]) DifferenceWithDelete(other Set[E]) Set[E] {
	diff := s.Clone()
	other.Range(func(e E) bool {
		diff.Delete(e)
		return true
	})
	return diff
}

// 交集 返回一个集合，元素包含在 当前集合 和 传入集合 中（两个集合都包含）
func (s *setImpl[E]) Intersection(other Set[E]) Set[E] {
	inter := &setImpl[E]{m: NewMap[E, struct{}]()}
	var big, small Set[E] = s, other
	if big.Size() < small.Size() {
		big, small = small, big
	}
	small.Range(func(e E) bool {
		if big.Has(e) {
			inter.Add(e)
		}
		return true
	})
	return inter
}

// 确定 当前集合 中的所有元素是否都在 传入集合 中
func (s *setImpl[E]) IsProperSubset(other Set[E]) bool {
	return s.Size() <= other.Size() && s.Difference(other).Size() == 0
}

// 确定 传入集合 中的元素都在 当前集合 中
func (s *setImpl[E]) IsProperSuperset(other Set[E]) bool {
	return other.IsProperSubset(s)
}

// 确定 当前集合 是否是 传入集合 的子集
func (s *setImpl[E]) IsSubset(other Set[E]) bool {
	return s.Size() <= other.Size() && s.Difference(other).Size() == 0
}

// 判断 传入集合 是否是 当前集合 的超集
func (s *setImpl[E]) IsSuperset(other Set[E]) bool {
	return other.IsSubset(s)
}

// 并集 返回一个集合，元素包含 当前集合 和 传入集合 的所有元素
func (s *setImpl[E]) Union(other Set[E]) Set[E] {
	union := s.Clone()
	other.Range(func(e E) bool {
		union.Add(e)
		return true
	})
	return union
}

func (s *setImpl[E]) String() string {
	return fmt.Sprint(s.Elements())
}
