from threeSum import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  [
        ("Example-1", [-1, 0, 1, 2, -1, -4], [
            [-1, 0, 1],
            [-1, -1, 2]
        ]),
]

@pytest.mark.parametrize(
    "example, nums, ans", 
    deepcopy(TEST_SUITS)
)
def test(example, nums, ans):
    res = SOL.threeSum(nums)
    assert len(res) == len(ans)

    res_set = set([tuple(sorted(arr)) for arr in res])
    ans_set = set([tuple(sorted(arr)) for arr in ans])
    assert res_set - ans_set == set()