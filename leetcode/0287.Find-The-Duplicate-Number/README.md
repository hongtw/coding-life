# [287. Find The Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/)

## Some different between Python3 and Golang

Python3:

```Python
a = [1, 2, 3]
pt = 0

pt, a[pt] = a[pt], 0
print(pt)    # 1
print(a)     # [1, 0, 3]
```

Golang:

```golang
package main

import (
    "fmt"
)

func main() {
    a := []int{1, 2, 3}
    pt := 0

    pt, a[pt] = a[pt], 0
    fmt.Println(pt)     // 1
    fmt.Println(a)      // [0 2 3]
}
```
