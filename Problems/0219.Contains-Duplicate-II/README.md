# [219. Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii/)

Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.

## Example 1

```text
Input: nums = [1,2,3,1], k = 3
Output: true
```

## Example 2

```text
Input: nums = [1,0,1,1], k = 1
Output: true
```

## Example 3

```text
Input: nums = [1,2,3,1,2,3], k = 2
Output: false
```

## 解題思路

- 方法1: 維護一個長度 k 的 cache，每次比對完新的值未在 cache 裡，則移除 k 個之前的值並新增該值
- 方法2: 建立 map 並記錄各值的 index，若新值在 map 裡，且 curr_index - prev_index <= k 的話回傳 True
