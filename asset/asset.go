package asset

import (
	`github.com/storezhang/uoa`
)

func (g *GetPreviewUrlReq) UoaType() uoa.Type {
	return uoa.Type(g.DirType)
}
