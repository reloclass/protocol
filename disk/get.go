package disk

import (
	`github.com/reloclass/core/cloudfile`
)

func (g *GetFilesReq) CloudDirType() cloudfile.DirType {
	return cloudfile.DirType(g.Type)
}

func (g *GetDiskReq) CloudFileType() cloudfile.DirType {
	return cloudfile.DirType(g.Type)
}
