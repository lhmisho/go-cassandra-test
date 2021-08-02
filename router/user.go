package router

import (
	"cassandraTest/database/scyllabdb"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandleUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	queryValues := r.URL.Query()
	username := queryValues.Get("username")
	userCity := queryValues.Get("usercity")
	fmt.Printf("Username %s\n", username)
	fmt.Printf("User city %s\n", userCity)

	var session = scyllabdb.GetSes()
	err := session.Query("INSERT INTO test_ks.user_dtl(user_id, username, city) VALUES(1, 'cstest', 'dhaka');").Exec()
	fmt.Println(err)
	//defer session.Close()


}
