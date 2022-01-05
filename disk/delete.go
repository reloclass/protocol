package disk

import (
	`github.com/reloclass/core/cloudfile`
)

func (g *DeleteReq) CloudFileType() cloudfile.Type {
	return cloudfile.Type(g.Type)
}
