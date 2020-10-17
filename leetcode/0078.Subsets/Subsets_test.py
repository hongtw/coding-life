
from Subsets import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
    ([1, 2, 3], [
        [3],
        [1],
        [2],
        [1,2,3],
        [1,3],
        [2,3],
        [1,2],
        []
        ])
]

@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, ans):
    res = SOL.subsets(nums)
    for sub in ans:
        assert sub in res
        
@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)
def testV2(nums, ans):
    res = SOL.subsetsV2(nums)
    for sub in ans:
        assert sub in res