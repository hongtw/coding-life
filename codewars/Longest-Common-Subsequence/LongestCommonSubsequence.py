def lcs(text1: str, text2: str) -> str:
    len1 = len(text1)
    len2 = len(text2)
    dp = [[(0, "")]*(len2+1) for _ in range(len1+1)]
    
    for r in range(1, len1+1):
        for c in range(1, len2+1):
            if text1[r-1] == text2[c-1]:
                dp[r][c] = (dp[r-1][c-1][0] + 1, dp[r-1][c-1][1] + text1[r-1])
            else:
                dp[r][c] = max(dp[r-1][c], dp[r][c-1], key=lambda item: item[0])
    return dp[len1][len2][1]

def lcs_hjcain(x, y):
    hashtable = {}
    for idx, c in enumerate(x):
        c_indices = hashtable.get(c, [])
        c_indices.append(idx)
        hashtable[c] = c_indices

    lcs = ""
    leftPos = -1
    for c in y:
        if c in hashtable: # hash table lookup complexity is O(1)
            available_indices = [ idx for idx in hashtable[c] if idx > leftPos ]
            if available_indices: # means this character can be used
                lcs += c
                leftPos = available_indices[0]
                hashtable[c] = available_indices[1:] # save for speed up

    return lcs
    
############
### 以下是在 codewars solution 上看到的各種奇技淫巧
### 擷取幾個來瞻仰

from itertools import combinations

# 精簡的暴力解
def lcs_bruteForce(x, y):
    def subsequences(s):
        """Returns set of all subsequences in s."""
        return set(''.join(c) for i in range(len(s) + 1) for c in combinations(s, i))
    """Returns longest matching subsequence of two strings."""
    return max(subsequences(x).intersection(subsequences(y)), key=len)

def lcs_recursive(x, y):
    if not x or not y: return ""
    if x[0] == y[0]: return x[0] + lcs(x[1:], y[1:])
    
    return max(lcs(x[1:], y), lcs(x, y[1:]), key=len) 

def benchmark():
    print("===== Benchmark =====")
    import timeit
    template = (
        '{0}( "abcdef" , "abc" );'
        '{0}( "abcdef" , "acf" );'
        '{0}( "132535365132535365" , "123456789" );'
    )
    for func in ("lcs", "lcs_hjcain", "lcs_bruteForce", "lcs_recursive"):
        print(f"Benchmark: {func}")
        res = timeit.timeit(template.format(func), setup=f"from __main__ import {func}", number=1000)
        print(f"{res:.5f} s\n-")

if __name__ == "__main__":
    x = "abcs"
    y = "sabc"
    print(lcs(x, y))
    print(lcs_hjcain(x, y))

    benchmark()