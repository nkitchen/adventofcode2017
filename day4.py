import fileinput
import re

sectorSum = 0

for line in fileinput.input():
    m = re.match( r"([a-z]+[-a-z]*[a-z]+)-(\d+)\[([a-z]{5})\]", line)
    if not m:
        continue

    freq = {}
    for c in m.group(1):
        if c.isalpha():
            freq[c] = 1 + freq.get(c, 0)

    byFreq = sorted(
        freq.items(),
        key=lambda i: (-i[1], i[0])
    )

    checksum = m.group(3)
    checksumExp = ''.join(i[0] for i in byFreq[:5])
    if checksum == checksumExp:
        sectorSum += int(m.group(2))

print sectorSum
