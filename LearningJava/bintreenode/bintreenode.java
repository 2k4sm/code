public class bintreenode{
    public static void main(String[] args) {
        TreeNode rootTn = new TreeNode(10);
        rootTn.left = new TreeNode(5);
        rootTn.right = new TreeNode(3);
        rootTn.left.right = new TreeNode(2);
        rootTn.left.left = new TreeNode(4);

        postOrder(rootTn);

    }

    public static void preOrder(TreeNode root){
        if (root == null){
            return;
        }
        System.out.println((root.val));
        preOrder(root.left);
        preOrder(root.right);
    }

    public static void postOrder(TreeNode root){
        if (root == null){
            return;
        }
        System.out.println((root.val));
        postOrder(root.right);
        postOrder(root.left);
         
    }

}

class TreeNode{
    TreeNode left;
    TreeNode right;
    int val;

    public TreeNode(int x){
        this.val = x;
        this.left = null;
        this.right = null;
    }
}


