package main

import "fmt"

func maximalNetworkRank(n int, roads [][]int) (ans int) {
	//建立个二维数组 n*n 的
	//有点像图的邻接矩阵这样子的
	connect := make([][]int, n)
	for i := 0; i < n; i++ {
		connect[i] = make([]int, n)
	}
	//节点的度
	degree := make([]int, n)
	for _, v := range roads {
		//两个城市有连接，则邻接矩阵置 1
		connect[v[0]][v[1]] = 1
		connect[v[1]][v[0]] = 1
		//度也要 +1
		degree[v[0]]++
		degree[v[1]]++
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dis := degree[i] + degree[j] - connect[i][j]
			ans = max(dis, ans)
		}
	}
	return
	//图入门吧算是，把邻接矩阵这些都用上了
}

func main() {
	fmt.Println(maximalNetworkRank(8, [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}, {5, 6}, {5, 7}})) //5
	fmt.Println(maximalNetworkRank(5, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4}})) //5
	fmt.Println(maximalNetworkRank(5, [][]int{{2, 3}, {0, 3}, {0, 4}, {4, 1}}))                 //4
	fmt.Println(maximalNetworkRank(6, [][]int{{2, 4}}))                                         //1
}
