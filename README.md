# MapSet - 一个泛型的 Map Set 类型库

## 使用

```shell
go get -u github.com/joeescn/mapset
```

## 示例

```golang
package main

import (
 "fmt"

 "github.com/joeescn/mapset"
)

func main() {
 s := mapset.NewSet[int]()
 s.Add(1)
 clone := s.Clone()
 clone.Add(2)

 fmt.Println(clone.IsSuperset(s), clone)
 // Output: true [1,2]

 m := mapset.NewMap[string, string]()
 m.Set("key", "value")
 fmt.Println(m.Get("key"))
 // Output: value true

}

```

## 注意

该实现不是线程安全的，如果需要需要线程安全，可以通过外部加锁实现。

```go
package main

import (
 "log"
 "math/rand"
 "runtime"
 "sync"

 "github.com/joeescn/mapset"
)

func main() {
 runtime.GOMAXPROCS(2)
 ints := rand.Perm(1000)

 s := mapset.NewSet[int]()
 lock := sync.Mutex{}

 var wg sync.WaitGroup
 wg.Add(len(ints))
 for i := 0; i < len(ints); i++ {
  go func(i int) {
   lock.Lock()
   defer lock.Unlock()
   s.Add(i)
   wg.Done()
  }(i)
 }

 wg.Wait()
 for _, i := range ints {
  if !s.Contains(i) {
   log.Fatalf("Set is missing element: %v", i)
  }
 }
}

```
