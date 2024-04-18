/**
 * houseRobber
 */
public class houseRobber {

  public static void main(String[] args) {
    int[] valueInHouses = { 1, 3, 2, 1 };
    int[] memo = new int[valueInHouses.length + 1];

    System.out.println(robbed(valueInHouses, valueInHouses.length, memo));
  }

  public static int robbed(int[] valueInHouse, int index, int[] memo) {

  }
}
