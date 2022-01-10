package filechunk

import (
	`github.com/reloclass/core/cloudfile`
)

func (i *InitReq) CloudDirType() cloudfile.DirType {
	return cloudfile.DirType(i.Type)
}
