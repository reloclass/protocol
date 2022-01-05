package disk

import (
	`github.com/reloclass/core/cloudfile`
)

func (g *GetFilesReq) CloudFileType() cloudfile.Type {
	return cloudfile.Type(g.Type)
}

func (g *GetDiskReq) CloudFileType() cloudfile.Type {
	return cloudfile.Type(g.Type)
}
