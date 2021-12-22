package disk

import `github.com/storezhang/uoa`

func (g *GetFilesReq) UoaType() uoa.Type {
	return uoa.Type(g.Type)
}

func (g *GetDiskReq) UoaType() uoa.Type {
	return uoa.Type(g.Type)
}
