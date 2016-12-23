#!/usr/bin/env python

import re
import sys

password = list(sys.argv[1])

def rotate(k):
   global password
   k %= len(password)
   password = password[-k:] + password[:-k]

for line in open(sys.argv[2]):
   #print password
   m = re.search(r"swap position (\d+) with position (\d+)", line)
   if m:
      x = int(m.group(1))
      y = int(m.group(2))
      password[x], password[y] = password[y], password[x]
      continue

   m = re.search(r"swap letter (\w) with letter (\w)", line)
   if m:
      x, y = m.group(1, 2)
      for i in range(len(password)):
         if password[i] == x:
            password[i] = y
         elif password[i] == y:
            password[i] = x
      continue

   m = re.search(r"rotate left (\d+) steps?", line)
   if m:
      rotate(-int(m.group(1)))
      continue

   m = re.search(r"rotate right (\d+) steps?", line)
   if m:
      rotate(int(m.group(1)))
      continue

   m = re.search(r"rotate based on position of letter (\w)", line)
   if m:
      c = m.group(1)
      for i in range(len(password)):
         if password[i] == c:
            if i >= 4 :
               d = 1
            else:
               d = 0
            rotate(1 + i + d)
            break
      continue

   m = re.search(r"reverse positions (\d+) through (\d+)", line)
   if m:
      x = int(m.group(1))
      y = int(m.group(2))
      password[x:y+1] = reversed(password[x:y+1])
      continue

   m = re.search(r"move position (\d+) to position (\d+)", line)
   if m:
      x = int(m.group(1))
      y = int(m.group(2))
      c = password.pop(x)
      password.insert(y, c)
      continue

   print "Unknown operation:", line,

print ''.join(password)
