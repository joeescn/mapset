package mapset_test

import (
	"testing"

	"github.com/joeescn/mapset"
	"github.com/stretchr/testify/assert"
)

func TestMap_Clear(t *testing.T) {
	m := mapset.NewMap[string, any]()
	m.Set("key", "value")
	m.Clear()
	assert.Equal(t, 0, m.Size())
}

func TestMap_Clone(t *testing.T) {
	m := mapset.NewMap[string, any]()
	m.Set("key", "value")

	clone := m.Clone()
	clone.Set("key1", "value1")

	assert.Equal(t, false, m.Size() == clone.Size())
}

func TestMap_IsEmpty(t *testing.T) {
	m := mapset.NewMap[string, any]()
	assert.Equal(t, true, m.IsEmpty())

	m.Set("key", "value")
	assert.Equal(t, false, m.IsEmpty())
}

func TestMap_Set(t *testing.T) {
	m := mapset.NewMap[string, any]()
	m.Set("key", "value")
	assert.Equal(t, true, m.Has("key"))
	assert.Equal(t, "value", m.GetOrDefault("key", "defaultValue"))
}

func TestMap_Get(t *testing.T) {
	m := mapset.NewMap[string, any]()
	m.Set("key", "value")
	value, ok := m.Get("key")
	assert.Equal(t, "value", value)
	assert.Equal(t, true, ok)
	value, ok = m.Get("unset")
	var def any
	assert.Equal(t, def, value)
	assert.Equal(t, false, ok)
}

func TestMap_GetOrDefault(t *testing.T) {
	m := mapset.NewMap[string, any]()
	m.Set("key", "value")
	assert.Equal(t, "value", m.GetOrDefault("key", "unset"))
	assert.Equal(t, "unset", m.GetOrDefault("unset", "unset"))
}

func TestMap_Has(t *testing.T) {
	m := mapset.NewMap[string, any]()
	m.Set("key", "value")
	assert.Equal(t, true, m.Has("key"))
	assert.Equal(t, false, m.Has("unset"))
}

func TestMap_Delete(t *testing.T) {
	m := mapset.NewMap[string, any]()
	m.Set("key", "value")
	assert.Equal(t, true, m.Has("key"))
	m.Delete("key")
	assert.Equal(t, false, m.Has("key"))
}

func TestMap_Size(t *testing.T) {
	m := mapset.NewMap[string, any]()
	assert.Equal(t, 0, m.Size())
	m.Set("key", "value")
	assert.Equal(t, 1, m.Size())
}

func TestMap_Range(t *testing.T) {
	m := mapset.NewMap[string, string]()
	m.Set("key", "value")
	m.Set("key1", "value")
	m.Set("key2", "value")
	flag := false
	count := 0
	m.Range(func(key string, val string) bool {
		count++
		flag = true
		return key == "key1" || key == "key2"
	})
	assert.Equal(t, true, count < m.Size() && flag)
}

func TestMap_Keys(t *testing.T) {
	m := mapset.NewMap[string, string]()
	m.Set("key", "value")
	assert.Equal(t, []string{"key"}, m.Keys())
}

func TestMap_Values(t *testing.T) {
	m := mapset.NewMap[string, string]()
	m.Set("key", "value")
	assert.Equal(t, []string{"value"}, m.Values())
}

func TestMap_Merge(t *testing.T) {
	m := mapset.NewMap[string, string]()
	m.Set("key", "value")
	merge := mapset.NewMap[string, string]()
	merge.Set("key", "mergeKey")
	merge.Set("key1", "mergeKey1")
	m.Merge(merge)
	assert.Equal(t, 2, m.Size())
}
