package model

import (
	"github.com/Cepave/fe/g"
	"github.com/Cepave/fe/model/dashboard"
	"github.com/Cepave/fe/model/uic"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() {
	// set default database
	config := g.Config()
	orm.RegisterDataBase("default", "mysql", config.Uic.Addr, config.Uic.Idle, config.Uic.Max)
	orm.RegisterDataBase("graph", "mysql", config.GraphDB.Addr, config.GraphDB.Idle, config.GraphDB.Max)
	orm.RegisterDataBase("falcon_portal", "mysql", config.FalconPortal.Addr, config.FalconPortal.Idle, config.FalconPortal.Max)

	// register model
	orm.RegisterModel(new(uic.User), new(uic.Team), new(uic.Session), new(uic.RelTeamUser), new(dashboard.Endpoint), new(dashboard.EndpointCounter), new(dashboard.HostGroup), new(dashboard.Hosts))

	if config.Log == "debug" {
		orm.Debug = true
	}
}
