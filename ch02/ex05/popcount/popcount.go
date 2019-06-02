package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var i uint64
	var r int

	if x == 0 {
		return 0
	}

	if x%2 == 0 {
		return 1
	}
	for i = x; i != 0; i-- {
		if i&(i-1) == 0 {
			r++
		}
	}
	return r
	// var r int
	// for i := 0; i < 64; i++ {
	// 	r += int(x >> uint(i) & 1)
	// 	fmt.Printf("%+v\n", x>>uint(i)&(x>>uint(i)-1))
	// }
	// return 0
	// for i := 0; i < 64; i++ {
	// }
}

func PopCountOld(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
