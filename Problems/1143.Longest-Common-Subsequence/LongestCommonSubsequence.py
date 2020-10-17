class Solution:

    # Runtime: 376 ms, faster than 90.48% of Python3 online submissions for Longest Common Subsequence.
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        len1 = len(text1)
        len2 = len(text2)
        dp = [[0]*(len2+1) for _ in range(len1+1)]
        
        for r in range(1, len1+1):
            for c in range(1, len2+1):
                if text1[r - 1] == text2[c-1]:
                    dp[r][c] = dp[r-1][c-1] + 1
                else:
                    dp[r][c] = max(dp[r-1][c], dp[r][c-1])

        # from pprint import pprint
        # pprint(dp)
        return dp[len1][len2]

    # Runtime: 336 ms, faster than 96.56% of Python3 online submissions for Longest Common Subsequence.
    def longestCommonSubsequenceV2(self, text1: str, text2: str) -> int:
        len1 = len(text1)
        len2 = len(text2)
        dp = [[0]*(len2+1) for _ in range(len1+1)]
        
        for r in range(1, len1+1):
            for c in range(1, len2+1):
                if text1[r - 1] == text2[c-1]:
                    dp[r][c] = dp[r-1][c-1] + 1
                else:
                    dp[r][c] = dp[r-1][c] if dp[r-1][c] > dp[r][c-1] else dp[r][c-1] 
        return dp[len1][len2]

if __name__ == "__main__":
    Solution().longestCommonSubsequence("babcde", "acxe")
    Solution().longestCommonSubsequence("babccde", "acxe")
    Solution().longestCommonSubsequence("abcba", "abcbcbaaaaaa")