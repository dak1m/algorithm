package binary_tree

// TreeNode 基本的二叉树节点
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// 后序遍历二叉树
func traverse(root *TreeNode) {
	if root != nil {
		traverse(root.left)
		traverse(root.right)
	}
}

// NTreeNode N叉鼠遍历
// 基本的 N 叉树节点
type NTreeNode struct {
	val      int
	children []*TreeNode
}

func nTraverse(root *NTreeNode) {
	for _, child := range root.children {
		traverse(child)
	}
}

// 二叉树层序遍历
func levelTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	q := make([]*TreeNode, 0)
	q = append(q, root)
	depth := 1

	// 从上到下遍历二叉树的每一层
	for len(q) > 0 {
		sz := len(q)

		// 从左到右遍历每一层的每个节点
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]

			// 将下一层节点放入队列
			if cur.left != nil {
				q = append(q, cur.left)
			}
			if cur.right != nil {
				q = append(q, cur.right)
			}
		}
		depth++
	}
}
