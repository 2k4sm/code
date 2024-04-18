public class frog {
  public static void main(String[] args) {
    int[] heights = { 10, 20, 30, 10 };
    int[] memo = new int[heights.length + 1];
    System.out.println(climb(heights, heights.length - 1, memo));
  }

  public static int climb(int[] heights, int index, int[] memo) {
    if (index == 0) {
      return 0;
    }

    if (memo[index] != 0) {
      return memo[index];
    }

    int takeOne = climb(heights, index - 1, memo) + Math.abs(heights[index] - heights[index - 1]);
    int takeTwo = Integer.MAX_VALUE;

    if (index > 1) {
      takeTwo = climb(heights, index - 2, memo) + Math.abs(heights[index] - heights[index - 2]);
    }

    return memo[index] = Math.min(takeOne, takeTwo);
  }

}
