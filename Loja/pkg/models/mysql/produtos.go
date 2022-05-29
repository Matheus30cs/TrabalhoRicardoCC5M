package mysql

import (
	"database/sql"
	"main/pkg/models"
)

type ProdutoModel struct{
  DB *sql.DB
}

func(m *ProdutoModel)Insert(nome string, preco float64, expires string) (int, error){
  stmt := `INSERT INTO produtos (nome, preco, created, expires) VALUES (?, ?, ?)`

  result, err := m.DB.Exec(stmt, nome, preco, expires)
  if err != nil{
    return 0, err
  }

  id, err := result.LastInsertId()
  
  if err != nil{
    return 0, err
  }

  return int(id), nil
}

func(m *ProdutoModel) Get(id int) (*models.Produto, error){

  stmt := `SELECT id, nome, preco, created, expires  FROM produtos
  WHERE expires > UTC_TIMESTAMP() AND id = ?`
  //WHERE preco > 0 AND id = ?`
  row := m.DB.QueryRow(stmt, id)

  s := &models.Produto{}

  err := row.Scan(&s.ID, &s.NOME, &s.PRECO, &s.Created, &s.Expires)
  
  if err == sql.ErrNoRows{
    return nil, models.ErrNoRecord
  } else if err != nil{
    return nil, err
  }
  
  return s, nil
}
func(m *ProdutoModel) Latest()([]*models.Produto, error){
  stmt := `SELECT id, nome, preco, created, expires  FROM produtos 
  WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
   //WHERE preco > 0 ORDER BY created DESC LIMIT 10`

  rows, err := m.DB.Query(stmt)
  if err != nil{
    return nil, err
  }
  defer rows.Close()

  produtos := []*models.Produto{}
  for rows.Next(){
    s := &models.Produto{}
    err = rows.Scan(&s.ID, &s.NOME, &s.PRECO, &s.Created, &s.Expires)
    if err != nil{
      return nil, err
    }
    produtos = append(produtos, s)
  }
  err = rows.Err()
  if err != nil{
    return nil, err
  }
  return produtos, nil
}