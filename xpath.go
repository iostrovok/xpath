package xpath

import (
	"fmt"

	"github.com/iostrovok/xpath/allnodes"
	"github.com/iostrovok/xpath/convert"
	"github.com/iostrovok/xpath/way"
)

type IXPath interface {
	Get(string) (interface{}, error)
	First(string) interface{}
}

type XPath struct {
	Data map[string]interface{}
}

func New(data map[string]interface{}) IXPath {

	a := XPath{
		Data: data,
	}

	_ = IXPath(a)

	return &a
}

func (xpath XPath) First(path string) interface{} {
	return nil
}

func (xpath XPath) Get(path string) (interface{}, error) {
	ways, err := way.New(path)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Get. Ways: Dump: %+v\n", ways[0].Dump())
	fmt.Printf("len(ways): %d\n", len(ways))

	if len(ways) == 1 {
		res, isArray := xpath._get(xpath.Data, ways[0], 0)

		fmt.Printf("Get: isArray: %t\n", isArray)
		fmt.Printf("Get: res: %+v\n", res)

		return res, nil
	}

	out := []interface{}{}

	// for _, way := range ways {
	// 	res := xpath._get(xpath.Data, way, 0)
	// 	fmt.Printf("Get: res: %+v\n", res)

	// 	if res != nil {
	// 		out = append(out, res...)
	// 	}

	// 	fmt.Printf("Get: out: %+v\n\n", out)

	// }

	// if len(out) == 0 {
	// 	return nil, nil
	// }

	return out, nil
}

func (xpath XPath) searchAny(data interface{}, way *way.Way) []interface{} {

	out := []interface{}{}

	allNodes := allnodes.New(data)
	node, find := allNodes.Next()
	for find {
		res, _ := xpath._get(node, way, 0)
		if res != nil {
			out = append(out, res)
		}
		node, find = allNodes.Next()
	}

	return out
}

func (xpath XPath) _get(data interface{}, way *way.Way, currentI int) (result interface{}, isArray bool) {

	result = nil
	isArray = false
	fmt.Printf("_get. Ways: Dump: %+v\n", way.Dump())

	findPath, findIteration := way.NextBy(currentI)

	if m, find := convert.IsStringMap(data); find {

		if !findIteration {
			result = data
			return
		}

		if value, find := m[findPath]; find {
			return xpath._get(value, way, currentI+1)
		}

		return
	}

	if m, find := convert.IsArray(data); find {

		isArray = true
		res := []interface{}{}

		if index, find := way.ArrayIndextBy(currentI - 1); find {

			if len(index) == 1 {
				indexID := index[0]
				if indexID < 0 {
					indexID = len(m) + indexID
					if indexID < 0 {
						return nil, false
					}
				}
				return xpath._get(m[indexID], way, currentI)
			}

			for _, indexID := range index {
				if indexID > len(m)-1 {
					break
				}

				if !findIteration {
					res = append(res, m[indexID])
				} else {
					resAdd, isSlice := xpath._get(m[indexID], way, currentI)
					res = appendSlice(res, resAdd, isSlice)
				}
			}

			return res, isArray
		}

		if rule, find := way.ArrayRuleBy(currentI - 1); find {

			for i := range m {

				if !rule(i, m[i]) {
					continue
				}
				if !findIteration {
					res = append(res, m[i])
				} else {
					resAdd, isSlice := xpath._get(m[i], way, currentI)
					res = appendSlice(res, resAdd, isSlice)
				}
			}

			return res, isArray
		}

		if !findIteration {
			return m, false
		}

		for i := range m {
			if !findIteration {
				res = append(res, m[i])
			} else {
				resAdd, isSlice := xpath._get(m[i], way, currentI)
				res = appendSlice(res, resAdd, isSlice)
			}
		}
		return res, isArray
	}

	if _, findNextIteration := way.NextBy(currentI + 1); !findNextIteration {
		result = data
	}

	return
}

func appendSlice(res []interface{}, resAdd interface{}, isSlice bool) []interface{} {

	if resAdd == nil {
		return res
	}

	if !isSlice {
		return append(res, resAdd)
	}

	if m, find := convert.IsArray(resAdd); find {
		res = append(res, m...)
	}

	return res
}
