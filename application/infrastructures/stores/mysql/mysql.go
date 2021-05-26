package mysql

import (
	"fmt"

	"github.com/go-xorm/xorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var db *xorm.Engine

func Init() error {
	var err error
	var cfgs = []string{"username", "password", "address", "databases"}
	var cfgVals = make([]interface{}, 0)
	for _, cfg := range cfgs {
		cfgVals = append(cfgVals, viper.GetString("mysql."+cfg))
	}
	db, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", cfgVals...))
	if db != nil {
		logrus.Info(db.Query("show tables"))
	}
	return err
}
