package disk

import (
	`github.com/storezhang/gox`
	`github.com/storezhang/uoa`
)

func (r *RenameReq) UoaType() uoa.Type {
	return uoa.Type(r.Type)
}

func (r *RenameReq) GoxFileType() gox.FileType {
	return gox.FileType(r.FileType)
}
