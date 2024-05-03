public class queens {
    public static void main(String[] args) {
        int n = 4;
        int[][] board = new int[n][n];
        new queens().NQueens(board, n, 0);
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                System.out.print(board[i][j] + " ");
            }
            System.out.println();
        }
    }

    public void NQueens(int[][] board, int n, int row) {
        if (row == n) {
            return;
        }
        for (int j = 0; j < n; j++) {
            if (checkConfig(board, n, row, j)) {
                board[row][j] = 1;
                NQueens(board, n, row + 1);
                board[row][j] = 0;
            }
        }
    }

    public boolean checkConfig(int[][] board, int n, int row, int col) {
        for (int i = 0; i < n; i++) {
            if (board[row][i] == 1) {
                return false;
            }

            if (board[i][col] == 1) {
                return false;
            }
            if ((row - i >= 0 && row + i < n && col - i >= 0 && col + i < n) && (board[row - i][col - i] == 1 || board[row - i][col + i] == 1
            || board[row + i][col - i] == 1 || board[row + 1][col + 1] == 1)) {
                return false;
            }
        }

        return true;
    }

}