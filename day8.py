import fileinput
import numpy as np
import re
import sys

width = int(sys.argv[1])
height = int(sys.argv[2])

screen = np.zeros((height, width))

for line in fileinput.input(sys.argv[3:]):
    m = re.search(r"rect (\d+)x(\d+)", line)
    if m:
        w = int(m.group(1))
        h = int(m.group(2))
        screen[:h,:w] = 1
        continue

    m = re.search(r"rotate column x=(\d+) by (\d+)", line)
    if m:
        x = int(m.group(1))
        d = int(m.group(2))
        screen[..., x] = np.roll(screen[..., x], d)
        continue

    m = re.search(r"rotate row y=(\d+) by (\d+)", line)
    if m:
        y = int(m.group(1))
        d = int(m.group(2))
        screen[y, ...] = np.roll(screen[y, ...], d)
        continue

print screen
print np.sum(screen)
