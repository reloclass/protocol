package filechunk

import `github.com/storezhang/uoa`

func (i *InitMultipartUploadReq) UoaType() uoa.Type {
	return uoa.Type(i.Type)
}
