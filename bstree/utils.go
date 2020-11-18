package bstree

import "github.com/rianby64/data-structures-self-study/list"

func insert(t *bstree, v interface{}, c func(a, b interface{}) bool) BStree {
	t.length++
	if t.payload == nil {
		t.SetValue(v)

		return t
	}

	node := &bstree{
		root:       t.root,
		comparator: c,
		parent:     t,
	}

	if c(t.payload.Value(), v) {
		if t.left == nil {
			t.left = node
		}

		return insert(t.left, v, c)
	}
	if t.right == nil {
		t.right = node
	}

	return insert(t.right, v, c)
}

func delete(t, parent *bstree, v interface{}, c func(a, b interface{}) bool) BStree {

	// case 1: t.left == nil && t.right == nil -> leaf
	if t.left == nil && t.right == nil {
		if parent.left == t {
			parent.left = nil
		}
		if parent.right == t {
			parent.right = nil
		}

		return parent
	}

	return nil
}

func castTobtree(b BStree) (*bstree, bool) {
	casted, ok := b.(*bstree)
	return casted, ok
}

func insertNode(t BStree, node BStree, c func(a, b interface{}) bool) BStree {
	ct, ok := castTobtree(t)
	if !ok {
		return t
	}

	cnode, ok := castTobtree(node)
	if !ok {
		return t
	}

	v := node.Value()
	if v == nil {
		return ct
	}

	if c(ct.payload.Value(), v) {
		if ct.left == nil {
			ct.left = cnode

			return node
		}

		return insertNode(ct.left, node, c)
	}

	if ct.right == nil {
		ct.right = cnode

		return node
	}

	return insertNode(ct.right, node, c)
}

func find(a interface{}, t *bstree, matcher, comparator comparator) BStree {
	if matcher(a, t.Value()) {
		return t
	}

	if t.left != nil {
		if t.comparator(a, t.left.Value()) {
			found := find(a, t.left, matcher, comparator)
			if found != nil {
				return found
			}
		}
	}

	if t.right != nil {
		return find(a, t.right, matcher, comparator)
	}

	return nil
}

func findmax(t BStree) BStree {
	right := t.Right()
	if right != nil {
		return findmax(right)
	}
	left := t.Left()
	if left != nil {
		return findmax(left)
	}
	return t
}

func inorder(root *bstree, l list.List) {
	// la condicion || (root.root == root && root.Value() == nil) es por falta de sentinela
	if root == nil || (root.root == root && root.Value() == nil) {
		return
	}

	if root.left != nil {
		inorder(root.left, l)
	}

	l.Last().Insert(root.Value())

	if root.right != nil {
		inorder(root.right, l)
	}
}
