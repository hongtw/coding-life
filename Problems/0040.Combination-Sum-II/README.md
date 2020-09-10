# [40. Combination Sum II](https://leetcode.com/problems/combination-sum-ii/)

## 解題思路

可由 [#39 (Combination Sum](https://leetcode.com/problems/combination-sum/)) 進行微調即可。先 sort 後就可以變很好處理:若 loop 時，現在的值與上次相同就略過，避免重複使用;而既然有 sort 過了，因此 loop 的 value 大於 target 時就可以提早 break，達到加速。