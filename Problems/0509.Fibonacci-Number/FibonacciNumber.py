class Solution:

    # Runtime: 24 ms, faster than 94.52% of Python3 online submissions for Fibonacci Number.
    def fib(self, N: int) -> int:
        if N in (1, 2): return 1
        curr = prev = 1
        count = 0
        for _ in range(2, N):
            count = curr + prev
            prev = curr
            curr = count
        
        return count