package filechunk

import (
	`github.com/reloclass/core/cloudfile`
)

func (i *InitReq) CloudFileType() cloudfile.Type {
	return cloudfile.Type(i.Type)
}
