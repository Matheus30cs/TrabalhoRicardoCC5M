package models

import (
	"errors"
  "time"
)

var ErrNoRecord = errors.New("models: no matching record Found")

type Produto struct{
  ID int
  NOME string
  PRECO float64
  Created time.Time
  Expires time.Time
}