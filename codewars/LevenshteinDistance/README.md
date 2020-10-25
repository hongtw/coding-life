## Benchmark

```shell
$ python levenshtein.py

[Test] levenshtein is OK
[Benchmark] levenshtein: 0.02800 s
-
[Test] levenshteinDP is OK
[Benchmark] levenshteinDP: 8.04390 s
-
[Test] levenshteinQueue is OK
[Benchmark] levenshteinQueue: 28.02818 s
-
All test is done
```

有趣的是 **levenshteinQueue** ，雖然在這邊 Benchmark 表現不好，但理論上很高效，且的確在 [LeetCode#72](https://leetcode.com/problems/edit-distance/) 使用 `Python` 拿下極高的[成績](https://github.com/hongtw/coding-life/tree/master/leetcode/0072.Edit-Distance)，
推測要夠多重複的字串才會獲得高效
