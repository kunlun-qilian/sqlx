package er_test

import (
	"github.com/kunlun-qilian/sqlx/v3/er"
	"github.com/kunlun-qilian/sqlx/v3/generator/__examples__/database"
	"github.com/kunlun-qilian/sqlx/v3/postgresqlconnector"
	"github.com/kunlun-qilian/utils/json"
)

func ExampleDatabaseERFromDB() {
	ers := er.DatabaseERFromDB(database.DBTest, &postgresqlconnector.PostgreSQLConnector{})
	_, _ = json.MarshalIndent(ers, "", "  ")
	// Output:
}
