package disk

import (
	`github.com/reloclass/core/cloudfile`
)

func (g *NewReq) CloudFileType() cloudfile.Type {
	return cloudfile.Type(g.Type)
}
