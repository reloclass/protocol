package disk

import (
	`github.com/reloclass/core/cloudfile`
	`github.com/storezhang/gox`
)

func (r *RenameReq) CloudFileType() cloudfile.Type {
	return cloudfile.Type(r.Type)
}

func (r *RenameReq) GoxFileType() gox.FileType {
	return gox.FileType(r.FileType)
}
