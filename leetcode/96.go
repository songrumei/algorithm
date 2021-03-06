package leetcode

/*
	不同的二叉搜索树
*/

// 以每个数字作为根节点，dp[i] 表示 1~i的的不同二叉搜索树的数量
// 创建以 3 为根、长度为 77 的不同二叉搜索树，整个序列是 [1, 2, 3, 4, 5, 6, 7]，
// 我们需要从左子序列 [1, 2] 即dp[2] 构建左子树，从右子序列 [4, 5, 6, 7] 即 dp[4] 构建右子树，然后将它们组合（即笛卡尔积）
// 注意到 G(n)G(n) 和序列的内容无关，只和序列的长度有关
// 比如 [1,2] 与 [4,5] 构建为树结果是一样的
func numTrees(n int) int {
	dp := make([]int, n+1)
	// 对于边界情况，当序列长度为 1（只有根）或为 0（空树）时，只有一种情况
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		// 从 1~i 每个节点都作为根节点，j表示根节点
		for j := 1; j <= i; j++ {
			// 做笛卡尔乘积
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
