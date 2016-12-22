import fileinput
import os
import re

M = os.environ.get("MAX")
if M:
   M = int(M)
else:
   M = 2 ** 32

blocks = []

for line in fileinput.input():
   m = re.search(r"(\d+)-(\d+)", line)
   if m:
      a = int(m.group(1))
      b = int(m.group(2))
      blocks.append((a, b))

blocks.sort()

n = 0
a = blocks[0][0]
if a > 0 :
   n += a

for i in xrange(len(blocks) - 1):
   b = blocks[i][1]
   a = blocks[i + 1][0]
   if b + 1 < a :
      n += a - (b + 1)

b = blocks[-1][1]
a = M
if b + 1 < a :
   n += a - (b + 1)

print n
