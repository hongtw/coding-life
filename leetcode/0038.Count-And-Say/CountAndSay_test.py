from CountAndSay import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS = [
    (1, "1"),
    (2, "11"),
    (3, "21"),
    (4, "1211"),
    (5, "111221"),
]


@pytest.mark.parametrize("nums, ans", deepcopy(TEST_SUITS))
def test(nums, ans):
    assert SOL.countAndSay(nums) == ans
