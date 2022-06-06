package main

//go run cmd/web/*
import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

  "main/pkg/models"
)

func(app *application) home (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/index.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) Produtos (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }
  
  files := []string{
    "./ui/html/produtos.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) Fogao (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos/Fogao"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }
  
  files := []string{
    "./ui/html/Fogao.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) Geladeira (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos/Geladeira"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/Geladeira.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) Lavadora (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos/Lavadora"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/Lavadora.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) Microondas (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos/Microondas"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/Microondas.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) CompraFogao (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos/Fogao/Compra"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/CompraFogao.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) CompraGeladeira (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos/Geladeira/Compra"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/CompraGeladeira.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) CompraLavadora (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos/Lavadora/Compra"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/CompraLavadora.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) CompraMicroondas (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos/Microondas/Compra"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/CompraMicroondas.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) ConfirmacaoCompra (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos/Compra/Confirmacao"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/ConfirmacaoCompra.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}


func(app *application) Cadastro (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Cadastro"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/Cadastro.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}


func(app *application) Login (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Login"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/Login.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}





func(app *application) ConversorMoedas (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos/ConversorMoedas"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/conversorMoedas.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}







func(app *application) PaginaListaProdutos (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produto/Lista"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/home.page.tmpl.html",
    "./ui/html/base.layout.tmpl.html",
    "./ui/html/footer.partial.tmpl.html",   
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}


//http://localhost:4000/snippet?id=123
func(app *application) showProduto(rw http.ResponseWriter, r *http.Request){
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    app.notFound(rw)
    return
  }

  s, err := app.produtos.Get(id)
  if err == models.ErrNoRecord {
    app.notFound(rw)
    return
  } else if err != nil{
    app.serverError(rw, err)
    return
  }
  
  files := []string{
    "./ui/html/show.page.tmpl.html",
    "./ui/html/base.layout.tmpl.html",
    "./ui/html/footer.partial.tmpl.html",   
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, s)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  
}

func(app *application) createProduto(rw http.ResponseWriter, r *http.Request){
  
  nome := "Aspirador de Pó"
  preco := 201.90
  created := "7"
  expires := "7"
  
  //id, err := app.snippets.Insert(title, content, expires)
  id, err := app.produtos.Insert(nome, preco, created, expires)
  if err != nil{
    app.serverError(rw, err)
    return
  }

  http.Redirect(rw, r, fmt.Sprintf("/Produto?id=%d", id),http.StatusSeeOther)

  if r.URL.Path != "/Produto/create"{
    app.notFound(rw)
    return
  }

  produtos, err := app.produtos.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/home.page.tmpl.html",
    "./ui/html/base.layout.tmpl.html",
    "./ui/html/footer.partial.tmpl.html",   
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, produtos)
  if err != nil{
    app.serverError(rw, err)
    return
  }

  //rw.Write([]byte("Criar novo Snippet"))
}
//curl -igp -X POST http://localhost:4000/snippet/create
//curl -i -X GET http://localhost:4000/snippet/create


/*
func(app *application) createSnippet(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    rw.Header().Set("Allow","POST")
    app.clientError(rw, http.StatusMethodNotAllowed)
    return
  }

  title := "Aula de hoje"
  content := "Tentando lidar com o Banco de Dados"
  expires := "7"

  id, err := app.snippets.Insert(title, content, expires)

  if err != nil{
    app.serverError(rw, err)
    return
  }

  http.Redirect(rw, r, fmt.Sprintf("/snippet?id=%d", id),http.StatusSeeOther)
  //rw.Write([]byte("Criar novo Snippet"))
}
//curl -igp -X POST http://localhost:4000/snippet/create
//curl -i -X GET http://localhost:4000/snippet/create*/