import java.util.ArrayList;

public class perms {
  static ArrayList<ArrayList<Integer>> permsContainer = new ArrayList<>();

  public static void main(String[] args) {
    int[] testVal = { 1, 1, 2 };
    int[] visited = new int[4];
    ArrayList<Integer> currPerm = new ArrayList<>();
    genPerms(testVal, visited, currPerm);

    for (ArrayList<Integer> conts : permsContainer) {
      for (int i : conts) {
        System.out.print(i + " ");
      }
      System.out.println();
    }

  }

  public static void genPerms(int[] val, int[] visited, ArrayList<Integer> currPerm) {
    if (currPerm.size() == val.length) {
      permsContainer.add(new ArrayList<>(currPerm));
      return;
    }

    for (int i = 0; i < val.length; i++) {
      if (visited[i] == 1 || (i > 0 && val[i] == val[i - 1] && visited[i - 1] == 0)) {
        continue;
      }
      visited[i] = 1;
      currPerm.add(val[i]);
      genPerms(val, visited, currPerm);
      visited[i] = 0;
      currPerm.remove(currPerm.size() - 1);
    }
  }
}
