from functools import lru_cache
from collections import deque

class Solution:

    # Runtime: 240 ms, faster than 19.64% of Python3 online submissions for Edit Distance.
    def minDistance(self, word1: str, word2: str) -> int:
        def lev(a, b, memo={}):
            if not a or not b:
                return max(len(a), len(b))
            if (a, b) in memo:
                return memo[(a, b)]
            if a[0] == b[0]:
                return lev(a[1:], b[1:])
            
            dist = 1 + min(
                lev(a[1:], b[1:]),  # substitution
                lev(a[0:], b[1:]),  # insert or delete
                lev(a[1:], b[0:]),  # insert or delete
            )
            memo[(a, b)] = dist
            return dist
        
        return lev(word1, word2)

    # Runtime: 244 ms, faster than 18.60% of Python3 online submissions for Edit Distance.
    def minDistanceV2(self, word1: str, word2: str) -> int:
        memo = {}
        def lev(a, b, memo):
            if not a or not b:
                return max(len(a), len(b))
            if (a, b) in memo:
                return memo[(a, b)]
            if a[0] == b[0]:
                return lev(a[1:], b[1:], memo)
            
            dist = 1 + min(
                lev(a[1:], b[1:], memo),  # substitution
                lev(a[0:], b[1:], memo),  # insert or delete
                lev(a[1:], b[0:], memo),  # insert or delete
            )
            memo[(a, b)] = dist
            return dist
        return lev(word1, word2, memo)

    # Runtime: 96 ms, faster than 95.57% of Python3 online submissions for Edit Distance.
    def minDistanceV3(self, word1: str, word2: str) -> int:

        @lru_cache(maxsize=None)
        def lev(a, b):
            if not a or not b:
                return max(len(a), len(b))
            if a[0] == b[0]:
                return lev(a[1:], b[1:])
            
            return 1 + min(
                lev(a[1:], b[1:]),  # substitution
                lev(a[0:], b[1:]),  # insert or delete
                lev(a[1:], b[0:]),  # insert or delete
            )
        return lev(word1, word2)

    # Runtime: 96 ms, faster than 95.57% of Python3 online submissions for Edit Distance.
    def minDistanceV3_1(self, word1: str, word2: str) -> int:

        @lru_cache(maxsize=None)
        def lev(a, b):
            if not a or not b:
                return max(len(a), len(b))
            if a[0] == b[0]:
                return lev(a[1:], b[1:])
            l1, l2 = len(a), len(b)
            p1, p2 = 0, 0
            while p1 < l1 and p2 < l2 and a[p1] == b[p2]:
                p1 += 1
                p2 += 1
            
            return 1 + min(
                lev(a[p1+1:], b[p2+1:]),  # substitution
                lev(a[p1:], b[p2+1:]),  # insert or delete
                lev(a[p1 + 1:], b[p2:]),  # insert or delete
            )
        return lev(word1, word2)


    # Runtime: 52 ms, faster than 99.78% of Python3 online submissions for Edit Distance.
    def minDistanceV4(self, word1: str, word2: str) -> int:
        l1, l2 = len(word1), len(word2)
        visited = set()
        queue = [(0, 0, 0)]
        while queue:
            p1, p2, dist = queue.pop(0)
            if (p1, p2) in visited or p1 > l1 or p2 > l2:
                continue
            if word1[p1:] == word2[p2:]:
                return dist

            visited.add((p1, p2))
            while p1 < l1 and p2 < l2 and word1[p1] == word2[p2]:
                p1 += 1
                p2 += 1
            
            queue.extend(
                [
                    (p1 +1, p2 +1, dist +1),
                    (p1, p2 +1, dist +1),
                    (p1 +1, p2, dist +1),
                ]
            )

    # Runtime: 40 ms, faster than 100.00% of Python3 online submissions for Edit Distance.
    def minDistanceV5(self, a: str, word2: b) -> int:
        la, lb = len(a), len(b)
        isVisited = set()
        queue = [(0, 0, 0)] # (p1, p2, dist)
        while queue:
            p1, p2, dist = queue.pop(0)
            if (p1, p2) in isVisited:
                continue
            isVisited.add((p1, p2))
            while p1 < la and p2 < lb and a[p1] == b[p2]:
                    p1 += 1
                    p2 += 1
            if p1 >= la and p2 >= lb:
                return dist

            dist += 1
            queue.extend([
                (p1 + 1, p2 + 1, dist),
                (p1, p2 + 1, dist),
                (p1 + 1, p2, dist),
            ])