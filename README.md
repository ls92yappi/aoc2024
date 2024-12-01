# adventofcode2024 #

My [2024 Advent of Code](https://adventofcode.com/2024) puzzle solutions, written in Go. No dependencies on anything outside the standard library. 
Each solution is documented between its source code file and the README for its day. As my friends and I always say, "if it's not documented, it's not done." Excluding README documentation, the `cloc --by-file ./**/*.go` of my source code after Day 01 was XXXX blank lines, XXXX comment lines, XXXX code lines. Those are pretty typical numbers for me, given the accompanying README files.  

The event started on December 1st, and I was actually available to start it on opening day this year.  


### Times and Ranks ###
```
      --------Part 1---------   --------Part 2---------
Day       Time    Rank  Score       Time    Rank  Score
  2   00:00:00              0   00:00:00              0
  1   00:00:00   98765      0   00:00:00   43210      0
```

### Brief overview of all files ###
```
README.md
integermath.go           // Convenience only - written in 2023
day01/
  day01/problem01.go       // D1 P1
  day01/problem02.go       // D1 P2
  day01/ex01.txt
  day01/ex02.txt
  day01/input01.txt
  day01/input02.txt
  day01/README.md
```


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
