# [532. K Diff Pairs In An Array](https://leetcode.com/problems/k-diff-pairs-in-an-array/)

## 解題思路

- 使用 counter 統計所有 nums 後，歷遍這些 unique 的 key，看 n + k 是否也在 counter 中
- 只需要檢查 n + k，不需要理會 n - k ，例如 (1, 3)、(3, 1) 屬於同一種解，只看 n + k 就能走遍所有組合
