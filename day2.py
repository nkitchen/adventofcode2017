import fileinput

keypad = [
   ". . 1 . .".split(),
   ". 2 3 4 .".split(),
   "5 6 7 8 9".split(),
   ". A B C .".split(),
   ". . D . .".split(),
]

delta = {
    'U': (-1, 0),
    'D': (1, 0),
    'L': (0, -1),
    'R': (0, 1),
}

pos = [2, 0]
res = []
for line in fileinput.input():
    for move in list(line.rstrip()):
       d = delta[move]
       n0 = pos[0] + d[0]
       n1 = pos[1] + d[1]
       if n0 < 0 or n0 >= len(keypad):
          continue
       if n1 < 0 or n1 >= len(keypad[n0]):
          continue
       if keypad[n0][n1] != ".":
          pos = [n0, n1]
    digit = keypad[pos[0]][pos[1]]
    res.append(digit)
print "".join(res)
