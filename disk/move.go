package disk

import (
	`github.com/reloclass/core/cloudfile`
)

func (g *MoveReq) CloudDirType() cloudfile.DirType {
	return cloudfile.DirType(g.Type)
}
