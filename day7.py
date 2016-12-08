import fileinput
import re

def tls(addr):
   for m in re.finditer(r"\[\w*(\w)(\w)\2\1\w*\]", addr):
      if m.group(1) != m.group(2):
         # ABBA in hypernet sequence
         return False

   for m in re.finditer(r"((\w)(\w)\3\2)", addr):
      if m.group(2) != m.group(3):
         # ABBA
         return m.group(1)

   return False

for line in fileinput.input():
    addr = line.strip()
    abba = tls(addr)
    if abba:
        print abba, addr
