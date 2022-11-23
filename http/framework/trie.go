package framework

import (
	"fmt"
	"strings"
)

func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

func NewTree() *Tree {
	return &Tree{
		root: &node{
			children: []*node{},
		},
	}
}

type Tree struct {
	root *node
}

type node struct {
	isLast   bool
	segment  string
	handler  ControllerHandler
	children []*node
}

func (n *node) filterChildNodes(segment string) []*node {
	if len(n.children) == 0 {
		return nil
	}

	if isWildSegment(segment) {
		return n.children
	}
	nodes := make([]*node, 0, len(n.children))

	for _, c := range n.children {
		if c.segment == segment || isWildSegment(c.segment) {
			nodes = append(nodes, c)
		}
	}
	return nodes
}

func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)
	segment := segments[0]

	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	cs := n.filterChildNodes(segment)
	if len(cs) == 0 {
		return nil
	}

	if len(segments) == 1 {
		for _, c := range cs {
			if c.isLast {
				return c
			}
		}
		return nil
	}

	for _, c := range cs {
		if match := c.matchNode(segments[1]); match != nil {
			return match
		}
	}
	return nil
}

func (t *Tree) AddRouter(uri string, h ControllerHandler) error {
	n := t.root
	if n.matchNode(uri) != nil {
		return fmt.Errorf("route exists: %s", uri)
	}

	segs := strings.Split(uri, "/")

	for i, s := range segs {
		if !isWildSegment(s) {
			s = strings.ToUpper(s)
		}
		isLast := i == len(segs)-1

		var objNode *node

		childNodes := n.filterChildNodes(s)

		if len(childNodes) > 0 {
			for _, c := range childNodes {
				if c.segment == s {
					objNode = c
					break
				}
			}
		}

		if objNode == nil {
			c := &node{}
			c.segment = s
			if isLast {
				c.isLast = true
				c.handler = h
			}
			n.children = append(n.children, c)
			objNode = c
		}
		n = objNode
	}

	return nil
}

func (t *Tree) FindHandler(uri string) ControllerHandler {
	n := t.root.matchNode(uri)
	if n == nil {
		return nil
	}
	return n.handler
}
