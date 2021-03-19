/*
 * @Author: mengdaoshizhongxinyang
 * @Date: 2021-03-19 15:34:29
 * @Description:
 */
package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func GormMysql() *gorm.DB {

	dsn := "root:666666@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		//global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil
		return nil
	} else {
		return db
	}
}
