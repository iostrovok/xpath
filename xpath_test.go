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

func skipAll(c *C, reason string) {
	// c.Skip("skipAll " + reason)
}

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
	skipAll(c, "bookstore/book/title/lang")

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
	skipAll(c, "cycle-1")

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
	skipAll(c, "cycle-2/book")

	expected := []interface{}{"Everyday Norwegian", "Everyday French", "Everyday Spanish"}
	path := "cycle-2/book"

	xp := load(c, file2)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestSimple04(c *C) {
	c.Skip("no reason")
	skipAll(c, "cycle-3/book")

	expected := []interface{}{[]interface{}{"Everyday Norwegian"}, []interface{}{"Everyday French"}, []interface{}{"Everyday Spanish"}}
	path := "cycle-3/book"

	xp := load(c, file2)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestSimple05(c *C) {
	c.Skip("no reason")
	skipAll(c, "cycle-4/book/title")

	expected := []interface{}{"Everyday Norwegian", "Everyday French", "Everyday Spanish"}
	path := "cycle-4/book/title"

	xp := load(c, file2)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestSimple06(c *C) {
	// c.Skip("no reason")
	skipAll(c, "cycle-5/allbooks/titles")

	expected := []interface{}{[]interface{}{"Everyday Norwegian", "Everyday French", "Everyday Spanish"}}
	path := "cycle-5/allbooks/titles"

	xp := load(c, file2)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestSimple07(c *C) {
	// c.Skip("no reason")
	skipAll(c, "cycle-5/allbooks[0]/titles")

	expected := []interface{}{"Everyday Norwegian", "Everyday French", "Everyday Spanish"}
	path := "cycle-5/allbooks[0]/titles"

	xp := load(c, file2)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestBook01(c *C) {
	// c.Skip("no reason")
	skipAll(c, "/bookstore/book/title/value")

	expected := []interface{}{"Everyday Italian", "Harry Potter", "XQuery Kick Start", "Learning XML"}
	path := "/bookstore/book/title/value"

	xp := load(c, file1)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestBook02(c *C) {
	c.Skip("no reason")
	skipAll(c, "/bookstore/book[0]/title/value")

	expected := []interface{}{"Everyday Italian"}
	path := "/bookstore/book[0]/title/value"

	xp := load(c, file1)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)
	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestBook03(c *C) {
	c.Skip("/bookstore/book[price>35]/title/value")
	skipAll(c, "/bookstore/book[price>35]/title/value")

	expected := []interface{}{"XQuery Kick Start", "Learning XML"}
	path := "/bookstore/book[price>35]/title/value"

	xp := load(c, file1)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	fmt.Printf("\n\nTestBook02. res: %+v\n\n", res)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestBook04(c *C) {
	c.Skip("//title/value | //price...")
	skipAll(c, "//title/value | //price...")

	paths := []string{
		"//title/value | //price",
		"//book/title/value | //book/price",
		"/bookstore/book/title/value | //price",
	}
	expected := []interface{}{
		"Everyday Italian",
		30.00,
		"Harry Potter",
		29.99,
		"XQuery Kick Start",
		49.99,
		"Learning XML",
		39.95,
	}

	xp := load(c, file1)

	for _, path := range paths {
		res, err := xp.Get(path)
		c.Assert(err, IsNil)
		c.Assert(res, NotNil)

		c.Assert(res, DeepEquals, expected)
	}
}

func (s *testPG) TestBook05(c *C) {
	c.Skip("//title/value | //price...")
	skipAll(c, "//title/value | //price...")

	path := "//*"

	expected := []interface{}{
		"Everyday Italian",
		"Giada De Laurentiis",
		2005,
		30.00,

		"Harry Potter",
		"J K. Rowling",
		2005,
		29.99,

		"XQuery Kick Start",
		"James McGovern",
		"Per Bothner",
		"Kurt Cagle",
		"James Linn",
		"Vaidyanathan Nagarajan",
		2003,
		49.99,

		"Learning XML",
		"Erik T. Ray",
		2003,
		39.95,
	}

	xp := load(c, file1)

	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)

}
