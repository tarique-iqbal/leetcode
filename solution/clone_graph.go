package solution

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)

	queue := []*Node{node}
	visited[node] = &Node{Val: node.Val}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, nei := range cur.Neighbors {
			if _, ok := visited[nei]; !ok {
				visited[nei] = &Node{Val: nei.Val}
				queue = append(queue, nei)
			}

			visited[cur].Neighbors = append(visited[cur].Neighbors, visited[nei])
		}
	}

	return visited[node]
}
