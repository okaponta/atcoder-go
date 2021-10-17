package util

func janken(a, b byte) int {
	if a == b {
		return 0
	}
	if a == 'G' && b == 'C' || a == 'C' && b == 'P' || a == 'P' && b == 'G' {
		return 1
	}
	return -1
}
