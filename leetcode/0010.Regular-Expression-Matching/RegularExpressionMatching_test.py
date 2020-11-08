
from RegularExpressionMatching import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ("aa", "a", False),
    ("aa", "a*", True),
    ("ab", ".*", True),
    ("aab", "c*a*b", True),
]

@pytest.mark.parametrize(
    "string, pattern, ans", 
    deepcopy(TEST_SUITS)
)
def test(string, pattern, ans):
    assert SOL.isMatch(string, pattern) == ans
    assert SOL.isMatchV2(string, pattern) == ans
    