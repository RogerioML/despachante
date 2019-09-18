package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	telegrama "github.com/RogerioML/despachante"
)

var (
	server   = "10.193.90.42"
	port     = 1433
	user     = "DSAWAPLDES"
	password = "ds@w@pldes"
	database = "DBTLMAIL"
)

func main() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("erro ao abrir conexao: ", err.Error())
	}
	if err = conn.Ping(); err != nil {
		log.Fatal("nao foi possivel conectar ao banco")
	}
	defer conn.Close()

	t, err := telegrama.consultaTelegrama(conn)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(t.Codigo)

}
