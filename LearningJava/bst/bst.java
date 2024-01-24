public class bst {
  public static void main(String Args[]) {
    TreeNode node = new TreeNode(20);
    insert(node, 14);
    insert(node, 13);
    insert(node, 11);
    insert(node, 10);
    insert(node, 15);
    insert(node, 16);
    insert(node, 17);
    insert(node, 12);
    insert(node, 1);
    insert(node, 3);

    preOrder(node);
    System.out.println(isBst(node).isBst);
  }

  public static void preOrder(TreeNode root) {
    if (root == null) {
      return;
    }
    System.out.println((root.val));
    preOrder(root.left);
    preOrder(root.right);
  }

  public static TreeNode insert(TreeNode root, int k) {
    // Base case

    if (root == null) {
      return new TreeNode(k);
    }

    // Recursive case;

    if (root.val > k) {
      root.left = insert(root.left, k);
    } else {
      root.right = insert(root.right, k);
    }

    return root;
  }

  public static Data isBst(TreeNode root) {
    if (root == null) {
      return new Data(true, Integer.MAX_VALUE, Integer.MIN_VALUE);
    }
    Data left = isBst(root.left);
    Data right = isBst(root.right);

    if (left.isBst && right.isBst && root.val > left.max &&
        root.val < right.min) {
      return new Data(true, checkmin(root.val, left.min, right.min),
                      checkmax(root.val, left.max, right.max));
    }
    return new Data(false, checkmin(root.val, left.min, right.min),
                    checkmax(root.val, left.max, right.max));
  }

  public static int checkmax(int val, int lmax, int rmax) {
    if (val > lmax) {
      if (val > rmax) {
        return val;
      } else {
        return rmax;
      }
    } else {
      if (lmax > rmax) {
        return lmax;
      } else {
        return rmax;
      }
    }
  }

  public static int checkmin(int val, int lmin, int rmin) {
    if (val < lmin) {
      if (val < rmin) {
        return val;
      } else {
        return rmin;
      }
    } else {
      if (lmin < rmin) {
        return lmin;
      } else {
        return rmin;
      }
    }
  }

  public static TreeNode delNode(TreeNode root, int val) {
    if (root == null) {
      return root;
    }

    if (root.val > val) {
      root.left = delNode(root.left, val);
    } else if (root.val < val) {
      root.right = delNode(root.right, val);
    } else {
      if (root.left == null && root.right == null) {
        return null;
      }

      if (root.left == null) {
        return root.right;
      }

      if (root.right == null) {
        return root.left;
      } else {
        int max = getmax(root.left);
        root.val = max;
        root.left = delNode(root.left, max);
      }
    }
    return root;
  }

  public static int getmax(TreeNode root) {
    TreeNode curNode = root;
    while (curNode.right != null) {
      curNode = curNode.right;
    }
    return curNode.val;
  }

  public static int getmin(TreeNode root) {
    TreeNode curNode = root;
    while (curNode.left != null) {
      curNode = curNode.left;
    }
    return curNode.val;
  }
}

class TreeNode {
  int val;
  TreeNode left;
  TreeNode right;

  public TreeNode(int x) {
    this.val = x;
    this.left = null;
    this.right = null;
  }
}

class Data {
  boolean isBst;
  int min;
  int max;

  public Data(boolean isBst, int min, int max) {
    this.isBst = isBst;
    this.min = min;
    this.max = max;
  }
}
