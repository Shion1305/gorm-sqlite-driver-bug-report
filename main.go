package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	var g = gen.NewGenerator(gen.Config{
		OutPath:           "./query",
		ModelPkgPath:      "./domain",
		FieldNullable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	sqliteConn := sqlite.Open("db.sqlite3")
	db, err := gorm.Open(sqliteConn, &gorm.Config{})
	if err != nil {
		return
	}

	g.UseDB(db)
	all := g.GenerateAllTable() // database to table model.
	g.ApplyBasic(all...)

	// Generate the code
	g.Execute()
}
