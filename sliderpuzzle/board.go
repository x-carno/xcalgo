package sliderpuzzle

import (
	"strconv"
	"strings"
)

type Board struct {
	tiles     [][]uint64
	dimension int
}

func NewBoard(dim uint32) *Board {
	return &Board{
		tiles:     make([][]uint64, dim),
		dimension: int(dim),
	}
}

func (b *Board) ToString() string {
	ret := strings.Builder{}
	ret.WriteString(strconv.Itoa(b.dimension))
	ret.WriteString("\n")

	for _, row := range b.tiles {
		ret.WriteString(" ")
		for _, col := range row {
			ret.WriteString(strconv.FormatUint(col, 10))
			ret.WriteString(" ")
		}
		ret.WriteString("\n")
	}
	return ret.String()
}

func (b *Board) Hamming() uint64 {
	var ret uint64 = 0

	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if b.tiles[i][j]%uint64(b.dimension) != uint64(j+1) {
				ret++
			}
		}
	}

	return ret
}

func (b *Board) Manhattan() uint64 {
	var ret uint64 = 0

	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			ret += b.TileDistance(b.tiles[i][j], i, j)
		}
	}

	return ret
}

func (b *Board) TileDistance(num uint64, curI, curJ int) uint64 {
	desireI := num/uint64(b.dimension) - 1

	j := num % uint64(b.dimension)
	var desireJ uint64
	if j == 0 {
		desireJ = uint64(b.dimension) - 1
	} else {
		desireJ = j - 1
	}

	var vecI uint64
	if desireI > uint64(curI) {
		vecI = desireI - uint64(curI)
	} else {
		vecI = uint64(curI) - desireI
	}

	var vecJ uint64
	if desireJ > uint64(curJ) {
		vecJ = desireJ - uint64(curJ)
	} else {
		vecJ = uint64(curJ) - desireJ
	}

	return vecI + vecJ
}
