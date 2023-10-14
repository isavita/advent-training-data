package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DFS(node int, adj map[int][]int, visited map[int]bool) {
	visited[node] = true
	for _, neighbor := range adj[node] {
		if !visited[neighbor] {
			DFS(neighbor, adj, visited)
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer file.Close()

	adj := make(map[int][]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " <-> ")
		from, _ := strconv.Atoi(parts[0])
		toNodes := strings.Split(parts[1], ", ")

		for _, toNode := range toNodes {
			to, _ := strconv.Atoi(toNode)
			adj[from] = append(adj[from], to)
			adj[to] = append(adj[to], from)
		}
	}

	visited := make(map[int]bool)
	groups := 0

	for node := range adj {
		if !visited[node] {
			DFS(node, adj, visited)
			groups++
		}
	}

	fmt.Println(groups)
}
