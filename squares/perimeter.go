package squares

import (
	"errors"
	"fmt"
)

//Clunky hashmap approach
func traverseBruteForce(coordinates []Point, startEdge Edge) ([]Edge, error) {
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
			//potential neighbouring squares
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
	//traverseBruteForce
	traversalPath := []Edge{startEdge}
	currentEdge := startEdge
	for {
		nextEdgeSquareCoordinate := squareMapByIndex[currentEdge.squareIndex]
		nextSquaresByPriority, associatedEdgeIndexes := getNextPossibleByPriority(nextEdgeSquareCoordinate, currentEdge.edge)
		for i := range nextSquaresByPriority {
			if index, found := squareMapByCoordinate[nextSquaresByPriority[i]]; found {
				edgeCandidate := Edge{index, associatedEdgeIndexes[i]}
				//check if edge candidate is a peripheral edge
				if peripheralEdgeSet[edgeCandidate] {
					if edgeCandidate == startEdge {
						//end traversal if you loop back to start
						return traversalPath, nil
					} else {
						traversalPath = append(traversalPath, edgeCandidate)
						currentEdge = edgeCandidate
						break
					}
				}
			}
		}
	}
}

// getNextPossibleByPriority
// constructs possible square and edge options for an edge, returned in descending priority
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

//More elegant, but not fully functionaly matrix approach
func traverseMatrix(coordinates []Point, startEdge Edge) ([]Edge, error) {
	//construct square search hashmap
	squareSetByCoordinate := make(map[Point]bool)
	for i := range coordinates {
		p := Point{coordinates[i].x, coordinates[i].y}
		squareSetByCoordinate[p] = true
	}
	//construct edge matrix
	var edgeMatrix [][4]bool
	for _, coordinate := range coordinates {
		row := [4]bool{}
		//potential neighbouring squares
		down := Point{coordinate.x, coordinate.y - 1}
		right := Point{coordinate.x + 1, coordinate.y}
		up := Point{coordinate.x, coordinate.y + 1}
		left := Point{coordinate.x - 1, coordinate.y}
		for i, option := range []Point{down, right, up, left} {
			//an edge is peripheral if it has nothing next to it
			if squareSetByCoordinate[option] {
				row[i] = true
			}
		}
		edgeMatrix = append(edgeMatrix, row)
	}

	x, y := startEdge.squareIndex, startEdge.edge
	length := len(edgeMatrix)

	//validate start edge
	if x >= length || y > 3 || !edgeMatrix[x][y] {
		return []Edge{}, errors.New(fmt.Sprintf("edge %+v invalid (non existant or non perimeter)", startEdge))
	}

	//traverseBruteForce matrix
	var traversalPath []Edge
	var reverse bool
	for x >= 0 {
		if x == length {
			reverse = true
			x = x - 2
			continue
		}
		if edgeMatrix[x][y] {
			traversalPath = append(traversalPath, Edge{squareIndex: x, edge: y})
			if y < 3 && edgeMatrix[x][y+1] {
				y++
			} else {
				if !reverse {
					x++
				} else {
					x--
				}
			}
		} else {
			if !reverse {
				x++
			} else {
				x--
			}
		}
	}
	return traversalPath, nil
}
