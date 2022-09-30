package squares

import (
	"errors"
	"fmt"
)

func traverse(coordinates []Point, startEdge Edge) ([]Edge, error) {
	//construct square search hashmap
	squareMapByCoordinate := make(map[Point]int)
	squareMapByIndex := make(map[int]Point)
	for i := range coordinates {
		squareMapByCoordinate[Point{coordinates[i].x, coordinates[i].y}] = i
		squareMapByIndex[i] = Point{coordinates[i].x, coordinates[i].y}
	}
	//construct a set of peripheral edges
	peripheralEdgeSet := make(map[Edge]bool)
	var e Edge
	for !peripheralEdgeSet[e] {
		for coordinate, squareIndex := range squareMapByCoordinate {
			down := Point{coordinate.x, coordinate.y - 1}
			right := Point{coordinate.x + 1, coordinate.y}
			up := Point{coordinate.x, coordinate.y + 1}
			left := Point{coordinate.x - 1, coordinate.y}
			for i, option := range []Point{down, right, up, left} {
				//an edge is peripheral if it has nothing next to it
				if _, found := squareMapByCoordinate[option]; !found {
					e := Edge{squareIndex, i}
					peripheralEdgeSet[e] = true
				}
			}
		}
	}
	//validate startEdge
	if !peripheralEdgeSet[startEdge] {
		return []Edge{}, errors.New(fmt.Sprintf("edge %+v invalid (non existant or non perimeter)", startEdge))
	}
	//traverse
	traversal := []Edge{startEdge}
	nextEdge := startEdge
	for {
		nextEdgeSquareCoordinate := squareMapByIndex[nextEdge.squareIndex]
		nextSquaresByPriority, associatedEdgeIndexes := getNextPossibleByPriority(nextEdgeSquareCoordinate, nextEdge.edge)
		for i := range nextSquaresByPriority {
			if index, found := squareMapByCoordinate[nextSquaresByPriority[i]]; found {
				edgeCandidate := Edge{index, associatedEdgeIndexes[i]}
				if peripheralEdgeSet[edgeCandidate] {
					if edgeCandidate == startEdge {
						return traversal, nil
					} else {
						traversal = append(traversal, edgeCandidate)
						nextEdge = edgeCandidate
						break
					}
				}
			}
		}
	}
}

func getNextPossibleByPriority(self Point, edgeIndex int) ([]Point, []int) {
	down := Point{self.x, self.y - 1}
	downRight := Point{self.x + 1, self.y - 1}
	right := Point{self.x + 1, self.y}
	upRight := Point{self.x + 1, self.y + 1}
	up := Point{self.x, self.y + 1}
	upLeft := Point{self.x - 1, self.y + 1}
	left := Point{self.x - 1, self.y}
	downLeft := Point{self.x - 1, self.y - 1}
	switch edgeIndex {
	case 0:
		return []Point{downRight, right, self}, []int{3, 0, 1}
	case 1:
		return []Point{upRight, up, self}, []int{0, 1, 2}
	case 2:
		return []Point{upLeft, left, self}, []int{1, 2, 3}
	case 3:
		return []Point{downLeft, down, self}, []int{2, 3, 0}
	}
	return []Point{}, []int{}
}
