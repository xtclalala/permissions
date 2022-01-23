package utils

var (
	SUCCESS = 200
	ERROR   = 500
)

var codeMsg = map[int]string{
	SUCCESS: "操作成功",
	ERROR:   "操作失败",
}

func GetErrorMessage(code int) string {
	return codeMsg[code]
}
