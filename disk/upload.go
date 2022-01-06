package disk

import (
	`github.com/reloclass/core/cloudfile`
)

func (g *UploadReq) CloudDirType() cloudfile.DirType {
	return cloudfile.DirType(g.Type)
}
