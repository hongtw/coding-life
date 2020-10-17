
from FibonacciNumber import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    (3, 2),
    (4, 3),
    (5, 5), 
]

@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, ans):
    assert SOL.fib(nums) == ans
    