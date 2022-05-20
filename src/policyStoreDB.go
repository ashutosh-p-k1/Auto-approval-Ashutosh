package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var err error
var count int
var sCount int32
var dbTempDir string
var schema = `
CREATE TABLE IF NOT EXISTS policy(
	id INTEGER, 
	apiversion varchar(20), 
	kind varchar(20), 
	name varchar(20), 
	namespace varchar(20)
	);
CREATE TABLE IF NOT EXISTS policy_spec(
	id INTEGER , 
	policy_id int,  
	label_key varchar(100), 
	label_value varchar(100), 
	spec varchar(2000));
`

func createDB() (db *sql.DB, fn string, dir string, err error) {
	count = 0
	sCount = 0
	dir, _ = ioutil.TempDir("/mnt/p/Auto approval development/2/Auto-approval-Ashutosh/temp", "dbtemp-")
	fn = filepath.Join(dir, "db")

	fmt.Println("The directory name is :->", dbTempDir)
	fmt.Println("The filepath is :->", fn)

	db, err = sql.Open("sqlite", fn)
	if err != nil {
		return db, fn, dir, err
	}
	_, err = db.Exec(schema)
	if err != nil {
		fmt.Println(err)
		return db, fn, dir, err
	}

	return db, fn, dir, nil
}
func insert(policy Policies, db *sql.DB, fn string, dir string) error {
	count += 1
	db.Ping()
	s := policy.GetSpecData()
	sqlstring := ""
	sqlstring = `insert into policy(id,apiversion,kind,name,namespace) values (` + fmt.Sprint(count) + `,'` + policy.GetApiVersion() + `','` + string(policy.GetKind()) + `','` + string(policy.GetMetaData()["name"]) + `','` + string(policy.GetMetaData()["namespace"]) + `')`

	if _, err = db.Exec(sqlstring); err != nil {
		fmt.Println("Error in SQL Exection in primary policies table")
		return err
	}

	for key, element := range policy.GetSpecSelector().MatchLabels {

		//insert every policy spec for matching label and key as a separate row to make searching easier.
		//increment sCount which is the primary key id for the spec table. Retain the count value which is the primary key for the policies table.
		sCount += 1
		sqlstring = `INSERT INTO policy_spec values (` + fmt.Sprint(sCount) + `,` + fmt.Sprint(count) + `,'` + element + `','` + key + `','` + fmt.Sprint(string(s)) + `')`
		if _, err = db.Exec(sqlstring); err != nil {
			return err
		}
	}

	return nil
}
