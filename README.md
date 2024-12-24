# adventofcode2024 #

My [2024 Advent of Code](https://adventofcode.com/2024) puzzle solutions, written in Go. No dependencies on anything outside the standard library. 
Each solution is documented between its source code file and the README for its day. As my friends and I always say, "if it's not documented, it's not done." Excluding README documentation, the `cloc --by-file ./**/*.go` of my source code after Day 05 was 231 blank lines, 209 comment lines, 799 code lines. Those are pretty typical numbers for me, given the accompanying README files.  

The event started on December 1st, and I was actually available to start it on opening day this year.  

#### 2023 bests ####
My best Part 1 rank was #3910 on Day 22. My best Part 2 rank was #4194 on Day 20.  
My fastest time to Part 1 was 00:19:55 on Day 15. My fastest time to Part 2 was 01:16:44 on Day 11. My fastest delta for time 2 was 00:24:26 on Day 11.  

#### 2024 bests ####
My best Part 1 rank was #1062 on Day 20. My best Part 2 rank was #4641 on Day 4.  
My fastest time to Part 1 was 00:25:12 on Day 20. My fastest time to Part 2 was 00:37:54 on Day 4. My fastest delta for time 2 was 00:05:31 on Day 18.  


### Times and Ranks ###
```
      --------Part 1---------   --------Part 2---------
Day       Time    Rank  Score       Time    Rank  Score
 24   01:23:40    4820      0   09:16:49    5104      0   P2 was complicated
 23   00:38:12    3492      0   03:57:43    6356      0   P2 worried too much about graph theory
 22   00:23:38    3172      0   03:36:49    5353      0   P2 Streamline, do less
 21   00:00:00    xxxx      0   00:00:00    xxxx      0   .    
 20   00:25:12    1062      0   03:58:43    5115      0   BfsDijkstra(), Manhattan() born here
 19   00:44:09    4428      0   01:25:33    4775      0   P2 Memoization saves the day again    
 18   01:19:51    5970      0   01:25:22    5413      0   Dijkstra's algorithm
 17   01:22:25    5612      0   05:38:18    4661      0   P2 octal arithmetic nightmare
 16   23:46:52   23316      0   00:00:00   xxxxx      0   Weighted mazes, manually eased P1    
 15   01:12:45    5570      0   12:07:02   13841      0   Tired after P1, went to bed    
 14   00:45:22    5062      0   04:27:26    9554      0   P2 useless hint without example
 13   01:34:12    7820      0   02:06:26    5716      0   regexp.FindSubmatch() P1 painful
 12   01:47:15    8942      0   03:12:04    6149      0   Ternary If()
 11   00:45:02    8952      0   01:26:13    6302      0   P2 Memoization example :)
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
day03/                  RegEx.FindAllString/FindAllStringIndex day
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
day11/                  Memoization of P2
day12/                  Routes counting
  day12/embed1.txt        // provided examples
  day12/embed2.txt        
  day12/mini1.txt         
  day12/mini2.txt         
  day12/abba2.txt         
  day12/eee2.txt          
day13/                  RegEx.FindSubmatch day and easy algebra for diophantine equations 
day14/                  Visualizer, wrote ModSolver(d1,r1,d2,r2) will add to utilities
  day14/AoC2024_Day14.png // screengrab of visualization
day15/                  Pushing boxes, sometimes pairs, collision detection
day16/                  Weighted mazes, least cost
  day16/maze1.txt         // 2nd provided example
  day16/simplified1.txt   // Manually simplified maze, based off of input1.txt
  day16b{ex,maze,input,simplified} // ASCII-visualized mazes
  day16/D16P1.png         // screengrab of bsimplified visualization that solved correctly
day17/                  Computer sim, then Octal suffix jumps for P2
  day17/two2.txt          // 2nd example provided
  day17/AoC_octalBruteForceDigits.png // screengrab from related YouTube video
  day17/d17_aoc2024.py    // my fork of JP's approach, came close and inspired
  day17/mrphlip.py        // my (barely) fork of MP's approach, did not solve
  day17/hneut.py          // HN's approach, solved it, partial inspiration
day18/                  Dijkstra's algorithm day
day19/                  Memoization of P2 again
day20/                  Dijkstra again, then BfsDijkstra(..., ignoreWalls, maxDist)
day21/                  .
day22/                  Score sequences of numeric diffs
  day22/given2.txt        // example for P2
  day22/p2brute.go        // first approach for P2 using brute force, not sure why wrong
day23/                  K-complete subgraphs  
day24/                  Logic gates validation
  day24/mini1.txt         // example for P1
  day24/p2first.go        // first approach for P2
  day24/der.txt           // derivations
  day24/depends.txt       // der, values only
  day24/maxdegs.txt       // der, max degree only
  day24/complexity.txt    // der, "[" only
  day24/permits.txt       // data structure when exploring P1
  day24/zGates.txt        // just the z-prefixed Gates
  day24/manual2.txt       // bonus manual solution, see daily README
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

Created a pait of bash functions to quicken my daily setup, in `~/.bash_aliases`:
```
#!/bin/bash

# already had the mdg() function to make and change working directory
mdg ()
{
    mkdir -p "$1"
    cd "$1"
}

# Create a new Advent of Code day$1 directory using suitable templates
newaocday ()
{
  newdaynum="$1"
  # move to root of repo
  cd $(git rev-parse --show-toplevel)
  if [ ! -d "day$newdaynum" ]; then
    mdg "day$newdaynum"
    touch ex1.txt ex2.txt input1.txt input2.txt problem2.go
    cp ../probTemplate.go problem1.go
    cp $HOME/Templates/aoc2024README.md README.md
    sed -i "s/XXX/${1}/g" README.md
    go mod init this
    go mod tidy
  fi
}

# below is accurate during December of 2024
# in Firefox, Ctrl+Shift+I = developer tools, then Storage tab->Cookies->session copy
export AOC_Cookie="12345678...90abcdef"

aocinput ()
{
  dirda=$(pwd)
  # get last 2 characters of current dir
  da=${dirda: -2}
  yr=$(date '+%Y')
  #https://adventofcode.com/2024/day/19/input
  infile="https://adventofcode.com/${yr}/day/${da}/input"
  curl "${infile}" -H "cookie: session=${AOC_Cookie}" -o input1.txt
  cp input1.txt input2.txt
}
```
I run `newaocday 07` to create December 7th's directory, any time on or before that date.  

Once midnight hits, I can run `aocinput` to pull down my copy of the input, using the session cookie I have stored in `AOC_Cookie` environment variable, that I previously cached from Firefox.  
