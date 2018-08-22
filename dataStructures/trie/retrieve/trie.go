package retrieve

import (
	"fmt"
)

// Trie - interface for caller
type Trie interface {
	Add(s string)
	IsMember(s string) bool
	Remove(s string)
	Print()
}

type node struct {
	key      rune
	leaf     bool
	children []*node
}

//NewTrie - creates a Trie and returns
func NewTrie() Trie {
	n := node{0, false, make([]*node, 26)}
	return n
}

func (n node) Print() {
	printTrie(&n, "")
}

func printTrie(n *node, s string) {
	if n == nil {
		return
	}

	ss := s + string(n.key)

	if n.leaf {
		fmt.Println(ss)
	}

	for _, c := range n.children {
		printTrie(c, ss)
	}
}

func (n node) Remove(s string) {
	r := []rune(s)
	remove(&n, r, 0)
}

func remove(n *node, r []rune, i int) bool {
	emptyChildren := func(n *node) bool {
		for _, c := range n.children {
			if c != nil {
				return false
			}
		}
		return true
	}

	if n == nil {
		return false
	}

	if i == len(r) {
		n.leaf = false
		return emptyChildren(n)
	} else {
		//if i < len(r) {
		c := n.children[r[i]-'a']
		if remove(c, r, i+1) {
			n.children[r[i]-'a'] = nil
			return emptyChildren(n) && !n.leaf
		} else {
			return false
		}
	}
}

func (n node) Add(s string) {
	r := []rune(s)
	p := &n
	c := n.children
	for _, k := range r {
		j := int(k - 'a')
		if c[j] == nil {
			c[j] = &node{k, false /*i == len(r)-1*/, make([]*node, 26)}
		}
		p = c[j]
		c = c[j].children
	}
	p.leaf = true
}

func (n node) IsMember(s string) bool {
	r := []rune(s)
	p := &n
	c := n.children
	for _, k := range r {
		i := int(k - 'a')
		if c[i] == nil {
			return false
		}
		p = c[i]
		c = c[i].children
	}
	return p.leaf == true
}
