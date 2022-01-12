package disk

import (
	`github.com/reloclass/core/cloudfile`
)

func (g *NewReq) CloudDirType() cloudfile.DirType {
	return cloudfile.DirType(g.Type)
}
