package xpath

import (
	. "gopkg.in/check.v1"
)

const (
	fileYahoo = "./test_data/yahoo.json"
)

func (s *testPG) TestYahoo01(c *C) {
	skipAll(c, "chart/result[0]/timestamp")

	path := "chart/result[0]/timestamp"
	expected := []interface{}{1560808560.0, 1560864600.0}

	xp := load(c, fileYahoo)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)
}

func (s *testPG) TestYahoo02(c *C) {
	// skipAll(c, "//result[0]/timestamp")

	path := "//result[0]/timestamp"
	expected := []interface{}{1560808560.0, 1560864600.0}

	xp := load(c, fileYahoo)
	res, err := xp.Get(path)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)

	c.Assert(res, DeepEquals, expected)
}
