ips = [
    "0.0.0.0",
    "127.0.0.1",
    "127.0.49.1",
    "127.0.64.1",
    "127.65.64.1",
    "127.98.0.1",
    "128.128.4.11",
    "128.57.107.76",
    "128.68.4.11",
    "128.96.107.55",
    "128.99.107.55",
    "128.99.58.55",
    "15.128.4.65",
    "26.56.4.23",
    "67.128.4.11",
    "7.7.7.8",
    "74.0.65.76",
    "77.255.255.254",
]

import re

pattern0_255 = re.compile(r"1\d\d|2[0-4]\d|25[0-5]|\d\d|\d")
assert all([pattern0_255.match(str(n)) for n in range(256)])

patten = re.compile(r"\b((1\d\d|2[0-4]\d|25[0-5]|\d\d|\d)\.){3}(1\d\d|2[0-4]\d|25[0-5]|\d\d|\d)\b")

for ip in ips:
    assert patten.search(ip) is not None, f"{ip} should match"

for ip in ["77.255.255.2555", "0.1111.2.2", "1111.2.3.4"]:
    assert patten.search(ip) is None, f"{ip} should not match"

print("Test Done")