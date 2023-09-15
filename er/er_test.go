package er_test

import (
	"github.com/go-courier/sqlx/v2/er"
	"github.com/go-courier/sqlx/v2/generator/__examples__/database"
	"github.com/go-courier/sqlx/v2/postgresqlconnector"
	"github.com/kunlun-qilian/utils/json"
)

func ExampleDatabaseERFromDB() {
	ers := er.DatabaseERFromDB(database.DBTest, &postgresqlconnector.PostgreSQLConnector{})
	_, _ = json.MarshalIndent(ers, "", "  ")
	// Output:
}
