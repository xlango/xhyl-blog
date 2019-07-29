package models


type Token struct {
	Token string `json:"token"`
}

type Result struct {
	Code int
	Msg  interface{}
}

//返回状态码，200
func Ok(rs interface{})  Result{
	return Result{
		Code:200,
		Msg:rs,
	}
}

//认证失败，401
func AuthErr()  Result{
	return Result{
		Code:401,
		Msg:"[error] Authentication failed",
	}
}

//参数错误，403
func ParamErr()  Result{
	return Result{
		Code:403,
		Msg:"[error] Request param failed",
	}
}

//请求页面找不到，404
func Nofound()  Result{
	return Result{
		Code:404,
		Msg:"[error] No found",
	}
}

//服务器内部错误，500
func Error(err interface{})  Result{
	return Result{
		Code:500,
		Msg:err,
	}
}