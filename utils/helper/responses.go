package helper

func Fail_Resp(msg string) map[string]interface{} {
	return map[string]interface{}{
		"Status":  "FAILED",
		"Message": msg,
	}

}

func Success_Resp(msg string) map[string]interface{} {
	return map[string]interface{}{
		"Status":  "SUCCESS",
		"Message": msg,
	}

}

func Success_DataResp(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"Status":  "SUCCESS",
		"Message": msg,
		"Data":    data,
	}

}
