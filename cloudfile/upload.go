package cloudfile

import `github.com/storezhang/uoa`

func (u *UploadReq) GetDirType() uoa.Type {
	return uoa.Type(u.DirType)
}
