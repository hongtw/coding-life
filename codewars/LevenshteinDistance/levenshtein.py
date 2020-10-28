def levenshtein(a, b, memo={}):
    if (a, b) in memo:
        return memo[(a, b)]
    if not a or not b:
        return max(len(a), len(b))

    isNeedSub = int(a[0] != b[0])
    dist = min(
        levenshtein(a[1:], b[1:]) + isNeedSub,  # substitution or pass
        levenshtein(a[0:], b[1:]) + 1,          # insert or delete
        levenshtein(a[1:], b[0:]) + 1,          # insert or delete
    )
    memo[(a, b)] = dist
    return dist


def levenshteinDP(a, b):
    ''' DP Table
        ''   s   i   t   t   i   n   g
    ''   0   1   2   3   4   5   6   7
    k    1   1   2   3   4   5   6   7 
    i    2   2   1   2   3   4   5   6
    t    3   3
    c
    h
    e
    n
    '''

    la = len(a)
    lb = len(b)
    dp = [[i] + [0]*(lb) for i in range(la + 1)]
    dp[0][:] = range(lb+1)


    for i in range(1, la + 1):
        for j in range(1, lb +1):
            isNeedSub = int(a[i-1] != b[j-1])
            dp[i][j] = min(
                dp[i-1][j-1] + isNeedSub,
                dp[i][j-1] + 1,
                dp[i-1][j] + 1,
            )

    # from pprint import pprint
    # pprint(dp)
    return dp[la][lb]

def levenshteinQueue(a, b):
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


def benchmark(func):
    import timeit
    import string
    import random
    def get_random_string(length):
        letters = string.ascii_lowercase
        return ''.join(random.choice(letters) for i in range(length))
        
        

    template = (
        f"{{0}}('kitten', 'sitting');"
        f"{{0}}('XDD', 'XD');"
        f"{{0}}('horse', 'ros');"
        f"{{0}}('intention', 'execution');"
        f"{{0}}('{get_random_string(100)}', '{get_random_string(150)}');"
    )

    # print(f"[Benchmark] Start {func.__name__}")
    res = timeit.timeit(template.format(func.__name__), setup=f"from __main__ import {func.__name__}", number=1000)
    print(f"[Benchmark] {func.__name__}: {res:.5f} s")

def test(func):
    assert func("kitten", "sitting") == 3
    assert func("XDD", "XD") == 1
    assert func("horse", "ros") == 3
    assert func("intention", "execution") == 5
    print(f"[Test] {func.__name__} is OK")


import timeit
for func in (levenshtein, levenshteinDP, levenshteinQueue):
    test(func)
    benchmark(func)
    print("-")
else:
    print("All test is done")

'''
python levenshtein.py
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
'''