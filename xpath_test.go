package xpath

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	. "gopkg.in/check.v1"
)

const (
	file1 = "./test_data/test.json"
	file2 = "./test_data/simple.json"
)

type testPG struct{}

var _ = Suite(&testPG{})

func TestService(t *testing.T) { TestingT(t) }

func LoadFile(c *C, file string) []byte {

	jsonBlob, err := ioutil.ReadFile(file)
	c.Assert(err, IsNil)

	return jsonBlob
}

func LoadRawData(c *C, file string) map[string]interface{} {
	rawData := map[string]interface{}{}

	err := json.Unmarshal(LoadFile(c, file), &rawData)
	c.Assert(err, IsNil)

	return rawData
}

func load(c *C, file string) IXPath {
	return New(LoadRawData(c, file))
}

func (s *testPG) TestSimple01(c *C) {
	// c.Skip("no reason")

	path := "bookstore/book/title/lang"
	expected := "en"

	xp := load(c, file2)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestSimple02(c *C) {
	// c.Skip("no reason")

	expected := []interface{}{1.0, 2.0, 3.0}
	path := "cycle-1"

	xp := load(c, file2)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestSimple03(c *C) {
	// c.Skip("no reason")

	expected := []interface{}{"Everyday Norwegian", "Everyday French", "Everyday Spanish"}
	path := "cycle-2/book"

	xp := load(c, file2)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestSimple04(c *C) {
	// c.Skip("no reason")

	expected := []interface{}{[]interface{}{"Everyday Norwegian"}, []interface{}{"Everyday French"}, []interface{}{"Everyday Spanish"}}
	path := "cycle-3/book"

	xp := load(c, file2)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestSimple05(c *C) {
	// c.Skip("no reason")

	expected := []interface{}{"Everyday Norwegian", "Everyday French", "Everyday Spanish"}
	path := "cycle-4/book/title"

	xp := load(c, file2)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestBook01(c *C) {
	// c.Skip("no reason")

	expected := []interface{}{"Everyday Italian", "Harry Potter", "XQuery Kick Start", "Learning XML"}
	path := "/bookstore/book/title/value"

	xp := load(c, file1)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestBook02(c *C) {
	// c.Skip("no reason")
	xp := load(c, file1)
	expected := []interface{}{"Everyday Italian"}
	path := "/bookstore/book[0]/title/value"
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	fmt.Printf("\n\nTestBook02. res: %+v\n\n", res)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestBook03(c *C) {
	c.Skip("no reason")
	xp := load(c, file1)
	expected := []interface{}{"XQuery Kick Start", "Learning XML"}
	path := "/bookstore/book[price>35]/title/value"
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	fmt.Printf("\n\nTestBook02. res: %+v\n\n", res)

	c.Assert(res, DeepEquals, expected)
}

/*

func (s *testPG) TestSimple03(c *C) {
	c.Skip("skip")

	xp := load(c)
	path := "/bookstore/book/price[text()]"
	c.Assert(xp.Get(path), DeepEquals, []float64{30.00, 29.99, 49.99, 39.95})
	c.Assert(xp.First(path), DeepEquals, 30.00)
}
*/
