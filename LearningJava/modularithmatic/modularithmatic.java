public class modularithmatic {
  public static void main(String[] args) {
    int Num = 100;

    // boolean[] ans = genPrime(Num);

    // for (int i = 0; i <= Num; i++) {
    // if (ans[i] == false) {
    // System.out.println("" + i + " " + ans[i]);
    // }
    // }
    factorsPrime(Num);
  }

  public static boolean isPrime(int number) {
    for (int i = 2; i * i <= number; i++) {
      if (number % i == 0) {
        return false;
      }
    }

    return true;
  }

  public static boolean[] genPrime(int Num) {
    boolean[] primes = new boolean[Num + 1];
    primes[0] = true;
    primes[1] = true;

    for (int i = 2; i * i <= Num; i++) {
      if (primes[i] == false) {
        for (int j = i * i; j <= Num; j += i) {
          primes[j] = true;
        }
      }
    }

    return primes;
  }

  public static void factorsPrime(int Num) {
    int temp = Num;
    for (int i = 2; i * i < temp; i++) {
      if (Num % i == 0) {
        Num /= i;
        System.out.println(i);
        i--;
      }
    }
  }
}
