//Copyright 2022 SamuelBanksTech
//
//Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation
//files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy,
//modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the
//Software is furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
//WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
//COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
//OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package systemicdb

import (
	"hash/fnv"
	"math"
	"time"
)

// strToInt function takes a string and returns an uint64 hash generated by the FNV1a algo
func strToInt(str string) (uint64, error) {

	h := fnv.New64a()
	_, err := h.Write([]byte(str))
	if err != nil {
		return 0, err
	}

	return h.Sum64(), nil
}

// Insert inserts an element into the AVL Tree and returns the newly created node (includes a rebalance of the tree)
func (t *systemicDB) Insert(keyStr string, value []byte, expiry time.Duration) *node {

	key, err := strToInt(keyStr)
	if err != nil {
		return nil
	}

	t.root = insert(t.root, key, keyStr, value, expiry)

	return t.root
}

func insert(n *node, key uint64, keyStr string, value []byte, expiry time.Duration) *node {

	if n == nil || n.key == key {
		t := time.Now()
		t = t.Add(expiry)

		return &node{KeyStr: keyStr, Value: value, Expiry: expiry, expiryUnix: t.Unix(), key: key, height: 0}
	}

	if key < n.key {
		n.left = insert(n.left, key, keyStr, value, expiry)
	} else {
		n.right = insert(n.right, key, keyStr, value, expiry)
	}

	n.height = 1 + int(math.Max(float64(height(n.left)), float64(height(n.right))))
	n = balance(n)
	return n
}

// Remove removes an element from the AVL Tree (includes a rebalance of the tree)
func (t *systemicDB) Remove(keyStr string) {
	key, err := strToInt(keyStr)
	if err != nil {
		return
	}

	t.root = remove(t.root, key)
}

func remove(n *node, key uint64) *node {
	if n == nil {
		return nil
	}

	if key < n.key {
		n.left = remove(n.left, key)
	} else if key > n.key {
		n.right = remove(n.right, key)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		min := minKey(n.right)
		n.key = min
		n.right = remove(n.right, min)
	}

	n.height = 1 + int(math.Max(float64(height(n.left)), float64(height(n.right))))
	n = balance(n)
	return n
}

// Exists searches for an element in the AVL Tree and return a bool (true is exists)
func (t *systemicDB) Exists(keyStr string) bool {
	key, err := strToInt(keyStr)
	if err != nil {
		return false
	}

	return exists(t.root, key)
}

func exists(n *node, key uint64) bool {
	if n == nil {
		return false
	}

	if key < n.key {
		return exists(n.left, key)
	} else if key > n.key {
		return exists(n.right, key)
	} else {
		return true
	}
}

// Get searches for an element in the AVL Tree and returns a NodeData struct which is important parts of the node that a user could need
func (t *systemicDB) Get(keyStr string) *NodeData {
	key, err := strToInt(keyStr)
	if err != nil {
		return nil
	}

	n := get(t.root, key)
	if n == nil {
		return nil
	}

	returnNode := NodeData{
		Key:    n.KeyStr,
		Value:  n.Value,
		Expiry: n.expiryUnix,
	}

	return &returnNode
}

func get(n *node, key uint64) *node {
	if n == nil {
		return nil
	}

	if key < n.key {
		return get(n.left, key)
	} else if key > n.key {
		return get(n.right, key)
	} else {
		return n
	}
}

// IsBalanced returns true if the AVL Tree is balanced.
func (t *systemicDB) IsBalanced() bool {
	return isBalanced(t.root)
}

func isBalanced(n *node) bool {
	if n == nil {
		return true
	}

	diff := height(n.left) - height(n.right)
	if diff > 1 || diff < -1 {
		return false
	}

	return isBalanced(n.left) && isBalanced(n.right)
}

// Min returns the minimum element in the AVL Tree.
func (t *systemicDB) Min() uint64 {
	return minKey(t.root)
}

func minKey(n *node) uint64 {
	if n == nil {
		return uint64(math.MaxInt64)
	}

	if n.left == nil {
		return n.key
	}

	return minKey(n.left)
}

// Max returns the maximum element in the AVL Tree.
func (t *systemicDB) Max() uint64 {
	return maxKey(t.root)
}

func maxKey(n *node) uint64 {
	if n == nil {
		return 0
	}

	if n.right == nil {
		return n.key
	}

	return maxKey(n.right)
}

func height(n *node) int {
	if n == nil {
		return -1
	}
	return n.height
}

// balance performs the all important balancing of the AVL Tree, this is called on every insert and remove to keep the tree balanced and running smoothly
func balance(n *node) *node {
	diff := height(n.left) - height(n.right)
	if diff > 1 {
		if height(n.left.left) >= height(n.left.right) {
			return rightRotate(n)
		} else {
			n.left = leftRotate(n.left)
			return rightRotate(n)
		}
	} else if diff < -1 {
		if height(n.right.right) >= height(n.right.left) {
			return leftRotate(n)
		} else {
			n.right = rightRotate(n.right)
			return leftRotate(n)
		}
	}

	return n
}

func rightRotate(n *node) *node {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.height = 1 + int(math.Max(float64(height(n.left)), float64(height(n.right))))
	newRoot.height = 1 + int(math.Max(float64(height(newRoot.left)), float64(height(newRoot.right))))
	return newRoot
}

func leftRotate(n *node) *node {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.height = 1 + int(math.Max(float64(height(n.left)), float64(height(n.right))))
	newRoot.height = 1 + int(math.Max(float64(height(newRoot.left)), float64(height(newRoot.right))))
	return newRoot
}
