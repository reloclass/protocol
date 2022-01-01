package cloudfile

import `github.com/storezhang/uoa`

// Dir 文件信息
type Dir struct {
	// 学校编码
	SchoolId int64
	// 文件ID
	FileId string
	// 用户ID
	UserId int64
	// 文件类型
	Type uoa.Type
}
