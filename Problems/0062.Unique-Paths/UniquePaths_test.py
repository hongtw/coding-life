
from UniquePaths import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    (3, 2, 3),
    (7, 3, 28),
    (23, 12, 193536720),
]

@pytest.mark.parametrize(
    "m, n, ans", 
    deepcopy(TEST_SUITS)
)
def test(m, n, ans):
    assert SOL.uniquePaths(m, n) == ans
    