import java.util.HashMap;

public class bintreenode{
    public static void main(String[] args) {
        int[] preorder = {1,2,3};
        int[] inorder = {2,1,3};
        TreeNode newTree = constructTree(preorder,inorder);
        preOrder(newTree);
        postOrder(newTree);


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
        postOrder(root.left);
        postOrder(root.right);
        System.out.println((root.val));
    }

    public static TreeNode constructTree(int[] preorder, int[] inorder){
        if (preorder.length == 0 || inorder.length == 0){
            return null;
        }
        HashMap<Integer,Integer> inorderMap = new HashMap<>();

        for (int i = 0; i < inorder.length; i++){
            inorderMap.put(inorder[i], i);
        }

        return build(preorder,0,0,inorder.length - 1,inorderMap);
    }

    public static TreeNode build(int[] preorder, int preOrderIndex,int inorderLow, int inorderHigh, HashMap<Integer,Integer> inorderMap){
        if (preOrderIndex > preorder.length || inorderLow > inorderHigh){
            return null;
        }

        int currval = preorder[preOrderIndex];
        TreeNode newTree = new TreeNode(currval);
        int mid = inorderMap.get(currval);

        newTree.left = build(preorder, preOrderIndex + 1, inorderLow, mid - 1, inorderMap);
        newTree.right = build(preorder, preOrderIndex + (mid - inorderLow) + 1, mid + 1, inorderHigh, inorderMap);

        return newTree;
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


