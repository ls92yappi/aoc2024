# adventofcode2024 #

My [2024 Advent of Code](https://adventofcode.com/2024) puzzle solutions, written in Go. No dependencies on anything outside the standard library. 
Each solution is documented between its source code file and the README for its day. As my friends and I always say, "if it's not documented, it's not done." Excluding README documentation, the `cloc --by-file ./**/*.go` of my source code after Day 02 was 136 blank lines, 107 comment lines, 385 code lines. Those are pretty typical numbers for me, given the accompanying README files.  

The event started on December 1st, and I was actually available to start it on opening day this year.  

#### 2023 bests ####
My best Part 1 rank was #3910 on Day 22. My best Part 2 rank was #4194 on Day 20.  
My fastest time to Part 1 was 00:19:55 on Day 15. My fastest time to Part 2 was 01:16:44 on Day 11. My fastest delta for time 2 was 00:24:26 on Day 11.  

#### 2024 bests ####
My best Part 1 rank was #7371 on Day 1. My best Part 2 rank was #8314 on Day 4.  
My fastest time to Part 1 was 00:25:50 on Day 1. My fastest time to Part 2 was 00:37:54 on Day 1. My fastest delta for time 2 was 00:12:17 on Day 4.  


### Times and Ranks ###
```
      --------Part 1---------   --------Part 2---------
Day       Time    Rank  Score       Time    Rank  Score
  4   00:25:37    5298      0   00:37:54    4641      0   
  3   00:28:02   10234      0   01:09:23   11865      0   
  2   00:31:21    9795      0   02:34:52   16100      0   Worked through a variety of off-by-1 bugs
  1   00:25:50    7371      0   00:46:26    8314      0   See https://adventofcode.com/2024/leaderboard/self
```


### Brief overview of all files ###
```
README.md
integermath.go           // refactored into ls92yappi/aoc repo
day01/
  day01/problem1.go       // D1 P1
  day01/problem2.go       // D1 P2
  day01/ex1.txt
  day01/ex2.txt
  day01/input1.txt
  day01/input2.txt
  day01/README.md
day02/
  day02/problem1.go       // D2 P1
  day02/problem2.go       // D2 P2
  day02/ex1.txt
  day02/ex2.txt
  day02/input1.txt
  day02/input2.txt
  day02/test2.txt         // BAD results from P1 for further testing
  day02/README.md
day03/
  day03/problem1.go       // D3 P1
  day03/problem2.go       // D3 P2
  day03/ex1.txt
  day03/ex2.txt
  day03/input1.txt
  day03/input2.txt
  day03/README.md
day04/
  day04/problem1.go       // D4 P1
  day04/problem2.go       // D4 P2
  day04/ex1.txt
  day04/ex2.txt
  day04/input1.txt
  day04/input2.txt
  day04/README.md
```


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
