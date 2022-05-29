package main

import (
	"database/sql"
	"flag"
	"log"
	"main/pkg/models/mysql"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct{
  errorLog *log.Logger
  infoLog *log.Logger
  produtos *mysql.ProdutoModel
}

//curl -i -X POST http://localhost:4000/snippet/create

func main() {
            //Nome da FLag, Valor padrão e Descrição
  addr := flag.String("addr", ":4000", "Porta da Rede")
  dsn := flag.String("dsn", "8VejBck38E:8q4dNy1ziK@tcp(remotemysql.com)/8VejBck38E?parseTime=true", "MySql DSN")
  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERRO:\t", log.Ldate|log.Ltime|log.Lshortfile)

  db, err := openDB(*dsn)
  if err != nil{
    errorLog.Fatal(err)
  }
  defer db.Close()
  
  app := &application{
    errorLog: errorLog, infoLog: infoLog, produtos: &mysql.ProdutoModel{DB:db},
  }

  srv := &http.Server{
    Addr: *addr, ErrorLog: errorLog, Handler: app.routes(),
  }
  
  infoLog.Printf("Inicializando o servidor na porta %s", *addr)
  err = srv.ListenAndServe()
  errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error){
  db, err := sql.Open("mysql", dsn)
  if err!= nil{
    return nil, err
  }
  if err = db.Ping(); err != nil{
    return nil, err
  }
  return db, nil
}