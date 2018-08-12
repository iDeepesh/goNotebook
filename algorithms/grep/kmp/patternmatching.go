package kmp

//GetPartialFunction - returns the partial/failure function table for given pattern string
func GetPartialFunction(p string) []int {
	pf := make([]int, len(p))

	for i := 0; i < len(p); i++ {
		pf[i] = i + 1
	}

	return pf
}
