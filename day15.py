#!/usr/bin/env python

import fileinput
import itertools
import re

p = {}
a = {}

for line in fileinput.input():
    m = re.search(r"Disc #(\d+) has (\d+) positions; at time=0, it is at position (\d+)[.]", line)
    if not m:
        continue

    i = int(m.group(1))
    p[i] = int(m.group(2))
    a[i] = int(m.group(3))

discs = sorted(p.keys())

# Disc #i is at position 0 at time=(-a[i] % p[i]).
# When the capsule reaches disc #i at time=(t0 + i),
# it will be at position (t0 + i + a[i]) % p[i].
i = discs[0]
times = itertools.count((-a[i] - i) % p[i], p[i])
times, copy = itertools.tee(times, 2)
print list(itertools.compress(copy, [1] * 5))

for i in (discs[1:]):
    def g(i):
        return lambda t: (t + i + a[i]) % p[i] == 0

    def f(t):
        print "({} + {} + {}) % {} == 0? {}".format(
            t, i, a[i], p[i], (t + i + a[i]) % p[i])
        return (t + i + a[i]) % p[i] == 0
    times = itertools.ifilter(
        # t == (-a[i] - i) % p[i] + k * p[i] for some k
        # lambda t: (t + i + a[i]) % p[i] == 0,
        g(i),
        times
    )

print next(times)
