package gcd

//FindGCD - uses euclid's algorothm
func FindGCD(g, s int) int {
	if s > g {
		s, g = g, s
	}

	r := g % s
	for r != 0 {
		s, g = r, s
		r = g % s
	}

	return s
}
