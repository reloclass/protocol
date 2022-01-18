package disk

import (
	`github.com/reloclass/core/cloudfile`
	`github.com/storezhang/gox`
)

func (r *RenameReq) CloudDirType() cloudfile.DirType {
	return cloudfile.DirType(r.Type)
}

func (r *RenameReq) GoxFileType() gox.FileType {
	return gox.FileType(r.FileType)
}
