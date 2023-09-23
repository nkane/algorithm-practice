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
	QuickFindType              = UnionDataType("quick_find")
	QuickUnionFindType         = UnionDataType("quick_union_find")
	WeightedQuickUnionFindType = UnionDataType("weighted_quick_union_find")
	QuickFindCompressedType    = UnionDataType("quick_find_compressed")
)

type UnionFind interface {
	Connected(p int, q int) bool
	Find(p int) int
	Union(p int, q int)
	GetCount() int
	GetIDs() []int
}

// QuickFind
// constructor: O(N)
// union: O(N)
// find: O(1)
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

// QuickUnionFind
// constructor: O(N)
// union: tree height
// find: tree height
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

// WeightedQuickUnionFind
// constructor: O(N)
// union: O(lg N)
// find: O(lg N)
type WeightedQuickUnionFind struct {
	IDs   []int
	Size  []int
	Count int
}

func (wf *WeightedQuickUnionFind) Connected(p int, q int) bool {
	pID := wf.Find(p)
	qID := wf.Find(q)
	return pID == qID
}

func (wf *WeightedQuickUnionFind) Find(p int) int {
	for p != wf.IDs[p] {
		p = wf.IDs[p]
	}
	return p
}

func (wf *WeightedQuickUnionFind) Union(p int, q int) {
	pID := wf.Find(p)
	qID := wf.Find(q)
	if pID == qID {
		return
	}
	if wf.Size[pID] < wf.Size[qID] {
		wf.IDs[pID] = qID
		wf.Size[qID] += wf.Size[pID]
	} else {
		wf.IDs[qID] = pID
		wf.Size[pID] += wf.Size[qID]
	}
	wf.Count--
}

func (wf *WeightedQuickUnionFind) GetCount() int {
	return wf.Count
}

func (wf *WeightedQuickUnionFind) GetIDs() []int {
	return wf.IDs
}

func CreateWeightedQuickUnionFind(n int) UnionFind {
	wf := WeightedQuickUnionFind{
		Count: n, IDs: make([]int, n),
		Size: make([]int, n),
	}
	for i := 0; i < n; i++ {
		wf.IDs[i] = i
		wf.Size[i] = 1
	}
	return &wf
}

type QuickFindPathCompression struct {
	IDs   []int
	Count int
}

func (qf *QuickFindPathCompression) Connected(p int, q int) bool {
	pID := qf.Find(p)
	qID := qf.Find(q)
	return pID == qID
}

func (qf *QuickFindPathCompression) Find(p int) int {
	root := p
	for qf.IDs[root] != root {
		root = qf.IDs[root]
	}
	// path compression
	for p != root {
		next := qf.IDs[p]
		qf.IDs[p] = root
		p = next
	}
	return root
}

func (qf *QuickFindPathCompression) Union(p int, q int) {
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

func (qf *QuickFindPathCompression) GetCount() int {
	return qf.Count
}

func (qf *QuickFindPathCompression) GetIDs() []int {
	return qf.IDs
}

func CreateQuickFindPathCompression(n int) UnionFind {
	qf := QuickFindPathCompression{
		Count: n,
		IDs:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		qf.IDs[i] = i
	}
	return &qf
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
	case WeightedQuickUnionFindType:
		uf = CreateWeightedQuickUnionFind(n)
	case QuickFindCompressedType:
		uf = CreateQuickFindPathCompression(n)
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
