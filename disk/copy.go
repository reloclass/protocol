package disk

import (
	`github.com/reloclass/core/cloudfile`
)

func (g *CopyReq) CloudFileType() cloudfile.Type {
	return cloudfile.Type(g.Type)
}
