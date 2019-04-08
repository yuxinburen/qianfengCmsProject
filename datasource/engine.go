package datasource

import (
	_ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/go-xorm/xorm"
	"irisDemo/CmsProject/model"
)

/**
 * 实例化数据库引擎方法：mysql的数据引擎
 */
func NewMysqlEngine() *xorm.Engine {

	//数据库引擎
	engine, err := xorm.NewEngine("mysql", "root:yu271400@/qfCms?charset=utf8")

	//根据实体创建表
	//err = engine.CreateTables(new(model.Admin))

	//同步数据库结构：主要负责对数据结构实体同步更新到数据库表
	/**
	 * 自动检测和创建表，这个检测是根据表的名字
	 * 自动检测和新增表中的字段，这个检测是根据字段名，同时对表中多余的字段给出警告信息
	 * 自动检测，创建和删除索引和唯一索引，这个检测是根据索引的一个或多个字段名，而不根据索引名称。因此这里需要注意，如果在一个有大量数据的表中引入新的索引，数据库可能需要一定的时间来建立索引。
	 * 自动转换varchar字段类型到text字段类型，自动警告其它字段类型在模型和数据库之间不一致的情况。
	 * 自动警告字段的默认值，是否为空信息在模型和数据库之间不匹配的情况
	 */
	//Sync2是Sync的基础上优化的方法
	err = engine.Sync2(new(model.Permission), new(model.City), new(model.Admin), new(model.AdminPermission), new(model.User))

	if err != nil {
		panic(err.Error())
	}

	//设置显示sql语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)

	return engine
}
