package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestColumnTypes(t *testing.T) {
	tc := []struct {
		name              string
		query             string
		expectedIsPrimary map[string]bool
	}{
		{
			name:  "What should be expected",
			query: "CREATE TABLE something\n(\n    column1 INT,\n    column2 INT,\n    column3 INT,\n    PRIMARY KEY (column1, column2)\n);",
			expectedIsPrimary: map[string]bool{
				"column1": true,
				"column2": true,
				"column3": false,
			},
		},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {

			// create a new db for sqlite
			dbName := fmt.Sprintf("db_%s.sqlite3", uuid.NewString())
			sqliteConn := sqlite.Open(dbName)
			db, err := gorm.Open(sqliteConn, &gorm.Config{})
			if err != nil {
				t.Errorf("Failed to open new db: %s", err)
			}

			// run table creation query
			db.Exec(tt.query)
			types, err := db.Migrator().ColumnTypes("something")
			if err != nil {
				t.Errorf("Failed to get column types: %s", err)
			}
			for _, ty := range types {
				pr, valid := ty.PrimaryKey()
				exp := tt.expectedIsPrimary[ty.Name()]
				if pr && valid && exp || !pr && !exp {
					continue
				}
				t.Errorf("Expected %s to be primary key: %t, got: isPrimaryKey: %t, isValid: %t",
					ty.Name(), exp, pr, valid)
			}
		})
	}
}
