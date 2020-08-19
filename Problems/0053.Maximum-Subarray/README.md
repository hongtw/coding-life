# [53. Maximum Subarray](https://leetcode.com/problems/maximum-subarray/)

Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

## Example

```text
Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

## 解題思路

一步步累加，如果這段加總加下一個值 n 卻比 n 小，那就從下一個值 n 開始重新累加，並在過程中記錄全局最大值
