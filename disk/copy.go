package disk

import `github.com/storezhang/uoa`

func (g *CopyReq) UoaType() uoa.Type {
	return uoa.Type(g.Type)
}
