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
	"fmt"
	"testing"
	"time"
)

var sdb *systemicDB
var testKey = "test-key"
var testValue = []byte("this is some test data")

func init() {
	sdb = NewSystemicDB()
}

func TestNewSystemicDB(t *testing.T) {

	if sdb.Meta.InitTime == 0 {
		t.Error("did not get a pointer to a new systemicdb instance")
	}
}

func TestSystemicDB_Insert(t *testing.T) {

	returnedNode := sdb.Insert(testKey, testValue, 10*time.Minute)

	if returnedNode == nil {
		t.Error("node was nil after inserting")
	}
}

func TestSystemicDB_Exists(t *testing.T) {

	got := sdb.Exists(testKey)
	want := true

	if got != want {
		t.Error("testing exists, wanted `true` got `false`")
	}

}

func TestSystemicDB_Get(t *testing.T) {
	got := sdb.Get(testKey)
	want := testValue
	wantTime := time.Now().Unix()

	if string(got.Value) != string(want) {
		t.Errorf("value returned from search was not the same as was set, wanted `%s` got `%s`", string(want), string(got.Value))
	}

	if got.Expiry <= wantTime {
		t.Error("expiry not set correctly")
	}
}

func TestSystemicDB_Remove(t *testing.T) {
	sdb.Remove(testKey)

	existsCheck := sdb.Exists(testKey)
	getCheck := sdb.Get(testKey)

	if existsCheck {
		t.Error("key should no longer exist")
	}

	if getCheck != nil {
		t.Error("key should no longer exist but is returning data")
	}
}

func TestSystemicDB_IsBalanced(t *testing.T) {
	sdb.Insert(testKey, testValue, 9*time.Second)
	sdb.Insert(testKey+"abc", testValue, 10*time.Minute)

	balanced := sdb.IsBalanced()

	if !balanced {
		t.Error("tree should be balanced after inserts and removes")
	}
}

func TestSystemicDB_Max(t *testing.T) {
	got := sdb.Max()
	var want uint64
	want = 15642096697511069609

	if got != want {
		t.Errorf("max should be %d but was %d", want, got)
	}
}

func TestSystemicDB_Min(t *testing.T) {
	got := sdb.Min()
	var want uint64
	want = 9875687615137630151

	if got != want {
		t.Errorf("min should be %d but was %d", want, got)
	}
}

func TestSystemicDB_GetCollectionCycleCount(t *testing.T) {

	time.Sleep(time.Millisecond * 200)

	if *sdb.GetCollectionCycleCount() == 0 {
		t.Error("expired data garbage collector does not appear to be running")
	}
}

func TestExpiredKeyGC(t *testing.T) {

	fmt.Println("Sleeping for 11 seconds to allow for GC to test removing expired keys")

	time.Sleep(time.Second * 11)

	if sdb.Exists(testKey) {
		t.Error("data still exists but should have been auto removed")
	}

}
