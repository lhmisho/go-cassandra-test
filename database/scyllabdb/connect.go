package scyllabdb

import (
	"cassandraTest/pkg/utils"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"sync"
)

var (
	doOnce sync.Once
	session *gocql.Session
	err error
)

func GetSes() *gocql.Session{
	doOnce.Do(func() {
		cluster := gocql.NewCluster(utils.GoDotEnvVariable("SCYLLA_HOST"))
		cluster.Keyspace = utils.GoDotEnvVariable("SCYLLA_KEYSPACE")
		cluster.Consistency = gocql.Quorum
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: utils.GoDotEnvVariable("SCYLLA_USER"),
			Password: utils.GoDotEnvVariable("SCYLLA_PASSWORD"),
		}
		session, err = cluster.CreateSession()
		if err != nil {
			log.Printf("Session tracker connection error: %v", err)
		}else{
			fmt.Println("DB connected successfully")
		}
	})
	return session
}