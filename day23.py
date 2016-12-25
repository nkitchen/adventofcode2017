import collections
import fileinput
import re

regs = collections.defaultdict(int)

cmds = []
for line in fileinput.input():
    line = line.strip()

    m = re.search(r"cpy ([-\w]+) ([-\w]+)", line)
    if m:
        cmds.append(['cpy', m.group(1), m.group(2)])
        continue

    m = re.search(r"inc (\w+)", line)
    if m:
        cmds.append(['inc', m.group(1)])
        continue

    m = re.search(r"dec (\w+)", line)
    if m:
        cmds.append(['dec', m.group(1)])
        continue

    m = re.search(r"jnz ([-\w]+) ([-\w]+)", line)
    if m:
        cmds.append(['jnz', m.group(1), m.group(2)])
        continue

    m = re.search(r"tgl ([-\w]+)", line)
    if m:
        cmds.append(['tgl', m.group(1)])
        continue

    assert not line

ip = 0

def value(x):
    try:
        return int(x)
    except ValueError:
        return regs[x]

def cpy(x, y):
    if re.match("[a-z]", y):
        regs[y] = value(x)
    return None

def jnz(x, y):
    if value(x) != 0:
        return ip + value(y)
    return None

def inc(x):
    if re.match("[a-z]", x):
        regs[x] += 1
    return None

def dec(x):
    if re.match("[a-z]", x):
        regs[x] -= 1
    return None

def tgl(x):
    tgt = ip + value(x)
    if tgt < 0 or tgt >= len(cmds):
        return

    if len(cmds[tgt]) == 2:
        if cmds[tgt][0] == 'inc':
            cmds[tgt][0] = 'dec'
        else:
            cmds[tgt][0] = 'inc'
    elif cmds[tgt][0] == 'jnz':
        cmds[tgt][0] = 'cpy'
    else:
        cmds[tgt][0] = 'jnz'
    return None

import pprint
regs["a"] = 12
while 0 <= ip < len(cmds):
    #cmds[ip].append('*')
    #pprint.pprint(cmds)
    #cmds[ip].pop()

    cmd = cmds[ip]
    #print ip, cmd

    f = globals()[cmd[0]]
    r = f(*cmd[1:])
    if r is None:
        ip += 1
    else:
        ip = r

    #pprint.pprint(regs)
    #print

    if ip < 0 or ip >= len(cmds):
        break

print regs["a"]
