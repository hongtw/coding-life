class Solution:

    # Runtime: 39 ms, faster than 98.81% of Python3 online submissions for Count and Say.
    # Memory Usage: 14 MB, less than 53.55% of Python3 online submissions for Count and Say.
    def countAndSay(self, n: int) -> str:
        if n == 1:
            return "1"

        prevAns = self.countAndSay(n - 1)
        char = prevChar = prevAns[0]
        count = 1
        ans = ""
        for char in prevAns[1:]:
            if char != prevChar:
                ans += f"{count}{prevChar}"
                prevChar = char
                count = 0
            count += 1

        ans += f"{count}{char}"

        return ans

