package helper

func Multiply(num1, num2 int) int {
	return num1 * num2
}

func Sum(num ...int) (total int) {
	if len(num) == 0 {
		return 0
	}

	if len(num) <= 1 {
		return num[0]
	}

	for _, n := range num {
		total += n
	}

	return
}

func IsPrime(num int) (isPrime bool) {
	counter := 0
	for i := 2; i < num; i++ {
		if num%i == 0 {
			counter++
		}
	}
	return counter == 0
}
