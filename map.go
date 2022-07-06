package mapset

type Map[K comparable, V any] interface {
	// 移除 Map 对象的所有键/值对
	Clear()
	// 返回新的 Map 对象
	Clone() Map[K, V]
	// 返回 hashMap 是否为空
	IsEmpty() bool
	// 设置键值对，返回该 Map 对象。
	Set(K, V) Map[K, V]
	// 返回键对应的值，返回找到的值（未找到返回 V 类型的零值）和 是否存在
	Get(k K) (V, bool)
	// 返回键对应的值，如果不存在，则返回 传入的默认值
	GetOrDefault(k K, defaultValue V) V
	// 返回一个布尔值，用于判断 Map 中是否包含键对应的值。
	Has(K) bool
	// 删除 Map 中的元素，删除成功返回 true，失败返回 false。
	Delete(K)
	// 返回 Map 对象键/值对的数量。
	Size() int
	// 迭代 Map 对象键/值对，返回 false 停止迭代
	Range(fn func(key K, val V) bool)
	// 返回一个新的 数组 对象，包含了 Map 对象中每个元素的键。
	Keys() []K
	// 返回一个新的 数组 对象，包含了 Map 对象中每个元素的值。
	Values() []V
	// 添加键值对到 Map 中
	Merge(x Map[K, V])
}

type mapImpl[K comparable, V any] map[K]V

func NewMap[K comparable, V any]() Map[K, V] {
	return &mapImpl[K, V]{}
}

// 移除 Map 对象的所有键/值对
func (m *mapImpl[K, V]) Clear() {
	*m = mapImpl[K, V]{}
}

// 设置键值对，返回该 Map 对象。
func (m *mapImpl[K, V]) Set(k K, v V) Map[K, V] {
	(*m)[k] = v
	return m
}

func (m mapImpl[K, V]) Clone() Map[K, V] {
	clone := &mapImpl[K, V]{}
	m.Range(func(k K, v V) bool {
		clone.Set(k, v)
		return true
	})
	return clone
}

// 返回 hashMap 是否为空
func (m mapImpl[K, V]) IsEmpty() bool {
	return len(m) == 0
}

// 返回键对应的值，返回找到的值（未找到返回 V 类型的零值）和 是否存在
func (m mapImpl[K, V]) Get(k K) (V, bool) {
	v, has := m[k]
	return v, has
}

// 返回键对应的值，如果不存在，则返回 传入的默认值
func (m mapImpl[K, V]) GetOrDefault(k K, defaultValue V) V {
	v, has := m[k]
	if has {
		return v
	}
	return defaultValue
}

// 返回一个布尔值，用于判断 Map 中是否包含键对应的值。
func (m mapImpl[K, V]) Has(k K) bool {
	_, has := m[k]
	return has
}

// 删除 Map 中的元素，删除成功返回 true，失败返回 false。
func (m mapImpl[K, V]) Delete(k K) {
	delete(m, k)
}

// 返回 Map 对象键/值对的数量。
func (m mapImpl[K, V]) Size() int {
	return len(m)
}

// 迭代 Map 对象键/值对，返回 false 停止迭代
func (m mapImpl[K, V]) Range(fn func(K, V) bool) {
	for k, v := range m {
		if stop := !fn(k, v); stop {
			return
		}
	}
}

// 返回一个新的 数组 对象，包含了 Map 对象中每个元素的键。
func (m mapImpl[K, V]) Keys() []K {
	var keys []K
	m.Range(func(k K, _ V) bool {
		keys = append(keys, k)
		return true
	})
	return keys
}

// 返回一个新的 数组 对象，包含了 Map 对象中每个元素的值。
func (m mapImpl[K, V]) Values() []V {
	var values []V
	m.Range(func(k K, v V) bool {
		values = append(values, v)
		return true
	})
	return values
}

// 添加键值对到 Map 中
func (m mapImpl[K, V]) Merge(x Map[K, V]) {
	x.Range(func(k K, v V) bool {
		m.Set(k, v)
		return true
	})
}
