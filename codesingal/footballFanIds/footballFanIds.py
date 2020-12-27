# https://app.codesignal.com/challenge/L9YzWMLMJbwSTGRkk
# Use Negative lookahead: (?!xxxx)

import os
import re
ROOT = "/root/devops"

invites = open(f"{ROOT}/invite.info").read().splitlines()
bans = open(f"{ROOT}/ban.info").read().splitlines()
pat = re.compile(
    f'{ROOT}/(?!({"|".join(bans)})/\w+\.info)({"|".join(invites)}.*\.info)')

allIds = []
for root, directories, files in os.walk(ROOT): 
    for file in files:
        fp = os.path.join(root, file)
        if pat.match(fp):
            allIds.extend(open(fp).read().splitlines())

for id in sorted(allIds):
    print(id)