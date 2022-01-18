package asset

import (
	`github.com/reloclass/core/cloudfile`
)

func (g *GetPreviewUrlReq) CloudDirType() cloudfile.DirType {
	return cloudfile.DirType(g.DirType)
}
