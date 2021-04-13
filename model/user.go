/*
 * @Author: mengdaoshizhongxinyang
 * @Date: 2021-03-19 17:12:13
 * @Description:
 */
package model

import "Chatin/global"

type User struct {
	Name     string
	Account  string
	Avatar   string
	Password string
	Config 	 string
	global.BASE_MODEL
}
