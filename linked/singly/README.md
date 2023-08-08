# 单向链表

## 基本使用

```go
package main

import (
	"fmt"

	"github.com/aide-cloud/slices/linked/singly"
)

func main() {
	h := singly.New(singly.WithValues(1, 2, 3, 4, 5))
	fmt.Println(h.Show())
	list := h.Slice()
	fmt.Println(fmt.Sprintf("cap: %d, len: %d, %v", cap(list), len(list), list))

	h.Append(6).Append(7).Append(8).Append(9).Append(10)
	fmt.Println(h.Show())
	list = h.Slice()
	fmt.Println(fmt.Sprintf("cap: %d, len: %d, %v", cap(list), len(list), list))
	h.Remove(1)
	fmt.Println(h.Show())

	h.RemoveValue(func(val int) bool {
		return val >= 8
	})

	h.Prepend(100).Append(200).Append(300).Append(400).Append(500)
	fmt.Println(h.Show())
	list = h.Slice()
	fmt.Println(fmt.Sprintf("cap: %d, len: %d, %v", cap(list), len(list), list))
}
```

## 输出

```bash
(1)-->(2)-->(3)-->(4)-->(5)
cap: 5, len: 5, [1 2 3 4 5]
(1)-->(2)-->(3)-->(4)-->(5)-->(6)-->(7)-->(8)-->(9)-->(10)
cap: 10, len: 10, [1 2 3 4 5 6 7 8 9 10]
(1)-->(3)-->(4)-->(5)-->(6)-->(7)-->(8)-->(9)-->(10)
(100)-->(1)-->(3)-->(4)-->(5)-->(6)-->(7)-->(200)-->(300)-->(400)-->(500)
cap: 11, len: 11, [100 1 3 4 5 6 7 200 300 400 500]
```