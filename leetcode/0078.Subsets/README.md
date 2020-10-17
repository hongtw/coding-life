# [78. Subsets](https://leetcode.com/problems/subsets/)

## 解題思路

起始一個空的 slice ，並且 loop nums，並對起始的 slice 內的元素逐一去 append 並再置入進去，動態的擴增此 slice。

---

### Golang 語言特性

動態的去 loop slice 並且過程中持續 append，loop 迴圈只會走 slice 起始的 element 而已，如下:

```golang
package main

import "fmt"

func main() {
    sl := []int{100, 200, 300}

    for _, n := range sl {
        sl = append(sl, n+1)
        fmt.Println(n, sl)
    }
}

// output
// 100 [100 200 300 101]
// 200 [100 200 300 101 201]
// 300 [100 200 300 101 201 301]
```

如果 Python3 用此寫法的話，則會陷入無窮迴圈

```python
a = [100, 200, 300]
for n in a:
    print(n, a)
    a.append(n+1)
```
