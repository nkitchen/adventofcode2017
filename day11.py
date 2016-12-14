#!/usr/bin/env python
import collections
import fileinput
import heapq
import itertools
import re
from pprint import pprint

State = collections.namedtuple(
    "State",
    "floorContents elevatorFloor stepsSoFar stepsToGo")

floorIndex = dict(
    first=1,
    second=2,
    third=3,
    fourth=4,
)

floorContents = [[] for i in range(5)]
for line in fileinput.input():
    m = re.search(r"The (\S+) floor contains", line)
    if not m:
        continue

    f = floorIndex[m.group(1)]

    for elem in re.findall(r"a (\S+)-compatible microchip", line):
        floorContents[f].append((elem, 'chip'))

    for elem in re.findall(r"a (\S+) generator", line):
        floorContents[f].append((elem, 'gen'))

def estimatedStepsRemaining(floorContents):
    # Heuristic: a direct ride to the assembler for each pair of items,
    # with no constraints
    n = 0
    for i, floor in enumerate(floorContents):
        d = 4 - i
        n += len(floor) / 2.0 * d
    return n

def itemsAreSafe(items):
    chips = set(x[0] for x in items if x[1] == 'chip')
    gens = set(x[0] for x in items if x[1] == 'gen')
    fried = len(chips - gens) and len(gens)
    return not fried

floorContents = tuple(map(tuple, floorContents))
s0 = State(floorContents=floorContents, elevatorFloor=1, stepsSoFar=0,
           stepsToGo=estimatedStepsRemaining(floorContents))

queue = [(s0.stepsToGo, s0)]
visited = set()

t = 0
while queue:
    t += 1
    print t, ':', len(queue)

    _, s = heapq.heappop(queue)
    pprint(tuple(s))

    if s.stepsToGo == 0:
        print s.stepsSoFar
        break

    visited.add((s.floorContents, s.elevatorFloor))

    nextFloors = []
    if s.elevatorFloor > 1:
        nextFloors.append(s.elevatorFloor - 1)
    if s.elevatorFloor < 4:
        nextFloors.append(s.elevatorFloor + 1)

    curItems = set(s.floorContents[s.elevatorFloor])
    for carried in itertools.imap(set, itertools.chain(
        itertools.combinations(curItems, 1),
        itertools.combinations(curItems, 2),
    )):
        itemsLeft = curItems - carried
        if not itemsAreSafe(itemsLeft):
            continue

        for f in nextFloors:
            nextItems = set(s.floorContents[f]) | carried
            if not itemsAreSafe(nextItems):
                continue

            floorContents = list(s.floorContents)
            floorContents[s.elevatorFloor] = tuple(itemsLeft)
            floorContents[f] = tuple(nextItems)
            floorContents = tuple(floorContents)

            if (floorContents, f) in visited:
                continue

            sn = State(floorContents=floorContents,
                       elevatorFloor=f,
                       stepsSoFar=1 + s.stepsSoFar,
                       stepsToGo=estimatedStepsRemaining(floorContents))
            d = sn.stepsSoFar + s.stepsToGo
            heapq.heappush(queue, (d, sn))

