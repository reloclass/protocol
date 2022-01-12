package disk

import (
	`github.com/reloclass/core/cloudfile`
)

func (g *DeleteReq) CloudDirType() cloudfile.DirType {
	return cloudfile.DirType(g.Type)
}
