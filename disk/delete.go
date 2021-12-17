package disk

import `github.com/storezhang/uoa`

func (g *DeleteReq) UoaType() uoa.Type {
	return uoa.Type(g.Type)
}
