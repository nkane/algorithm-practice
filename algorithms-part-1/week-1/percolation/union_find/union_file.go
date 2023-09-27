package union_find

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type UnionFile struct {
	N     int
	Pairs [][]int
}

func CreateUnionFile(path string) (*UnionFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	fs := bufio.NewScanner(file)
	if !fs.Scan() {
		return nil, errors.New("failed to read first line of file")
	}
	nString := fs.Text()
	n, err := strconv.Atoi(nString)
	if err != nil {
		return nil, err
	}
	unionFile := UnionFile{
		N:     n,
		Pairs: make([][]int, 0),
	}
	i := 0
	for fs.Scan() {
		unionFile.Pairs = append(unionFile.Pairs, make([]int, 2))
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
			unionFile.Pairs[i][0] = p
			unionFile.Pairs[i][1] = q
			i++
		}
	}
	return &unionFile, nil
}
