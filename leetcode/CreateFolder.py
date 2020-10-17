import re
from argparse import ArgumentParser
from pathlib import Path
SELF = Path(__file__).resolve()
PROBLEMS_FOLDER = SELF.parent

def parse_problem_name(url):
    # https://leetcode.com/problems/two-sum/

    if not url.startswith("https://leetcode.com/problems/"):
        raise RuntimeError("Invalid url")
    
    name = url.split("https://leetcode.com/problems/")[1].strip('/')
    name_tokens = [
        token.capitalize() if not re.match("i*$", token) else token.upper()
        for token in name.split('-')
    ]
    capitalized_dash_name = '-'.join(name_tokens)
    capitalized_space_name = ' '.join(name_tokens)

    camelName = ''.join(name_tokens)
    return capitalized_dash_name, capitalized_space_name, camelName

ap = ArgumentParser(__file__)
ap.add_argument("-n", "--num", required=True, help="number of problems in leetcode")
ap.add_argument("-u", "--url", required=True, help="url of problems in leetcode")
args = ap.parse_args()

capitalized_dash_name, capitalized_space_name, camelName = parse_problem_name(args.url)
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

TEST_SUITS =  [
    ([1, 2, 3], 4)
]

@pytest.mark.parametrize(
    "nums, ans", 
    deepcopy(TEST_SUITS)
)
def test(nums, ans):
    assert SOL.func(nums) == ans
    """
)
go_file = (new_folder / f"{camelName}.go")
go_file.touch()
if go_file.stat().st_size == 0:
    go_file.write_text("package leetcode\n\n")

readme = new_folder / f"README.md"
readme.touch()
if readme.stat().st_size == 0:
    readme.write_text(
        f"# [{args.num}. {capitalized_dash_name.replace('-', ' ')}]({args.url})\n"
    )

project_readme = PROBLEMS_FOLDER.parent / 'README.md'
content = project_readme.read_text(encoding='utf8')
if args.url not in content:
    with project_readme.open('a', encoding='utf8') as f:
        f.write(f"- [{args.num.strip('0')}. {capitalized_space_name}](https://github.com/hongtw/leetcode/tree/master/Problems/{args.num:0>4}.{capitalized_dash_name})")