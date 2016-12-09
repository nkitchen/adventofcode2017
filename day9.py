#!/usr/bin/env python

import re
import sys

markerRe = re.compile(r"[(](\d+)x(\d+)[)]")

def decompressedLen(input):
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
        repeated = input[m.end() : m.end() + chars]
        n += reps * decompressedLen(repeated)
        i = m.end() + chars

    return n

print decompressedLen(sys.stdin.read().strip())
