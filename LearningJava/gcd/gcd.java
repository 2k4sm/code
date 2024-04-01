import java.util.Scanner;

public class gcd {
  public static void main(String[] args) {
    System.out.println(findGcd(15, 2));
  }

  public static int findGcd(int a, int b) {
    if (a == 0 || b == 0) {
      return a + b;
    }

    if (a == 1 || b == 1) {
      return 1;
    }

    return findGcd(Math.max(a, b) % Math.min(a, b), Math.min(a, b));
  }
}
