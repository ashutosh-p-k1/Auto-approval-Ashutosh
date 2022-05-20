package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

func checkOut(db *sql.DB, fn string, dir string) error {

	//cleaning up of existing folders

	conf := confReader() //opening conf here first
	fmt.Println("The base0 dir is :=>", conf.Basedir)

	conf.Basedir += stringWithCharset(10)

	// Clone the given repository to the given directory
	//username : "ashutoshpandeyengineer"
	//password : "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"
	if conf.Username != "" && conf.Password != "" { //if username and token is available
		/*//TODO: add git authentication
		_, err = git.PlainClone(conf.b_dir, true, &git.CloneOptions{
			URL: conf.b_url,
			Auth: &http.BasicAuth{
				Username: conf.u, // anything except an empty string
				Password: conf.t,
			},
		})
		if err != nil {
			log.Println("Error in checkout is :-> ", err)
		}
		fmt.Println("satisfied:")
		*/
	} else {

		log.Print("\n Initiating a main repo clone ", conf.Baseurl)

		r, err := git.PlainClone(conf.Basedir, false, &git.CloneOptions{
			URL:      conf.Baseurl,
			Progress: os.Stdout,
		})

		CheckIfError(err)

		ref, err2 := r.Head()
		log.Print("\n\n The head is pointing to:.. ", ref.Hash())

		CheckIfError(err2)
		//traverse through files.
		if err := readAndLoadFiles(conf.Basedir, db, fn, dir, err); err != nil {
			return err
		}

		//load them into a btree

	} //end of else

	return nil
}
