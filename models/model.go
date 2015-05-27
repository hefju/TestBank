package models
import (
    _ "github.com/mattn/go-sqlite3"
    "github.com/go-xorm/core"
    "github.com/go-xorm/xorm"
    "log"

)
var engine *xorm.Engine
func init() {
    var err error
    //	engine, err = xorm.NewEngine("odbc", "driver={SQL Server};Server=192.168.1.200; Database=charge; uid=sa; pwd=123;")
    engine, err = xorm.NewEngine("sqlite3", "./TestBank.db")

    if err != nil {
        log.Fatalln("xorm create error", err)
    }
   // engine.ShowSQL = true
    engine.SetMapper(core.SameMapper{})
    // engine.CreateTables(new(tp_charge_billing))
     err = engine.Sync2(new(Topics)) //, new(Group))
    if err != nil {
        log.Fatalln("xorm sync error", err)
    }
}

