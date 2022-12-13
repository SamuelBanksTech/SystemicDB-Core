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
	"math"
	"time"
)

var ExpiredCollectionCycles int64

// startExpiredGC this is only called by the NewSystemicDB function it starts the garbage collector that is responsible
// for removing keys that have reached their expiry unix time
func (t *systemicDB) startExpiredGC() {
	go func() {
		for {
			//I suppose that if this was left running for decades it would need resetting at some point
			if ExpiredCollectionCycles == math.MaxInt64 {
				ExpiredCollectionCycles = 0
			}

			ExpiredCollectionCycles++
			time.Sleep(time.Second * 10)
			t.removeExpired()
		}
	}()
}

// GetCollectionCycleCount returns a int64 of a count of how many times the expiredGC has run, this is used for testing
func (t *systemicDB) GetCollectionCycleCount() *int64 {
	return &ExpiredCollectionCycles
}

func (t *systemicDB) removeExpired() {
	traverseTree(t.root, t)
}

// traverseTree gently traverses the tree and passes each node to the cleanNodes function for expiry checking
func traverseTree(node *node, root *systemicDB) {
	if node == nil {
		return
	}

	// Visit the left subtree
	traverseTree(node.left, root)

	// Call the cleanNodes function for the current node
	cleanNodes(node, root)

	// Visit the right subtree
	traverseTree(node.right, root)
}

// cleanNodes simple check to see if the node has reach or exceeded the expiry unix timestamp, if so it gets deleted
func cleanNodes(node *node, root *systemicDB) {
	t := time.Now().Unix()

	if node.expiryUnix <= t {
		root.Remove(node.KeyStr)
	}
}
