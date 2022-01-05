package disk

import (
	`github.com/reloclass/core/cloudfile`
)

func (g *UploadReq) CloudFileType() cloudfile.Type {
	return cloudfile.Type(g.Type)
}
