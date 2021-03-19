/*
 * @Author: mengdaoshizhongxinyang
 * @Date: 2021-03-19 17:26:48
 * @Description:
 */
package logic

type MSG map[string]interface{}

func OK(data interface{}, msg string) MSG {
	return MSG{"data": data, "errCode": 0, "message": msg}
}
func ERROR(msg string, errCode int16) MSG {
	return MSG{"errCode": errCode, "message": msg}
}
