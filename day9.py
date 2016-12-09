#!/usr/bin/env python

import re
import sys

input = sys.stdin.read().strip()

markerRe = re.compile(r"[(](\d+)x(\d+)[)]")

i = 0
n = 0
while i < len(input):
    m = markerRe.search(input, i)
    if not m:
        n += len(input) - i
        break

    j = m.start()
    n += j - i
    chars = int(m.group(1))
    reps = int(m.group(2))
    n += chars * reps
    i = m.end() + chars

print n

