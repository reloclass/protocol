package disk

import `github.com/storezhang/uoa`

func (g *UploadReq) UoaType() uoa.Type {
	return uoa.Type(g.Type)
}
