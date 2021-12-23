package filechunk

import `github.com/storezhang/uoa`

func (i *InitReq) UoaType() uoa.Type {
	return uoa.Type(i.Type)
}
