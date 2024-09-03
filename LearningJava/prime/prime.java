import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class prime {
    public static void main(String[] args) {
        Scanner scan = new Scanner(System.in);

        int n = scan.nextInt();

        scan.close();
        List<Integer> list = new ArrayList<>();

        findPrime(n, list);

        for (int i = 0; i < list.size(); i++) {
            System.out.println(list.get(i) + " ");
        }
    }

    public static void findPrime(int n, List<Integer> list) {
        for (int i = 1; i * i * 1L <= n; i++) {
            if (n % i == 0) {
                int f1 = 1;
                int f2 = n / i;

                list.add(f1);
                if (f1 != f2) {
                    list.add(f2);
                }
            }
        }
    }
}