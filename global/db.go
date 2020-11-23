package global

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var (
	DBLink *gorm.DB
)

func SetupDBLink() (error) {
	var err error
	DBLink, err = gorm.Open(DatabaseSetting.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			DatabaseSetting.UserName,
			DatabaseSetting.Password,
			DatabaseSetting.Host,
			DatabaseSetting.DBName,
			DatabaseSetting.Charset,
			DatabaseSetting.ParseTime,
	))
	if err != nil {
		return err
	}

	if ServerSetting.RunMode == "debug" {
		DBLink.LogMode(true)
	}
	DBLink.SingularTable(true)
	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	DBLink.DB().SetMaxIdleConns(DatabaseSetting.MaxIdleConns)
	DBLink.DB().SetMaxOpenConns(DatabaseSetting.MaxOpenConns)
	//otgorm.AddGormCallbacks(db)
	return nil
}

/*
func SetupDBLink() error {
	var err error
	DBLink, err = gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/dig?charset=utf8&parseTime=True&loc=Local")
	if err == nil  {
		// 全局禁用表名复数
		DBLink.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
		//打开sql日志
		DBLink.LogMode(true)
		return nil
	} else {
		return err
	}
}
*/