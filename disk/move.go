package disk

import (
	`github.com/reloclass/core/cloudfile`
)

func (g *MoveReq) CloudFileType() cloudfile.Type {
	return cloudfile.Type(g.Type)
}
