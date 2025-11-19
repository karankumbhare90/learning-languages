package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// Sort by ID so we can validate sequence and build in order.
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	n := len(records)
	nodes := make([]*Node, n)

	// Validate all IDs and parent relationships
	for i, r := range records {
		if r.ID != i {
			return nil, errors.New("record IDs must be continuous and start at 0")
		}
		if r.ID == 0 {
			if r.Parent != 0 {
				return nil, errors.New("root record must have parent ID = 0")
			}
		} else {
			if r.Parent >= r.ID {
				return nil, errors.New("parent ID must be lower than child ID")
			}
		}

		nodes[i] = &Node{ID: r.ID}
	}

	// Link children
	for _, r := range records {
		if r.ID == 0 {
			continue // root has no parent link
		}
		parent := nodes[r.Parent]
		parent.Children = append(parent.Children, nodes[r.ID])
	}

	return nodes[0], nil
}
