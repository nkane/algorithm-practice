package algorithms

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type UnionFind struct {
	IDs   []int
	Count int
}

func (uf *UnionFind) Connected(p int, q int) bool {
	// TODO(nick): implement
	return false
}

func (uf *UnionFind) Union(p int, q int) {
	// TODO(nick): implement
}

func CreateUnionFind(n int) *UnionFind {
	uf := UnionFind{
		Count: n,
		IDs:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.IDs[i] = i
	}
	return &uf
}

func CreateUnionFindFromFile(path string) (*UnionFind, error) {
	rf, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fs := bufio.NewScanner(rf)
	fs.Split(bufio.ScanLines)
	if !fs.Scan() {
		return nil, errors.New("failed to read first line")
	}
	nString := fs.Text()
	n, err := strconv.Atoi(nString)
	if err != nil {
		return nil, err
	}
	uf := CreateUnionFind(int(n))
	for fs.Scan() {
		line := fs.Text()
		tokens := strings.Fields(line)
		if len(tokens) >= 0 {
			firstToken := tokens[0]
			secondToken := tokens[1]
			p, err := strconv.Atoi(firstToken)
			if err != nil {
				return nil, err
			}
			q, err := strconv.Atoi(secondToken)
			if err != nil {
				return nil, err
			}
			if uf.Connected(p, q) {
				continue
			}
			uf.Union(p, q)
			log.Printf("%d %d\n", p, q)
		}
	}
	return uf, nil
}
