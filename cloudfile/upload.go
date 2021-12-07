package cloudfile

import `github.com/storezhang/uoa`

func (u *UploadReq) UoaType() uoa.Type {
	return uoa.Type(u.DirType)
}
