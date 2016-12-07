import fileinput

input = [map(int, line.split()) for line in fileinput.input()]
i = 0
n = 0
while i < len(input):
    for j in range(0, len(input[i])):
        sides = [input[k][j] for k in range(i, i + 3)]
        a, b, c = sorted(sides)
        if a + b > c:
            n += 1
    i += 3
print n
