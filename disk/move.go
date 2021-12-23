package disk

import `github.com/storezhang/uoa`

func (g *MoveReq) UoaType() uoa.Type {
	return uoa.Type(g.Type)
}
