#!/usr/bin/env python

import re
import sys

def rotate(word, k):
   k %= len(word)
   return word[-k:] + word[:-k]

def descramble(password, ops):
   for line in ops:
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
          password = rotate(password, int(m.group(1)))
          continue

       m = re.search(r"rotate right (\d+) steps?", line)
       if m:
          password = rotate(password, -int(m.group(1)))
          continue

       m = re.search(r"rotate based on position of letter (\w)", line)
       if m:
          c = m.group(1)
          done = False
          for k in range(0, len(password)):
             if done:
                break
             w = rotate(password, k)
             for i in range(len(w)):
                if w[i] == c:
                   if i >= 4 :
                      d = 1
                   else:
                      d = 0
                   wr = rotate(w, 1 + i + d)
                   if wr == password:
                      password = w
                      done = True
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
          c = password.pop(y)
          password.insert(x, c)
          continue

       print "Unknown operation:", line,

   return password

password = list(sys.argv[1])

ops = list(open(sys.argv[2]))

password = descramble(password, reversed(ops))
print ''.join(password)
