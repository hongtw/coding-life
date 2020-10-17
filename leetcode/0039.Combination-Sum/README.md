# [39. Combination Sum](https://leetcode.com/problems/combination-sum/)

## 解題思路

1. 使用遞迴處理這個問題
2. 優化: 遞迴的function，比起每次都回傳新的 array result 到最外層去 append，直接傳遞 result 的位址給遞迴的function在過程中 append 這樣更快，並且在 golang 和 python 速度上都有增加
3. 在處理 golang 時，若加進 result 的 slice 沒有特別 copy，在某些測資下會發生預期外的錯誤。
