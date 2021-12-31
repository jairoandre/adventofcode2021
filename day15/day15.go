package day15

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	ID          int
	Risk        int
	Connections map[int]*Node
}

func NewNode(id, risk int) *Node {
	return &Node{
		ID:          id,
		Risk:        risk,
		Connections: make(map[int]*Node),
	}
}

func (n *Node) Link(other *Node) {
	n.Connections[other.ID] = other
	other.Connections[n.ID] = n
}

func Solution() {
	file, _ := os.Open("day15/input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()
	allNodes := make(map[int]*Node)
	id := 0
	cols := 0
	rows := 0
	input := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
		if cols == 0 {
			cols = len(line)
		}
		for _, v := range line {
			risk, _ := strconv.Atoi(string(v))
			allNodes[id] = NewNode(id, risk)
			id++
		}
		rows++
	}

	// Make connections
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			currId := x + y*cols
			currNode := allNodes[currId]
			if x > 0 {
				currNode.Link(allNodes[currId-1])
			}
			if x < cols-1 {
				currNode.Link(allNodes[currId+1])
			}
			if y > 0 {
				currNode.Link(allNodes[currId-cols])
			}
			if y < rows-1 {
				currNode.Link(allNodes[currId+cols])
			}
		}
	}
	Part1(allNodes)

}

func Dijkstra(allNodes map[int]*Node) int {
	start := allNodes[0]
	dist := make([]int, 0)
	maxInt := int(^uint(0) >> 1)
	for _, _ = range allNodes {
		dist = append(dist, maxInt)
	}
	// Do not consider the initial risk
	dist[start.ID] = 0
	visited := make(map[int]bool)
	queue := make(chan *Node, 300)
	queue <- start
	for len(queue) > 0 {
		currNode := <-queue
		if visited[currNode.ID] {
			continue
		}
		visited[currNode.ID] = true
		// expand all connections
		for otherId, otherNode := range currNode.Connections {
			thisCost := dist[currNode.ID] + otherNode.Risk
			if dist[otherId] > thisCost {
				dist[otherId] = thisCost
				queue <- otherNode
			}
		}
	}
	return dist[len(allNodes)-1]
}

func Part1(allNodes map[int]*Node) {
	fmt.Println("Part 1 answer:", Dijkstra(allNodes))
}
