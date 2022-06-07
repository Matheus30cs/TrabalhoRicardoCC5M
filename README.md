# TrabalhoRicardoCC5M

• Descrição do projeto

O Projeto consiste em um site de Compra de Produtos Eletrodomésticos, o Nome do Site se chama KingEletros.
Basicamente o Site temos 4 abas:

1 - Home/Inicio
2 - Produtos
3 - Login
4 - Cadastro

Dentro de Produtos temos os Eletrodomésticos disponíveis na Loja. Abaixo da seleção dos produtos temos o Link "Eletrodomésticos no Banco de dados", que vai mostrar uma página com detalhes técnicos dos produtos registrados no Banco de Dados.

Nessa página temos as abas: 

1 - Home - Que vai levar de volta para a página principal do site.
2 - Listas Produtos - Página em que estamos atualmente, a página da lista dos produtos do Banco.
3 - Adicionar Aspirador - Insere o Produto Aspirador no Banco.
4 - Adicionar Liquidificador - Insere o Produto Liquidificador no Banco.


Voltando para os produtos do site principal, quando clicamos na imagem de um produto nós somos direcionados a uma página com as informações dele, botão quero comprar abaixo da página temos o Link Conversor de 
Moedas, que é um sistema onde podemos converter o valor real para Dolar, Euro ou Libra caso a moeda que você tenha posse seja diferente do real.

Quando clicamos no botão Quero Comprar, somos levados a uma tela de compra e no fim precisamos inserir o nosso Nome + CEP antes de confirmar a compra.

Depois de confirmado, pronto, compra feita com sucesso.

Login e Cadastro são funcionalidades que tentei implementar.

Não faz muita diferença para o resto do site mas eram funções que queria colocar.

No Cadastro nós temos: E-Mail, Nome, Sobrenome, Senha e Confirmar Senha
Já no Login temos: E-Mail e Senha.

• Link para vídeo de apresentação do projeto (mostrando o funcionamento de todas funcionalidades)

https://www.youtube.com/watch?v=dWNUknLCBbc

• Documentação (como descrição dos dados no BD)

A Tabela que resolvi criar foi a Tabela Produtos, e cada tabela de produtos terá: id, Nome, preco, created, expires (os dois últimos servem para determinar o tempo que o produto estará disponível). Todas as informações inseridas são de produtos que estão disponíveis no site por enquanto. se no futuro mais produtos novos chegarem, a tabela receberá informações dos novos produtos.



![Banco de Dados](https://user-images.githubusercontent.com/62408199/172451632-9d6b26a4-fe4b-4d10-a9b8-34254ddb9962.PNG)





-- Create a `produtos` table.
CREATE TABLE produtos (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
nome VARCHAR(100) NOT NULL,
preco DECIMAL(9,2) NOT NULL,
created DATETIME NOT NULL,
expires DATETIME NOT NULL
);

-- Add an index on the created column.
CREATE INDEX idx_produtos_created ON produtos(created);

-- Add some dummy records.
INSERT INTO produtos (nome, preco, created, expires) VALUES (
'Fogão', 
799.90, 
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO produtos (nome, preco, created, expires) VALUES (
'Geladeira', 
3199.00, 
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO produtos (nome, preco, created, expires) VALUES (
'Lavadora', 
3399.00, 
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO produtos (nome, preco, created, expires) VALUES (
'Micro-ondas', 
579.00, 
UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

SELECT id, nome, preco, expires FROM produtos;



