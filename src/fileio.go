package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func readAndLoadFiles(fp string, db *sql.DB, fn string, dir string, err error) error {
	filepath.Walk(fp, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err.Error())
		}
		isPolicy, _ := filepath.Match("*yaml*", info.Name())

		policy := Policy{}

		if isPolicy {
			dat, err := ioutil.ReadFile(filepath.Dir(path) + "/" + info.Name())
			CheckIfError(err)

			if err != nil {
				log.Fatal(err)
			}

			//parse default as cilium policy
			err2 := yaml.Unmarshal(dat, &policy) // right now only unmarshalling only to network policy...

			if err2 != nil {

				log.Println("Umarshall error for policy ", err2)
				fmt.Println(string(dat))
				//should skip this policy
			}
			if err := insert(policy, db, fn, dir); err != nil {
				log.Println(err)
				os.Exit(1)
			}
		}
		return nil
	})
	return nil
}
