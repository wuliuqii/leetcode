package tree

// G(n): 长度为 n 的序列能构成的不同二叉搜索树的个数
// F(i,n): 以 i 为根、序列为长度为 n 的不同二叉搜索树的个数 (1<=i<n)
// F(i,n) = G(i-1)*G(n-i)
// G(n) = ∑F(i,n), G(0)=G(1)=1
// G(n) = ∑G(i-1)*G(n-i)
func numTrees(n int) int {
	G := make([]int, n+1)
	G[0], G[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			G[i] += G[j-1] * G[i-j]
		}
	}
	return G[n]
}
