import fileinput
import os
import re

M = os.environ.get("MAX")
if M:
   M = int(M)
else:
   M = 2 ** 32


endpoints = []

for line in fileinput.input():
   m = re.search(r"(\d+)-(\d+)", line)
   if m:
      a = int(m.group(1))
      b = int(m.group(2))
      endpoints.append([a, 1])
      endpoints.append([b + 1, -1])

nBlocks = 0
nAddrs = 0
curBlockStart = None

endpoints.sort()
for x, d in endpoints:
   nBlocks += d
   if d > 0 and nBlocks == 1 :
      curBlockStart = x
   elif nBlocks == 0 :
      nAddrs += x - curBlockStart

print M - nAddrs
