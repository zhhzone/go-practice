package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])

}

func PopCount2(x uint64) int {
	var sum byte

	for i := uint(0); i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}
	return int(sum)
}

func PopCount3(x uint64) int {
	sum := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			sum++
		}
	}

	return sum
}

func PopCount4(x uint64) int {
	sum := 0
	for x != 0 {
		x = x & (x - 1)
		sum++
	}

	return sum
}
