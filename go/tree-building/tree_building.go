package tree

import (
	"errors"
	"fmt"
	"sort"
)

// Record represents an input
type Record struct {
	ID, Parent int
}

// Node represents a node of a tree
type Node struct {
	ID       int
	Children []*Node
}

// Mismatch represents an error
type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

// Build construct a tree based on records input
func Build(records []Record) (*Node, error) {

	// Sort input records by its ID
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if len(records) <= 0 {
		return nil, nil
	}

	var root Node

	for i, rec := range records {

		if i == 0 && rec.ID == 0 {

			if rec.ID != 0 || rec.Parent != 0 {
				return nil, errors.New("Invalid Parent")
			}

			root.ID = records[0].ID
		} else {

			if rec.ID >= len(records) {
				return nil, fmt.Errorf("Out of Bounds")
			}

			if rec.ID == rec.Parent {
				return nil, fmt.Errorf("ID is equal to Parent")
			}

			if rec.ID < rec.Parent {
				return nil, fmt.Errorf("Parent greater than ID")
			}

			err := chk(&root, rec.ID)
			if err != nil {
				return nil, err
			}

			if p := findParent(rec.Parent, &root); p != nil {
				n := Node{ID: rec.ID}
				p.Children = append(p.Children, &n)
			}
		}
	}

	return &root, nil
}

func findParent(id int, nodes *Node) *Node {
	if id == nodes.ID {
		return nodes
	}

	var ret *Node
	for _, c := range nodes.Children {
		ret = findParent(id, c)
		if ret != nil {
			return ret
		}
	}

	return ret
}

func chk(n *Node, m int) (err error) {
	if n.ID > m {
		return fmt.Errorf("z")
	} else if n.ID == m {
		return fmt.Errorf("y")
	} else {
		for i := 0; i < len(n.Children); i++ {
			err = chk(n.Children[i], m)
			if err != nil {
				return
			}
		}
		return
	}
}
