# [15. 3sum](https://leetcode.com/problems/3sum/)

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

## Note

The solution set must not contain duplicate triplets.

## Example

```text
Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```

## 解題思路

- v1: 使用 unique 且排序過的 nums 去做雙層迴圈
- v2: 在做過 `16. 3sum-closest`後，回頭使用 #16 的思維去解這題。對 nums 排序後，使用 3 個指針以及大小於 0 或等於 0  做指針的控制進退。
  - 不知道是否 Python 排序比較耗時，使用此法後速度掉很多