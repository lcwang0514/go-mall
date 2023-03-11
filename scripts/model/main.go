package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

const dsn = "admin:password@tcp(localhost:3306)/gomall?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	 db, err := gorm.Open(mysql.Open(dsn))
	 if err != nil {
		panic(fmt.Errorf("cannot establish db connection: #{err}"))
	 }

	 g := gen.NewGenerator(gen.Config{
		OutPath: "../../source/server/logic/orm/dal",
		ModelPkgPath: "../../source/server/logic/orm/model",
		Mode: gen.WithDefaultQuery | gen.WithoutContext,
	 })

	 // 设置目标 db
	 g.UseDB(db)
	 g.ApplyBasic(g.GenerateAllTable()...)
	 g.ApplyBasic(g.GenerateModel("accounts", gen.FieldType("type", "sdk.StoreType")))
	 g.Execute()
}
