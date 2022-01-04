package user

const (
	// LoginSchool 后台登陆
	LoginSchool LoginType = "1"
	// LoginClient 客户端登陆
	LoginClient LoginType = "2"
)

// LoginType 登陆类型
type LoginType string
