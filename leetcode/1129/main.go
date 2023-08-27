package main

// 给定一个整数 n，即有向图中的节点数，其中节点标记为 0 到 n - 1。图中的每条边为红色或者蓝色，并且可能存在自环或平行边。
// 给定两个数组 redEdges 和 blueEdges，其中：
// redEdges[i] = [ai, bi] 表示图中存在一条从节点 ai 到节点 bi 的红色有向边，
// blueEdges[j] = [uj, vj] 表示图中存在一条从节点 uj 到节点 vj 的蓝色有向边。
// 返回长度为 n 的数组 answer，其中 answer[X] 是从节点 0 到节点 X 的红色边和蓝色边交替出现的最短路径的长度。如果不存在这样的路径，那么 answer[x] = -1。
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	type pair struct{ x, color int }
	graph := make([][]pair, n)
	for _, edge := range redEdges {
		graph[edge[0]] = append(graph[edge[0]], pair{edge[1], 0})
	}
	for _, edge := range blueEdges {
		graph[edge[0]] = append(graph[edge[0]], pair{edge[1], 1})
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}

	visited := make([][2]bool, n)
	visited[0] = [2]bool{true, true}
	queue := []pair{{0, 0}, {0, 1}}
	for level := 0; len(queue) > 0; level++ {
		q := queue
		queue = nil
		for _, p := range q {
			x := p.x
			if dis[x] < 0 {
				dis[x] = level
			}

			for _, to := range graph[x] {
				if to.color != p.color && !visited[to.x][to.color] {
					visited[to.x][to.color] = true
					queue = append(queue, to)
				}
			}
		}
	}

	return dis
}
