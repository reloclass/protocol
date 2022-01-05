package cloudfile

import (
	`github.com/reloclass/core/cloudfile`
)

func (u *UploadReq) CloudFileType() cloudfile.Type {
	return cloudfile.Type(u.DirType)
}
