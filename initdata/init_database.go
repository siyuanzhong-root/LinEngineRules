package initdata

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path/filepath"
)

/**
CREATE_USER:SYZ
CREATE_TIME:2023/4/18 16:31
CREATE_BY:GoLand.LinEngineRules
*/

var EngineDB *gorm.DB

// init 初始化数据库
func init() {
	// 数据库初始化
	var err error
	dir := os.Getenv("DATABASE_DIR")
	EngineDB, err = gorm.Open(sqlite.Open(fmt.Sprintf(filepath.Join(dir, "engine.db"))), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("open engine.db err:%v", err))
	}
	EngineDB.Logger = logger.Default.LogMode(logger.Info)

}
