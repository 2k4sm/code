import java.util.ArrayList;
import java.util.List;

public class ngenerate {
    public static void main(String[] args) {
        // List<List<Integer>> ans = new ArrayList<>();
        // generate(ans, 3, new ArrayList<>());

        int[] A = { 1, 2, 3 };

        ArrayList<ArrayList<Integer>> container = new ArrayList<>();

        ArrayList<Integer> list = new ArrayList<>();

        gensequence(container, list, 0, A);
        for (List<Integer> midans : container) {
            for (int val : midans) {
                System.out.print(val + " ");
            }
            System.out.println();
        }
    }

    public static void generate(List<List<Integer>> ans, int N, ArrayList<Integer> midans) {
        if (midans.size() == N) {
            ans.add((ArrayList<Integer>) midans.clone());
            return;
        }

        midans.add(1);
        generate(ans, N, midans);
        midans.remove(midans.size() - 1);
        midans.add(2);
        generate(ans, N, midans);
        midans.remove(midans.size() - 1);
    }

    public static void gensequence(ArrayList<ArrayList<Integer>> container, ArrayList<Integer> list, int index,
            int[] A) {
        if (index == A.length) {
            container.add((ArrayList<Integer>) list.clone());
            return;
        }

        list.add(A[index]);
        gensequence(container, list, index + 1, A);
        list.remove(list.size() - 1);
        gensequence(container, list, index + 1, A);
    }
}