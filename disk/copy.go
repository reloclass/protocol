package disk

import (
	`github.com/reloclass/core/cloudfile`
)

func (g *CopyReq) CloudDirType() cloudfile.DirType {
	return cloudfile.DirType(g.Type)
}
