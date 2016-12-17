import md5
import re
import sys

salt = sys.argv[1]

hashes = {}

def hash(i):
   h = hashes.get(i)
   if not h:
       s = salt + str(i)
       for j in xrange(2017):
           s = md5.new(s).hexdigest()
       h = s
       hashes[i] = h
   return h

def isKey(i):
    m = re.search(r"(.)\1\1", hash(i))
    if not m:
        return False

    s = m.group(1) * 5
    for j in range(i + 1, i + 1001):
        if s in hash(j):
            return True

    return False

n = 64
print '.' * n
keys = []
i = 0
while len(keys) < n:
    if isKey(i):
        keys.append(i)
        sys.stdout.write('=')
        sys.stdout.flush()
    i += 1
sys.stdout.write('\n')
print keys[-1]
