package union_find

type UnionFind interface {
	Union(p int, q int)
	Find(p int) int
	Connected(p int, q int) bool
	Count() int
	IDs() []int
}
