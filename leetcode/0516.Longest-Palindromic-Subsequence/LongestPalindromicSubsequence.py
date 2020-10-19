class Solution:

    # Runtime: 496 ms, faster than 98.41% of Python3 online submissions for Longest Palindromic Subsequence.
    def longestPalindromeSubseq(self, s: str) -> int:
        n = len(s)
        if s == s[::-1]:
            # without this early return section, Runtime will extend to 1228 ms
            # Runtime: 1228 ms, faster than 81.11% of Python3 online submissions for Longest Palindromic Subsequence.
            return n

        dp = [[0]*n for _ in range(n)]
        for i in range(n):
            dp[i][i] = 1

        # final: i --> 0 and  j --> n
        # xxxxxxxxxx
        #      i  j
        for i in range(n-2, -1, -1):
            for j in range(i+1, n):
                if s[i] == s[j]:
                    dp[i][j] = dp[i+1][j-1] + 2
                else:
                    dp[i][j] = dp[i+1][j] if dp[i+1][j] > dp[i][j-1] else dp[i][j-1]
                        
        import pprint
        print(s)
        pprint.pprint(dp)
        return dp[0][-1]

if __name__ == "__main__":
    print(Solution().longestPalindromeSubseq("abxba"))
