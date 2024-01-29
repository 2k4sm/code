package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
}

func lot(root *TreeNode) [][]int {
	var levelList [][]int
	queue := []*TreeNode{}
	if root == nil {
		return levelList
	}
	queue = append(queue, root)
	for(len(queue) != 0){
		level := make([]int,0)
		for int i = 0; i < len(queue);i++{
			curval := queue[0]
			queue = queue[1:]
			level = append(level, curval.Val);
			
			if curval.Left != nil{
				queue = append(queue, curval.Left)	
			}
			
			if curval.Right != nil {
				queue = append(queue, curval.Right)	
			}
		}
	}
}
