package gopostgres

import (
	"io/ioutil"
	"runtime"
	"strings"
)

//This function receive path of a sql file and executes it if not done before
//
//THIS FUNCTION IS NOT COMPLETELY FOOL PROOF, USE IT AT YOUR OWN RISK
func (db Postgres) Migrate(path string) {
	db.createMigrationTableIfNotExists()
	list := migrationFilesList(path)
	db.migrateFiles(list)
}

func migrationFilesList(path string) []string {
	contentList, _ := ioutil.ReadDir(path)
	var list []string
	for _, f := range contentList {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".sql") {
			list = append(list, path +  "/" + f.Name())
		}
	}

	return list
}

func (db Postgres) createMigrationTableIfNotExists() error {
	exists, _ := db.tableExists("migrations")
	if exists {
		return nil
	}

	_, f, _, _ := runtime.Caller(0)
	path := strings.TrimSuffix(f, "/migrate.go") + "/sql/create_migrations_table.sql"

	query, err := getQueryFromSQLFile(path)
	if err != nil {
		return err
	}
	_, err = db.execQuery(*query)
	if err != nil {
		return err
	}
	return nil
}

func (db Postgres) migrateFiles(filePaths []string) error {
	migratedPaths, err := db.getMigratedFiles()
	if err != nil {
		return err
	}
	var queries []query
	for _, path := range filePaths {
		if !stringSliceContains(migratedPaths, path) {
			q, err := getQueryFromSQLFile(path)
			if err != nil {
				return err
			}
			queries  = append(queries, *q)

			migrationQuery := query{
				Statement: queryAddMigration,
				Args:      []interface{}{path},
			}
			queries = append(queries, migrationQuery)
		}
	}

	err = db.execQueriesWithTransactions(queries)
	if err != nil {
		return err
	}

	return err
}

func (db Postgres) getMigratedFiles() ([]string, error) {
	res, err := db.FindAllWhere("migrations", []string{"path"})
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, r := range res {
		paths = append(paths, r[0].(string))
	}

	return paths, nil
}

func stringSliceContains(stringSlice []string, stringToSearchFor string) bool {
	for _, s := range stringSlice {
		if s == stringToSearchFor {
			return true
		}
	}

	return false
}