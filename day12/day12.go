package day12

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Name        string
	IsBig       bool
	Connections map[string]*Node
}

func IsBigName(name string) bool {
	firstChar := name[0]
	if firstChar >= "a"[0] && firstChar <= "z"[0] {
		return false
	}
	return true
}

func NewNode(name string) *Node {
	return &Node{
		Name:        name,
		IsBig:       IsBigName(name),
		Connections: make(map[string]*Node),
	}
}

func (n *Node) Link(other *Node) {
	n.Connections[other.Name] = other
	other.Connections[n.Name] = n
}

func GetOrCreate(nodes map[string]*Node, name string) *Node {
	node := nodes[name]
	if node == nil {
		node = NewNode(name)
		nodes[name] = node
	}
	return node
}

func PrintPath(path []*Node) {
	result := make([]string, 0)
	for _, node := range path {
		result = append(result, node.Name)
	}
	fmt.Println(strings.Join(result, "-"))
}

func PrintAllPaths(nodes map[string]*Node, start, end *Node, visited []*Node, previousTwice *string) int {
	path := make([]*Node, 0)
	path = append(path, visited...)
	visited = append(visited, start)
	path = append(path, start)
	count := 0
MainLoop:
	for _, node := range start.Connections {
		twice := previousTwice
		if node == end {
			path = append(path, end)
			//PrintPath(path)
			count++
			continue
		}
		if !node.IsBig {
			for _, visitedNode := range visited {
				if visitedNode == node {
					if twice == nil && node.Name != "start" && node.Name != "end" && start.Name != node.Name {
						// do nothing
						twice = &node.Name
						break
					}
					continue MainLoop
				}
			}
		}
		// Skip subsequents to same nodes
		if node.Name == start.Name {
			continue
		}
		count += PrintAllPaths(nodes, node, end, visited, twice)
	}
	return count
}

func Solution() {
	file, _ := os.Open("day12/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	nodes := make(map[string]*Node)

	for scanner.Scan() {
		line := scanner.Text()
		names := strings.Split(line, "-")
		node1 := GetOrCreate(nodes, names[0])
		node2 := GetOrCreate(nodes, names[1])
		node1.Link(node2)
	}

	result := PrintAllPaths(nodes, nodes["start"], nodes["end"], make([]*Node, 0), nil)

	fmt.Println(result)

}
