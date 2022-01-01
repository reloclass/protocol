package cloudfile

import `github.com/storezhang/uoa`

// 文件类型
const (
	// TypeProductFile 产品文件
	TypeProductFile uoa.Type = "product"
	// TypeProductRelease 产品发布文件
	TypeProductRelease uoa.Type = "product-release"
	// TypePublicDisk 公共云盘
	TypePublicDisk uoa.Type = "public"
	// TypePrivateDisk 私有云盘
	TypePrivateDisk uoa.Type = "private"
	// TypeFileResource 普通文件
	TypeFileResource uoa.Type = "resource"
	// TypeSystemFile 系统文件文件
	TypeSystemFile uoa.Type = "system"
	// TypeCourse 课程资源
	TypeCourse uoa.Type = "course"
	// TypeCourseTimeRecord 课程录课资源
	TypeCourseTimeRecord uoa.Type = "record"
	// TypeOrgRelease 版本发布文件
	TypeOrgRelease uoa.Type = "org-release"
)
