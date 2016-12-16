import collections
import fileinput
import re

cmds = [line.strip() for line in fileinput.input()]

regs = collections.defaultdict(int)
regs['c'] = 1

ip = 0

def value(x):
    try:
        return int(x)
    except ValueError:
        return regs[x]

while 0 <= ip < len(cmds):
    cmd = cmds[ip]
    #print ip, cmd

    p = (r"(?P<cpy>cpy (?P<xCpy>[-\w]+) (?P<yCpy>\w+))|"
         r"(?P<inc>inc (?P<xInc>\w+))|"
         r"(?P<dec>dec (?P<xDec>\w+))|"
         r"(?P<jnz>jnz (?P<xJnz>[-\w]+) (?P<yJnz>[-\w]+))" )
    m = re.match(p, cmd)
    if m.group('jnz'):
        x = value(m.group('xJnz'))
        y = value(m.group('yJnz'))
        if x != 0:
            ip += y
            continue
    elif m.group('cpy'):
        x = value(m.group('xCpy'))
        y = m.group('yCpy')
        regs[y] = x
    elif m.group('inc'):
        x = m.group('xInc')
        regs[x] += 1
    elif m.group('dec'):
        x = m.group('xDec')
        regs[x] -= 1
    ip += 1

print regs["a"]
