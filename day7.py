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

abaRe = re.compile( r"((\w)\w\2)" )

def ssl(addr):
   supers = re.findall( r"\w+(?=\[|$)", addr )
   hypers = re.findall( r"(?<=\[)\w+(?=\])", addr )

   for s in supers:
      def _hits():
         i = 0
         while i < len(s):
            m = abaRe.search( s, i )
            if m is None:
               return
            yield m.group(0)
            i = m.start() + 1

      for aba in _hits():
         if aba[0] == aba[1]:
            continue

         bab = aba[1] + aba[0] + aba[1]
         for h in hypers:
            if bab in h:
               return True

   return False

for line in fileinput.input():
    addr = line.strip()
    if ssl(addr):
        print addr
