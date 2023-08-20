// You are given an integer A. Your task is to find an integer X, such that:

//     Number of prime factors in the prime factorization of A and X are the same.
//     Prime factorization of A and X differs by exactly one prime factor.

// For ex: if A = 12 (2*2*3), then possible value of X can be 18 (3*2*3), 30(2*5*3), 20(2*2*5) etc.

// Among all possible values of X, return the one which has minimum possible value of abs(A-X). If there are multiple values of X possible, return the largest one.

package nsimnum

func Nsimilar(A int) int {
	AfactorList := Factors(A)

	//Run a loop till A*2 and store all the possible numbers. Then check the numbers which have three factors and return the nearest one.

}

func Factors(A int) (fs []int) {
	//find prime factor by reccursively dividing with all the prime numbers till zero.

}
