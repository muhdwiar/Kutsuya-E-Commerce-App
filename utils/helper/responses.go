package helper

import "net/http"

func Fail_Resp(msg string) map[string]interface{} {
	return map[string]interface{}{
		"Status":  "Failed",
		"Message": msg,
	}

}

func Success_Resp(msg string) map[string]interface{} {
	return map[string]interface{}{
		"Status":  "Success",
		"Message": msg,
	}

}

func Success_DataResp(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"Status":  "Success",
		"Message": msg,
		"Data":    data,
		"code":    http.StatusOK,
	}

}
