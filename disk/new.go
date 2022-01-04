package disk

import `github.com/storezhang/uoa`

func (g *NewReq) UoaType() uoa.Type {
	return uoa.Type(g.Type)
}
