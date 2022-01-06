package cloudfile

import (
	`github.com/reloclass/core/cloudfile`
)

func (u *UploadReq) CloudDirType() cloudfile.DirType {
	return cloudfile.DirType(u.DirType)
}
