# [16. 3sum Closest](https://leetcode.com/problems/3sum-closest/)

## 解題思路

1. 先 sort
2. 固定第一個數字後，剩下的兩個數由兩個指針控制
    - 指針1 為第一個數的下一位
    - 指針2 從最後開始
    - 如果現在三數總和大於 target，指針2向左靠近，若小於 target，指針1 向右靠近，逐漸夾逼出最接近 target 的值。