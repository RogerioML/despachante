package despachante

import (
	"database/sql"
	"log"
)

type Telegrama struct {
	Codigo string
}

func consultaTelegrama(con sql.DB) (Telegrama, error) {
	var telegrama Telegrama

	qry := `SELECT TOP 1 NUM_TX_REGISTRO_SRO FROM LOGTELEGRAMASSGM`

	row := con.QueryRow(qry)
	if err := row.Scan(&telegrama.Codigo); err != nil {
		log.Fatal(err.Error())
		return telegrama, err
	}

	return telegrama, nil

}
