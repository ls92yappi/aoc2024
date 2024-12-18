# based off of https://github.com/jonathanpaulson/AdventOfCode/blob/master/2024/17.py

import sys
import re
import heapq
from collections import defaultdict, Counter, deque
def pr(s):
    print(s)
    pc.copy(s)
sys.setrecursionlimit(10**6)
DIRS = [(-1,0),(0,1),(1,0),(0,-1)] # up right down left
def ints(s):
    return [int(x) for x in re.findall('-?\d+', s)]

infile = sys.argv[1] if len(sys.argv)>=2 else 'input2.txt'
ans = 0
D = open(infile).read().strip()

regs, program = D.split('\n\n')
A,B,C = ints(regs)
program = program.split(':')[1].strip().split(',')
program = [int(x) for x in program]
#print(A,B,C,program)


def run(Ast, part2):
    def getCombo(x):
        if x in [0,1,2,3]:
            return x
        if x==4:
            return A
        if x==5:
            return B
        if x==6:
            return C
        return -1
    A = Ast
    B = 0
    C = 0
    ip = 0
    out = []
    while True:
        if ip>=len(program):
            return out
        cmd = program[ip]
        op = program[ip+1]
        combo = getCombo(op)

        #print(ip, len(program), cmd)
        if cmd == 0:
            A = A // 2**combo
            ip += 2
        elif cmd == 1:
            B = B ^ op
            ip += 2
        elif cmd == 2:
            B = combo%8
            ip += 2
        elif cmd == 3:
            if A != 0:
                ip = op
            else:
                ip += 2
        elif cmd == 4:
            B = B ^ C
            ip += 2
        elif cmd == 5:
            out.append(int(combo%8))
            if part2 and out[len(out)-1] != program[len(out)-1]:
                return out
            ip += 2
        elif cmd == 6:
            B = A // 2**combo
            ip += 2
        elif cmd == 7:
            C = A // 2**combo
            ip += 2

#part1 = run(A, False)
#print(','.join([str(x) for x in part1]))

#A = Ast
#0o1 1 0 16
#0o14 12 1 16
#0o1277 703 2 16
#0o14040 6176 ...
#0o15240 6816 4 16
#0o1015240 268960 5 16
#0o3015240 793248 6 16

#A = Ast * 8**1 + 0o5
#0o15 13 0 16
#0o105 69 1 16
#0o2165 1141 2 16
#0o7155 3693 3 16
#0o47155 20077 4 16
#0o4057155 1072749 5 16
#0o4257155 1138285 6 16

#A = Ast * 8**4 + 0o7155
#0o1404257155 202464877 7 16
#0o1414257155 204562029 8 16

#A = Ast * 8**6 + 0o257155
#0o31574257155 3455147629 9 16
#0o35574257155 3992018541 10 16

#A = Ast * 8**8 + 0o74257155
#0o2055474257155 143561678445 11 16
#0o2257474257155 161009983085 12 16
#0o22257474257155 1260521610861 13 16

#A = Ast * 8**11 + 0o57474257155
#0o222257474257155 10056614633069 14 16
#0o2562257474257155 95818521599597 15 16
#236556009954925 -- submitted, Too High

#A = Ast * 8**9 + 0o474257155
# dialing back the number of Octal digits from 11 to 9 gave
# me the same thresholds for new digits and the same result
# just at a slower convergence rate

#A = Ast * 8**1 + 0o7
#     0o17 15 0 16
#     0o67 55 1 16
#   0o1277 703 2 16
#  0o15277 6847 4 16
#0o1015277 268991 5 16
#0o3015277 793279 6 16

#A = Ast * 8**3 + 0o277
#  0o127661277 23028415 7 16
# 0o1414257277 204562111 8 16
#0o31574257277 3455147711 9 16
#0o35574257277 3992018623 10 16

#A = Ast * 8**6 + 0o257277
#  0o31574257277 3455147711 9 16
#  0o35574257277 3992018623 10 16
#0o2055414257277 143549095615 11 16
#0o2257414257277 160997400255 12 16

#A = Ast * 8**8 + 0o14257277
#  0o2055414257277 143549095615 11 16
#  0o2257414257277 160997400255 12 16
# 0o22257414257277 1260509028031 13 16
#0o222257414257277 10056602050239 14 16

#A = Ast * 8**9 + 0o414257277
# 0o222257414257277 10056602050239 14 16
#0o2562257414257277 95818509016767 15 16
#0o2562257414257277 95818509016767 15 16
#236555997372095

# soln = 236556009954925 // Incorrect - Too High0o155 route
#        236555997372095 // Also      - Too High 0o277 route
#        236555997372013 // Correct from HyperNeutrino
# diff = 12582830 = 0o57777656, but 0o6 route got me nowhere

# This problem is more an art than it is a science
Ast = 0
best = 0
while True:
    Ast += 1
    #A = Ast
    # trying 0o240 didn't get me past 6
    #A = Ast * 8**1 + 0o7
    #A = Ast * 8**3 + 0o277
    #A = Ast * 8**6 + 0o257277
    #A = Ast * 8**8 + 0o14257277
    #A = Ast * 8**9 + 0o414257277
    A = Ast * 8**10 + 0o7414257277
    #A = Ast * 8**4 + 0o7155
    #A = Ast * 8**6 + 0o257155
    #A = Ast * 8**8 + 0o74257155
    #A = Ast * 8**9 + 0o474257155
    #A = Ast * 8**11 + 0o57474257155
    out = run(A, True)
    if out == program:
        print(A)
        break
    elif len(out) > best:
        print(oct(A), A, best, len(program))
        best = len(out)