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

import "time"

type systemicDB struct {
	root *node
	Meta sdbMeta
}

type sdbMeta struct {
	NodeCount int64
	InitTime  int64
}

type NodeData struct {
	Key    string
	Value  []byte
	Expiry int64
}

type node struct {
	KeyStr     string
	Value      []byte
	Expiry     time.Duration
	expiryUnix int64
	key        uint64
	left       *node
	right      *node
	height     int
}

// NewSystemicDB simply creates a new instance of the SystemicDB AVL Tree, starts the expired garbage collector and returns a pointer
func NewSystemicDB() *systemicDB {
	sdb := systemicDB{
		Meta: sdbMeta{
			NodeCount: 0,
			InitTime:  time.Now().Unix(),
		},
	}

	sdb.startExpiredGC()

	return &sdb
}
