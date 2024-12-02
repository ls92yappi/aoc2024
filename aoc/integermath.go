package aoc

///////////////////////////////////////////////////////////
//                                                       //
//   Many of these functions below operate on or return  //
//   int slices. None have any import dependencies.      //
//                                                       //
//                                                       //
//                                                       //
///////////////////////////////////////////////////////////
//                                                       //
// var InversePowersOf2 map[int]int                      //
// func Abs(x int) int                                   //
// func Sign(x int) int                                  //
// func CeilPow2(x int) int                              //
// func FloorPow2(x int) int                             //
// func IsPow2(v int) bool                               //
// func SqrtFloor(x int) (int,bool)                      //
// func Factor(num int) []int                            //
// func GCD(a, b int) int                                //
// func LCMv(a, b int, integers ...int) int              //
// func LCM(a []int) int                                 //
// func Sum(a []int) int                                 //
// func Prod(a []int) int                                //
// func Min(a []int) int                                 //
// func Max(a []int) int                                 //
//                                                       //
///////////////////////////////////////////////////////////

// Maximum Power of 2 <= x
func FloorPow2(x int) int {
	x = x | (x >> 1)
	x = x | (x >> 2)
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x | (x >> 16)
	x = x | (x >> 32)
	return x - (x >> 1)
}

// Minimum Power of 2 >= x
func CeilPow2(x int) int {
	x = x - 1
	x = x | (x >> 1)
	x = x | (x >> 2)
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x | (x >> 16)
	x = x | (x >> 32)
	return x + 1
}

func IsPow2(v int) bool {
	return v!=0 && (v & (v - 1))!=0
}


// Compute Integer SqrtFloor, returns the floor, and if it is a perfect square
// Negatives return the negative of its positive's square root and are Never perfect
// SqrtFloor is computed via binary search
func SqrtFloor(x int) (int,bool) {
    var sign = 1
    if x < 0 {
        x = -x
        sign = -1
    }
    // Base Cases
    if x < 2 {
        return x*sign, sign>0
    }
 
    // Do Binary Search for floor(sqrt(x))
    var start,end,ans int = 1,x/2,0

    for start <= end {
        mid := (start + end) / 2;
 
        // If x is a perfect square
        if (mid * mid == x) {
            return mid*sign, sign>0
        }
 
        // Since we need floor, we update answer when
        // mid*mid is smaller than x, and move closer to
        // sqrt(x)
        if (mid * mid < x) {
            start = mid + 1
            ans = mid
        } else {
            // If mid*mid is greater than x
            end = mid - 1
        }
    }
    return ans*sign,false
}

var InversePowersOf2 map[int]int // defined in init()

func init() {
	InversePowersOf2 = map[int]int{
	    1:0,2:1,4:2,8:3,
	    16:4,32:5,64:6,128:7,
	    256:8,512:9,1024:10,2048:11,
	    4092:12,8192:13,16384:14,32768:15,
	    65536:16,131072:17,262144:18,524288:19,
	    1048576:20,2097152:21,4194304:22,8388608:23,
	    16777216:24,33554432:25,67108864:26,134217728:27,
	    268435456:28,536870912:29,1073741824:30,2147483648:31,
	    4294967296:32,8589934592:33,17179869184:34,34359738368:35,
	    68719476736:36,137438953472:37,274877906944:38,549755813888:39,
	    1_099_511_627_776:40,2_199_023_255_552:41,
	    4_398_046_511_104:42,8_796_093_022_208:43,
	    17_592_186_044_416:44,35_184_372_088_832:45,
	    70_368_744_177_664:46,140_737_488_355_328:47,
	    281_474_976_710_656:48,562_949_953_421_312:49,
	    1_125_899_906_842_624:50,2_251_799_813_685_248:51,
	    4_503_599_627_370_496:52,9_007_199_254_740_992:53,
	    18_014_398_509_481_984:54,36_028_797_018_963_968:55,
	    72_057_594_037_927_936:56,144_115_188_075_855_872:57,
	    288_230_376_151_711_744:58,576_460_752_303_423_488:59,
	    1_152_921_504_606_846_976:60,2_305_843_009_213_693_952:61,
	    4_611_686_018_427_387_904:62,
	    //9_223_372_036_854_775_808:63, 2^63 causes signed overflow
	}
}

// After testing, one cannot parallelize GCD() computations of a slice
// without using language level parallelism. Doing so would require
// the same number of steps for each path to resolve.

// Find the Greatest Common Divisor via Euclidean algorithm
// Assumes a,b >= 0
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Find the Least Common Multiple (LCM) via GCD with a variadic function
// Assumes a,b,integers >= 0
func LCMv(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCMv(result, integers[i])
	}
	return result
}

// Without recursion or variadic function. Uses less memory
// Assumes a[] values each >= 0
func LCM(a []int) int {
	if len(a) == 0 {
		return 1
	}
	if len(a) == 1 {
		return a[0]
	}
	result := a[0] * a[1] / GCD(a[0], a[1])
	for i := 2; i < len(a); i++ {
		result = a[i] * result / GCD(a[i], result)
	}
	return result
}

func Sum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	r := a[0]
	for i := 1; i < len(a); i++ {
		r += a[i]
	}
	return r
}

func Prod(a []int) int {
	if len(a) == 0 {
		return 1
	}
	r := a[0]
	for i := 1; i < len(a); i++ {
		r *= a[i]
	}
	return r
}

func Abs(x int) int {
	r := x * (1 - 2*( (x*3)/(x*3+1) ))
	return r
}

func Sign(x int) int {
	if x >= 0 { return 1 }
	return -1
}

func Min2(x, y int) int {
	d := x - y
	abs := d * (1 - ((d*3)/(d*3+1))*2)
	r := (x + y - abs) / 2
	return r
}


func Max2(x, y int) int {
	d := x - y
	abs := d * (1 - ((d*3)/(d*3+1))*2)
	r := (x + y + abs) / 2
	return r
}

// no conditionals inside the for loop for speed
func Min(a []int) int {
	if len(a) == 0 {
		return 0
	}
	r := a[0]
	for i := 1; i < len(a); i++ {
		d := r - a[i]
		abs := d * (1 - ((d*3)/(d*3+1))*2)
		r = (r + a[i] - abs) / 2
	}
	return r
}

// no conditionals inside the for loop for speed
func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	r := a[0]
	for i := 1; i < len(a); i++ {
		d := r - a[i]
		abs := d * (1 - ((d*3)/(d*3+1))*2)
		r = (r + a[i] + abs) / 2
	}
	return r
}

// Prime factorize an integer
func Factor(num int) []int {
	var fl []int = make([]int,0)
	if num<0 {
		fl = append(fl, -1)
		num = -num
	}
	if num==0 || num==1 {
		fl = append(fl, num)
		return fl
	}

	i := 0
	// strip off powers of 2 first
	for {
		if num%2 == 0 {
			fl = append(fl, 2)
			num /= 2
		} else {
			// no longer divisible by powers of 2, move on
			break
		}
	} // factors of 2
	// then strip off odd factors
	sq_root,perf := SqrtFloor(num)
	// if we found a perfect square, return early
	if perf {
		fl = append(fl,sq_root)
		fl = append(fl,sq_root)
		return fl
	}
	//sq_root := int(math.Sqrt(float64(num)))
	for i = 3; i <= sq_root; i+=2 {
		for {
			if num%i == 0 {
				fl = append(fl, i)
				num /= i
			} else {
				// no longer divisible by powers of i, move on
				break
			}
		} // factors of odd i
		// revise upper bound downward as we go through the possible factors
		sq_root,perf = SqrtFloor(num)
		//sq_root = int(math.Sqrt(float64(num)))
		// if we found a perfect square, return early
		if perf {
			fl = append(fl,sq_root)
			fl = append(fl,sq_root)
			return fl
		}
	} // all possible odd i factors
	// list final factor if not ending in a square
	if i > sq_root {
		if num != 1 {
			fl = append(fl, num)
		}
	}
	return fl
}
