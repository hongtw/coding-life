from functools import lru_cache

class Solution:

    # Runtime: 32 ms, faster than 98.87% of Python3 online submissions for Regular Expression Matching.
    def isMatch(self, s: str, p: str) -> bool:
        len_s = len(s)
        len_p = len(p)
        
        @lru_cache(None)
        def dp(i, j):
            if j == len_p:
                return i == len_s
            
            # print("***", i, j, s[i], p[j])
            isFirstOK = i < len_s and p[j] in (".", s[i])

            if j + 1 < len_p and p[j + 1] == '*':
                return dp(i, j + 2) or \
                    isFirstOK and dp(i + 1, j)
            else:
                return isFirstOK and dp(i+1, j+1)
        return dp(0, 0)

if __name__ == "__main__":
    print(Solution().isMatch('aa', "a*"))