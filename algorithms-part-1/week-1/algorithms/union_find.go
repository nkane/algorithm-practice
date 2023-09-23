package algorithms

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type UnionDataType string

var (
	QuickFindType             = UnionDataType("quick_find")
	QuickUnionFindType        = UnionDataType("quick_union_find")
	WeightedQuickUionFindType = UnionDataType("weighted_quick_union_find")
)

type UnionFind interface {
	Connected(p int, q int) bool
	Find(p int) int
	Union(p int, q int)
	GetCount() int
	GetIDs() []int
}

type QuickFind struct {
	IDs   []int
	Count int
}

func (qf *QuickFind) Connected(p int, q int) bool {
	pID := qf.Find(p)
	qID := qf.Find(q)
	return pID == qID
}

func (qf *QuickFind) Find(p int) int {
	return qf.IDs[p]
}

func (qf *QuickFind) Union(p int, q int) {
	// put p and q into the same component
	pID := qf.Find(p)
	qID := qf.Find(q)
	if pID == qID {
		return
	}
	for i := 0; i < len(qf.IDs); i++ {
		if qf.IDs[i] == pID {
			qf.IDs[i] = qID
		}
	}
	qf.Count--
}

func (qf *QuickFind) GetCount() int {
	return qf.Count
}

func (qf *QuickFind) GetIDs() []int {
	return qf.IDs
}

func CreateQuickFind(n int) UnionFind {
	qf := QuickFind{
		Count: n,
		IDs:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		qf.IDs[i] = i
	}
	return &qf
}

type QuickUnionFind struct {
	IDs   []int
	Count int
}

func (quf *QuickUnionFind) Connected(p int, q int) bool {
	pID := quf.Find(p)
	qID := quf.Find(q)
	return pID == qID
}

func (quf *QuickUnionFind) Find(p int) int {
	for p != quf.IDs[p] {
		p = quf.IDs[p]
	}
	return p
}

func (quf *QuickUnionFind) Union(p int, q int) {
	// put p and q into the same component
	pID := quf.Find(p)
	qID := quf.Find(q)
	if pID == qID {
		return
	}
	quf.IDs[pID] = qID
	quf.Count--
}

func (quf *QuickUnionFind) GetCount() int {
	return quf.Count
}

func (quf *QuickUnionFind) GetIDs() []int {
	return quf.IDs
}

func CreateQuickUnionFind(n int) UnionFind {
	uf := QuickUnionFind{
		Count: n,
		IDs:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.IDs[i] = i
	}
	return &uf
}

func CreateUnionFindFromFile(path string, t UnionDataType) (UnionFind, error) {
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
	log.Printf("init components: %d", n)
	var uf UnionFind
	switch t {
	case QuickUnionFindType:
		uf = CreateQuickUnionFind(n)
	case QuickFindType:
		uf = CreateQuickFind(n)
	}
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
	log.Printf("components: %d", uf.GetCount())
	return uf, nil
}
