import fileinput
import re

used = {}
avail = {}

for line in fileinput.input():
    m = re.search(r"/dev/grid/(\S+)\s+\S+\s+(\d+)T\s+(\d+)T", line)
    if m:
        node = m.group(1)
        u = int(m.group(2))
        a = int(m.group(3))
        used[node] = u
        avail[node] = a

n = 0
keys = sorted(used.keys())
for a in keys:
    if used[a] == 0:
        continue
    for b in keys:
        if a == b:
            continue
        if used[a] <= avail[b]:
            n += 1

print n
