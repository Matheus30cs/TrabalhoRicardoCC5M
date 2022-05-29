function convert(){
  //pegar o valor
  var quant = document.getElementById("quantia");
  var valorConverter = quant.value; //Recebe o valor o valor digitado pelo usuario;

  //pegar a moeda de entrada
  var moedaIn = document.querySelector('input[name="in"]:checked');
  var mIn = moedaIn.value;  //pega o valor selecionado pelo usuario.

  //pegar a moeda de saida
  var moedaOut = document.querySelector('input[name="out"]:checked');
  var mOut = moedaOut.value; //pega o valor selecionado pelo usuario.

  var valorEmReal = 0;
  var valorConvertido = 0;
  //converter
  if(mIn == "dolar"){
    valorEmReal = valorConverter*5.46;
  }else if(mIn == "euro"){
    valorEmReal = valorConverter*5.89;
  }else if(mIn == "libra"){
    valorEmReal = valorConverter*6.75;
  }else{
    valorEmReal = valorConverter;
  }

  if(mOut == "dolar"){
    valorConvertido = valorEmReal/5.46;
  }else if(mOut == "euro"){
    valorConvertido = valorEmReal/5.89;
  }else if(mOut == "libra"){
    valorConvertido = valorEmReal/6.75;
  }else{
    valorConvertido = valorEmReal;
  }

  //exibir o resultado
  var outRes = document.getElementById("resultado");
  outRes.innerHTML = valorConvertido;
}