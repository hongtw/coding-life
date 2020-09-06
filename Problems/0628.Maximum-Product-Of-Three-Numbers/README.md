# [628. Maximum Product Of Three Numbers](https://leetcode.com/problems/maximum-product-of-three-numbers/)

## 解題思路

### 方法1

sort 過後，以下兩者取大值

- 最大一位配最小兩位(最小兩位為負值的情況) `nums[-1]*nums[0]*nums[1]`
- 最大三位: `nums[-1]*nums[-2]*nums[-3]`

時間複雜度為 `O(n log n)`，但使用 Python 在 leetcode 提交，反而方法1比較快，可能測資不夠複雜

### 方法2

loop array，找出最大的三個數以及最小的兩個數去比對乘積，找出最大值，時間複雜度為 `O(n)`
