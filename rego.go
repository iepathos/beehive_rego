// rego contains RethinkDB Go helper functions
package rego

import (
	"log"
	"os"

	r "gopkg.in/gorethink/gorethink.v3"
)

func CreateDatabase(databaseName string) {
	// connect to rethinkdb
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Creating database", databaseName)
	_, err = r.DBCreate(databaseName).Run(session)
	if err != nil {
		log.Println(err.Error())
	}
}

func CreateTable(tableName string) {
	// connect to rethinkdb
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	db := r.DB(os.Getenv("DBNAME"))

	log.Println("Creating table", tableName)
	if _, err := db.TableCreate(tableName).RunWrite(session); err != nil {
		log.Println(err)
	}
}
