import fileinput

debug = False
def dprint(*args):
    if debug:
        print " ".join( str( a ) for a in args )

def hq():
    x = [0, 0]
    v = [0, 1]
    visits = {tuple(x): 1}
    for line in fileinput.input():
        steps = line.split( ", " )
        for s in steps:
            if s[0] == "L":
                v = [-v[1], v[0]]
            elif s[0] == "R":
                v = [v[1], -v[0]]
            else:
                assert False, "Bad step " + s
            d = int(s[1:])
            dprint(v)
            for k in range(1, d + 1):
                x[0] += v[0]
                x[1] += v[1]
                tx = tuple(x)
                visits[tx] = 1 + visits.get(tx, 0)
                if visits[tx] == 2:
                    return x
    return None
x = hq()
print x, abs(x[0]) + abs(x[1])

# vim: set shiftwidth=4 :
