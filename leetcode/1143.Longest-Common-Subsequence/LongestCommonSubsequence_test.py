
from LongestCommonSubsequence import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ("babcde", "acxe", 3),
    ("babccde", "acxe", 3),
    ("oxcpqrsvwf", "shmtulqrypy", 2),
    ("abcba", "abcbcba", 5),
]

@pytest.mark.parametrize(
    "text1, text2, ans", 
    deepcopy(TEST_SUITS)
)
def test(text1, text2, ans):
    assert SOL.longestCommonSubsequence(text1, text2) == ans
    