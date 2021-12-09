package disk

import `github.com/storezhang/gox`

func (r *RenameReq) GoxFileType() gox.FileType {
	return gox.FileType(r.Type)
}
