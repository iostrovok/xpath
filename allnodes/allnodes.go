package allnodes

import (
	"context"

	"github.com/iostrovok/xpath/convert"
)

type AllNodes struct {
	data  interface{}
	chRes chan interface{}
	ctx   context.Context
}

func New(data interface{}) *AllNodes {
	out := &AllNodes{
		data:  data,
		ctx:   context.Background(),
		chRes: make(chan interface{}, 2),
	}

	go getAllNodes(out.ctx, out.data, out.chRes)

	return out
}

func (a *AllNodes) Next() (interface{}, bool) {

	select {
	case <-a.ctx.Done():
		return nil, false
	case res, ok := <-a.chRes:
		return res, ok
	}

	return nil, false
}

func getAllNodes(ctx context.Context, d interface{}, chRes chan interface{}) {
	defer close(chRes)
	oneNode(ctx, d, chRes)
}

func oneNode(ctx context.Context, d interface{}, chRes chan interface{}) {
	if m, find := convert.IsStringMap(d); find {

		if pushToCh(ctx, m, chRes) {
			for _, v := range m {
				oneNode(ctx, v, chRes)
			}
		}

		return
	}

	if m, find := convert.IsArray(d); find {

		if pushToCh(ctx, m, chRes) {
			for _, v := range m {
				oneNode(ctx, v, chRes)
			}
		}

		return
	}
}

func pushToCh(ctx context.Context, m interface{}, chRes chan interface{}) bool {
	select {
	case <-ctx.Done():
		return false
	case chRes <- m:
		return true
	}

	return true
}
