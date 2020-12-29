import re

## Lookbehind
'''
(?=...)
Matches if ... matches next, but doesn’t consume any of the string. 
This is called a lookahead assertion.
'''

lk = re.compile("John(?= Lee)")
print(f"Lookbehind: {lk}")
for sample in [
    "John Lee",
    "XDD John Lee",
    "John XLee",
    ]:
    print(f"{sample: <18}: {lk.search(sample)}")
# output
'''
Positive Lookbehind: re.compile('John(?= Lee)')
John Lee          : <re.Match object; span=(0, 4), match='John'>
XDD John Lee      : <re.Match object; span=(4, 8), match='John'>
John XLee         : None
'''


## Negative Lookbehind
'''
(?!...)
Matches if ... doesn’t match next. 
This is a negative lookahead assertion. For example, Isaac (?!Asimov) will match 'Isaac ' only if it’s not followed by 'Asimov'.
'''
nk = re.compile("John(?! Lee)")
print(f"\nNegative Lookbehind: {nk}")
for sample in [
    "John Lee",
    "XDD John Lee",
    "John XLee",
    ]:
    print(f"{sample: <18}: {nk.search(sample)}")