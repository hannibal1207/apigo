package beer_test

import (
	"Documentos/pos_trybe/go/api_go/apigo/core/beer"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestStore(t *testing.T) {
	b := &beer.Beer{
		ID:    1,
		Name:  "Heineken",
		Type:  beer.TypeLager,
		Style: beer.StylePale,
	}

	db, err := sql.Open("sqlite3", "../../data/beer_test.db")
	if err != nil {
		t.Fatalf("Erro conectando ao banco de dados %s", err.Error())
	}
	err = clearDB(db)
	if err != nil {
		t.Fatalf("Erro limpando o banco de dados: %s", err.Error())
	}
	defer db.Close()
	service := beer.NewService(db)
	err = service.Store(b)
	if err != nil {
		t.Fatalf("Erro salvando no banco de dados: %s", err.Error())
	}
	saved, err := service.Get(1)
	if err != nil {
		t.Fatalf("Erro buscando do banco de dados: %s", err.Error())
	}
	if saved.ID != 1 {
		t.Fatalf("Dados invalidos. Esperando %d, recebido %d", 1, saved.ID)
	}
}

func clearDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from beer")
	tx.Commit()
	return err
}
