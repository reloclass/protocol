package asset

import (
	`github.com/reloclass/core/cloudfile`
)

func (g *GetPreviewUrlReq) CloudFileType() cloudfile.Type {
	return cloudfile.Type(g.DirType)
}
