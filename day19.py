import sys

# The number of elves remaining
n = int(sys.argv[1])

# The first elf still remaining
a = 1

# The last elf still remaining
b = n

# The difference between consecutive elves remaining
k = 2

while a < b:
   if n % 2 == 0:
      # Last elf is out.
      b -= (k - 1)
      if b < a:
         b = a
   else:
      # First elf is out.
      a += k
      if a > b:
         a = b

   n //= 2
   k *= 2

print b

