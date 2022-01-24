package utils

var (
	SUCCESS = 200
	ERROR   = 500

	// Controller Error
	ParamsResolveFault = 1001
)

var codeMsg = map[int]string{
	SUCCESS: "操作成功",
	ERROR:   "操作失败",

	ParamsResolveFault: "参数解析失败",
}

func GetErrorMessage(code int) string {
	return codeMsg[code]
}
