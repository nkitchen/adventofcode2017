#!/usr/bin/env python

import string
import sys

nRows = int(sys.argv[1])

tab = string.maketrans(".^", "01")
rtab = string.maketrans("01", ".^")

rowStr = next(sys.stdin).strip()
w = len(rowStr)
r = int(rowStr.translate(tab), 2)
rows = [r]

while len(rows) < nRows:
   rp = rows[-1]
   rpl = rp >> 1
   rpr = rp << 1
   rule1 = rpl & rp & ~rpr
   rule2 = ~rpl & rp & rpr
   rule3 = rpl & ~rp & ~rpr
   rule4 = ~rpl & ~rp & rpr
   r = rule1 | rule2 | rule3 | rule4
   r &= (1 << w) - 1
   rows.append(r)

#for r in rows:
#   print "{0:0{1}b}".format(r, w).translate(rtab)

def bitCount(r):
   n = 0
   while r != 0:
      n += 1
      r &= r - 1
   return n

safe = 0
for r in rows:
   safe += len(rowStr) - bitCount(r) 

print safe
