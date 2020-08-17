from argparse import ArgumentParser
from pathlib import Path
SELF = Path(__file__).resolve()
PROBLEMS_FOLDER = SELF.parent

def parse_problem_name(url):
    # https://leetcode.com/problems/two-sum/

    if not url.startswith("https://leetcode.com/problems/"):
        raise RuntimeError("Invalid url")
    
    name = url.split("https://leetcode.com/problems/")[1].strip('/')
    capitalized_dash_name = '-'.join(map(str.capitalize, name.split('-')))
    camelName = ''.join(map(str.capitalize, name.split('-')))
    return capitalized_dash_name, camelName

ap = ArgumentParser(__file__)
ap.add_argument("-n", "--num", required=True, help="number of problems in leetcode")
ap.add_argument("-u", "--url", required=True, help="url of problems in leetcode")
args = ap.parse_args()

capitalized_dash_name, camelName = parse_problem_name(args.url)
new_folder = PROBLEMS_FOLDER / f"{args.num:0>4}.{capitalized_dash_name}"
new_folder.mkdir(exist_ok=True)
(new_folder / f"{camelName}.py").touch()
pytest_file = (new_folder / f"{camelName}_test.py")
pytest_file.touch()
if pytest_file.stat().st_size == 0:
    pytest_file.write_text(
f"""
from {camelName} import Solution
import pytest
from copy import deepcopy

SOL = Solution()

TEST_SUITS =  []

@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, ans):
    assert SOL.func(nums) == ans
    """
    )
(new_folder / f"{camelName}.go").touch()
(new_folder / f"{camelName}_test.go").touch()
readme = new_folder / f"README.md"
readme.touch()
if readme.stat().st_size == 0:
    readme.write_text(
        f"# [{args.num}. {capitalized_dash_name.replace('-', ' ')}]({args.url})\n"
    )