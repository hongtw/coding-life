
from MaximumProductOfThreeNumbers import Solution
import pytest
import random
import timeit
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([1, 2, 3], 6),
    ([1, 2, 3, 4], 24),
    ([-3, -5, 2, 3], 45),
    ([-1,-2,-3], -6),
    ([-1,-2,-3, -4], -6),
]

BENCHMARK_SUITS = list(range(1000)) + list(range(-1, -1000, -1))
random.shuffle(BENCHMARK_SUITS)

@pytest.mark.parametrize(
    "nums, ans", 
    TEST_SUITS
)
def test(nums, ans):
    assert SOL.maximumProduct(deepcopy(nums)) == ans
    assert SOL.maximumProductV2(deepcopy(nums)) == ans
