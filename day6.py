import fileinput

freq = []

for line in fileinput.input():
    line = line.rstrip()
    while len(freq) < len(line):
        freq.append({})

    for i, c in enumerate(line):
        freq[i][c] = 1 + freq[i].get(c, 0)

msg = ""
for f in freq:
    m = sorted(f.items(), key=lambda e: e[1])[0]
    msg += m[0]
print msg
