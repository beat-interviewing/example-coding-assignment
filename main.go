func FibFn() func() uint64 {
    a, b := uint64(0), uint64(1)
	return func() uint64 {
		a, b = b, a+b
		return a
	}
}

func main() {
	fib := FibFn()
	for j := 0; j < 10; j++ {
		fmt.Printf("%v ", fib())
	}
}