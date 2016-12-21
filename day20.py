import fileinput
import re

blocks = []

for line in fileinput.input():
   m = re.search(r"(\d+)-(\d+)", line)
   if m:
      a = int(m.group(1))
      b = int(m.group(2))
      blocks.append((a, b))

blocks.sort()

for a, b in blocks:
   x = b + 1
   if not any(c <= x <= d for c, d in blocks):
      print x
      break
