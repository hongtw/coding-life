
from LongestPalindromicSubsequence import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ("bbbab", 4),
    ("cbbd", 2)
]

@pytest.mark.parametrize(
    "s, ans", 
    deepcopy(TEST_SUITS)
)
def test(s, ans):
    assert SOL.longestPalindromeSubseq(s) == ans
    