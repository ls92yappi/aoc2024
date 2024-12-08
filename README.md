# adventofcode2024 #

My [2024 Advent of Code](https://adventofcode.com/2024) puzzle solutions, written in Go. No dependencies on anything outside the standard library. 
Each solution is documented between its source code file and the README for its day. As my friends and I always say, "if it's not documented, it's not done." Excluding README documentation, the `cloc --by-file ./**/*.go` of my source code after Day 05 was 231 blank lines, 209 comment lines, 799 code lines. Those are pretty typical numbers for me, given the accompanying README files.  

The event started on December 1st, and I was actually available to start it on opening day this year.  

#### 2023 bests ####
My best Part 1 rank was #3910 on Day 22. My best Part 2 rank was #4194 on Day 20.  
My fastest time to Part 1 was 00:19:55 on Day 15. My fastest time to Part 2 was 01:16:44 on Day 11. My fastest delta for time 2 was 00:24:26 on Day 11.  

#### 2024 bests ####
My best Part 1 rank was #5298 on Day 4. My best Part 2 rank was #4641 on Day 4.  
My fastest time to Part 1 was 00:25:37 on Day 4. My fastest time to Part 2 was 00:37:54 on Day 4. My fastest delta for time 2 was 00:08:34 on Day 10.  


### Times and Ranks ###
```
      --------Part 1---------   --------Part 2---------
Day       Time    Rank  Score       Time    Rank  Score
 10   01:40:34   10134      0   01:48:58    9811      0   Fiddly P1 setup
  9   03:57:49   17179      0   05:36:29   12916      0   bytes -> []int, P2 -> []struct{int,int}
  8   00:47:23    6171      0   00:59:23    5662      0   
  7   00:26:20    4786      0   02:48:48   12230      0   P2 Recursive ops
  6   00:47:00    8513      0   01:33:29    6052      0   
  5   00:30:38    6529      0   01:54:49   11514      0   P2 off-by-1: append(fixed,orig[i+1:j]
  4   00:25:37    5298      0   00:37:54    4641      0   
  3   00:28:02   10234      0   01:09:23   11865      0   
  2   00:31:21    9795      0   02:34:52   16100      0   Worked through a variety of off-by-1 bugs
  1   00:25:50    7371      0   00:46:26    8314      0   
```
See https://adventofcode.com/2024/leaderboard/self  


### Brief overview of all files ###
```
README.md
integermath.go           // refactored into ls92yappi/aoc repo
dayNN/
  dayNN/problem1.go       // Dn P1
  dayNN/problem2.go       // Dn P2
  dayNN/ex1.txt
  dayNN/ex2.txt
  dayNN/input1.txt
  dayNN/input2.txt
  dayNN/README.md
day01/                  List (skipped frequency table)
day02/                  Ranged increasing/decreasing filter with removal
  day02/test2.txt         // BAD results from P1 for further testing
day03/                  RegEx day
day04/                  Crossword
day05/                  Insertion sort with slices (skipped toposort)
  day05/test2.txt         // Simplified known bad sequence from P2
day06/                  Maze Day (faked loop detection with timing out)
day07/                  Recursive operation varying and shortcircuiting
day08/                  Hash table and grid day
day09/                  Defrag modeling
  day09/row1.txt          // One row long from input, included 2-digit values
  day09/short2.txt        // Case where first useful Gap after values to move
day10/                  Routes counting
  day10/test1.txt         // Simplest example provided
```
Note that `ex1.txt` and `ex2.txt` along with `input1.txt` and `input2.txt` are redundant.  


Simple helper script to speed up runs at `~/bin/solve`:
```bash
#!/bin/bash

# solve $1 [$2] -- runs `go run problem$1.go input$1.txt` or `go run problem$1.go $2$1.txt`
# examples:
#   solve 2
#   solve 1 ex

if [ -z $2 ]; then
  go run problem$1.go input$1.txt
else
  go run problem$1.go $2$1.txt
fi
```
