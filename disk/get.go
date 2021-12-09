package disk

import (
	`github.com/storezhang/uoa`
)

func (g *GetDirFileReq) UoaType() uoa.Type {
	return uoa.Type(g.Type)
}
