#!/usr/bin/env python

import sys

initStr = '110010110100'

def format(n, bits):
    return "{0:0{1}b}".format(bits, n)

def expand(n, bits):
    mask = (1 << n) - 1
    invBits = bits ^ mask
    invDigits = reversed(format(n, invBits))
    m = 2 * n + 1
    invStr = ''.join(invDigits)
    expBits = bits << (n + 1) | int(invStr, 2)
    return m, expBits

def checksum(n, bits):
    while n % 2 == 0:
        s = 0
        for i in range(0, n / 2):
            a = 1 & (bits >> (2 * i))
            b = 1 & (bits >> (2 * i + 1))
            if a == b:
                s |= 1 << i
        n /= 2
        bits = s
    return n, s

initStr = sys.argv[1]
n = int(sys.argv[2])

k = len(initStr)
bits = int(initStr, 2)

while k < n:
    k, bits = expand(k, bits)

bits >>= k - n

#print format(n, bits)

cn, cbits = checksum(n, bits)
print format(cn, cbits)
