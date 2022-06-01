package main

import "net/http"

func (app *application) routes() *http.ServeMux {
  mux := http.NewServeMux()
  mux.HandleFunc("/", app.home)
  mux.HandleFunc("/Produtos", app.Produtos)
  mux.HandleFunc("/Produtos/Fogao", app.Fogao)
  mux.HandleFunc("/Produtos/Fogao/Compra", app.CompraFogao)
  mux.HandleFunc("/Produtos/Geladeira", app.Geladeira)
  mux.HandleFunc("/Produtos/Geladeira/Compra", app.CompraGeladeira)
  mux.HandleFunc("/Produtos/Lavadora", app.Lavadora)
  mux.HandleFunc("/Produtos/Lavadora/Compra", app.CompraLavadora)
  mux.HandleFunc("/Produtos/Microondas", app.Microondas)
  mux.HandleFunc("/Produtos/Microondas/Compra", app.CompraMicroondas)
  mux.HandleFunc("/Produtos/Compra/Confirmacao", app.ConfirmacaoCompra)
  mux.HandleFunc("/Produtos/ConversorMoedas", app.ConversorMoedas)
  mux.HandleFunc("/Produto", app.showProduto)
  mux.HandleFunc("/Cadastro", app.Cadastro)
  mux.HandleFunc("/Login", app.Login)
  mux.HandleFunc("/Produto/create", app.createProduto)
  mux.HandleFunc("/Produto/Lista", app.PaginaListaProdutos)
  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/",http.StripPrefix("/static", fileServer))

  return mux
}