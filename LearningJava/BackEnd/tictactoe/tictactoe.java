import java.util.Scanner;

public class tictactoe {
  public static void main(String[] args) {
    int n = 3;
    char[][] board = new char[n][n];

    Scanner scan = new Scanner(System.in);

    System.out.println("Please Enter Player 1 Name:- ");
    String p1 = scan.nextLine();

    System.out.println("Please Enter Player 2 Name:- ");
    String p2 = scan.nextLine();

    // True: Player 1;
    boolean turn = true;

    for (int i = 0; i < n; i++) {
      for (int j = 0; j < n; j++) {
        board[i][j] = '-';
      }
    }

    System.out.println("Player 1, Select your Symbol: ");
    String s1 = scan.next();
    char sym1 = s1.charAt(0);
    System.out.println("Player 2, Select your Symbol: ");
    String s2 = scan.next();
    char sym2 = s2.charAt(0);
    boolean win = false;
    char symbol = sym1;
    while (!win) {
      int count = 0;
      while (true) {
        printBoard(board, n);
        if (turn) {
          System.out.println("Player 1, please make your turn: ");
          symbol = sym1;
        } else {
          System.out.println("Player 2, please make your turn: ");
          symbol = sym2;
        }

        System.out.println("Please input the row: ");
        int row = scan.nextInt();

        System.out.println("Please input the column: ");
        int col = scan.nextInt();

        if (row < 0 || row >= n || col < 0 || col >= n) {
          System.out.println("Invalid Input");
        } else if (board[row][col] != '-') {
          System.out.println("Cell is Blocked");
        } else {
          board[row][col] = symbol;
          count++;
        }

        if (count == n * n) {
          printBoard(board, n);
          System.out.println("Draw");
          win = true;
          break;
        }

        if (checkWinner(board, symbol)) {
          System.out.println(symbol + "Wins");
          win = true;
          break;
        }

        turn = !turn;
      }
    }
  }

  public static void printBoard(char[][] board, int n) {
    for (int i = 0; i < n; i++) {
      for (int j = 0; j < n; j++) {
        System.out.print(board[i][j] + " ");
      }
      System.out.println();
    }
  }

  public static boolean checkWinner(char[][] board, char sym) {
    for (int i = 0; i < board.length; i++) {
      if (board[i][0] == sym && board[i][0] == board[i][1] &&
          board[i][1] == board[i][2]) {
        return true;
      } else if (board[0][i] == sym && board[0][i] == board[1][i] &&
          board[1][i] == board[2][i]) {
        return true;
      } else if ((board[2][0] == sym && board[1][1] == sym &&
          board[0][2] == sym) ||
          (board[0][0] == sym && board[1][1] == sym &&
              board[2][2] == sym)) {
        return true;
      }
    }
    return false;
  }
}
