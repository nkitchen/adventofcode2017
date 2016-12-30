import collections
import fileinput
import itertools
import re
import sys

regs = None

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

    m = re.search(r"out ([-\w]+)", line)
    if m:
       cmds.append(['out', m.group(1)])
       continue

    assert not line

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
        return value(y)
    return None

def inc(x):
    if re.match("[a-z]", x):
        regs[x] += 1
    return None

def dec(x):
    if re.match("[a-z]", x):
        regs[x] -= 1
    return None

def run(a0):
   global regs
   regs = collections.defaultdict(int)
   regs["a"] = a0

   output = []
   ip = 0
   while 0 <= ip < len(cmds):
       cmd = cmds[ip]

       if cmd[0] == "out":
          output.append(value(cmd[1]))
          if len(output) >= 20:
             return output
       else:
          f = globals()[cmd[0]]
          r = f(*cmd[1:])
          if r is None:
              ip += 1
          else:
              ip += r

       if ip < 0 or ip >= len(cmds):
           break

for a0 in itertools.count(1):
   sys.stderr.write(".")
   s20 = list(itertools.islice(run(a0), 20))
   #print s20
   #print [0, 1] * 10
   #print
   if s20 == [0, 1] * 10:
      print a0
      break


