package main

func shortestPathLength(graph [][]int) int {
	nodes := len(graph)
	if nodes <= 1 {
		return 0
	}
	type state struct {
		node    int
		visited int
	}
	seen := make([][]bool, 1<<nodes)
	for i := range seen {
		seen[i] = make([]bool, nodes)
	}
	queue := make([]*state, nodes)
	for i := range queue {
		v := 1 << i
		queue[i] = &state{node: i, visited: v}
		seen[v][i] = true
	}
	endingMask := 1<<nodes - 1
	steps := 0
	for {
		nextQueue := []*state{}
		for _, s := range queue {
			for _, neighbor := range graph[s.node] {
				v := s.visited | 1<<neighbor
				if v == endingMask {
					return 1 + steps
				}
				if !seen[v][neighbor] {
					seen[v][neighbor] = true
					nextQueue = append(nextQueue, &state{node: neighbor, visited: v})
				}
			}
		}
		steps++
		queue = nextQueue
	}
}
