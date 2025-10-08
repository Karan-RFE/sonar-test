package main

func IsNumberDivisibleBy1000(num int, divisor int) (result bool) {
	if divisor == 0 {
		return false
	}
	return num%divisor == 0
}
