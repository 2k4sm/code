public class bst{
  public static void main(String Args[]){

  }

  public TreeNode insert(TreeNode root, int k){
    //Base case

    if (root.left == null){
      root.left = new TreeNode(k);
    }
    if (root.right == null){
      root.right = new TreeNode(k);
    }

    //Recursive case;
    
    if (root.val > k){
      return insert(root.left,k);
    }else{
      return insert(root.right,k);
    }

  }
}

class TreeNode{
  int val;
  TreeNode left;
  TreeNode right;

  public TreeNode(int x){
    this.val = x;
    this.left = null;
    this.right = null;
  }
}
