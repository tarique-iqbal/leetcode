package solution

import (
	"testing"
)

func buildGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	nodes := make([]*Node, len(adjList))
	for i := range adjList {
		nodes[i] = &Node{Val: i + 1}
	}
	for i, neighbors := range adjList {
		for _, neiIdx := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neiIdx-1])
		}
	}
	return nodes[0]
}

func deepEqualGraph(n1, n2 *Node) bool {
	return deepEqualHelper(n1, n2, map[*Node]*Node{})
}

func deepEqualHelper(n1, n2 *Node, visited map[*Node]*Node) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil {
		return false
	}
	if v, ok := visited[n1]; ok {
		return v == n2
	}

	if n1.Val != n2.Val || len(n1.Neighbors) != len(n2.Neighbors) {
		return false
	}

	visited[n1] = n2

	for i := range n1.Neighbors {
		if !deepEqualHelper(n1.Neighbors[i], n2.Neighbors[i], visited) {
			return false
		}
	}
	return true
}

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
	}{
		{
			name:    "Empty graph",
			adjList: [][]int{},
		},
		{
			name:    "Single node, no neighbors",
			adjList: [][]int{{}},
		},
		{
			name:    "Two nodes connected",
			adjList: [][]int{{2}, {1}},
		},
		{
			name:    "Four nodes in a cycle",
			adjList: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			name:    "Tree structure",
			adjList: [][]int{{2, 3}, {}, {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adjList)
			cloned := cloneGraph(original)

			if original == nil && cloned != nil {
				t.Fatal("Expected nil for cloned graph, got non-nil")
			}
			if original != nil && cloned == nil {
				t.Fatal("Expected non-nil for cloned graph, got nil")
			}

			if !deepEqualGraph(original, cloned) {
				t.Errorf("Cloned graph does not match original for test case: %s", tt.name)
			}

			if original != nil && original == cloned {
				t.Errorf("Expected different pointers, got same pointer for test case: %s", tt.name)
			}
		})
	}
}
