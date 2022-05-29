# TrabalhoRicardoCC5M

• Descrição do projeto

• Link para vídeo de apresentação do projeto (mostrando o funcionamento de todas funcionalidades)


• Documentação (como descrição dos dados no BD)

A Tabela que resolvi criar foi a Tabela Produtos, e cada tabela de produtos terá: id, Nome, preco, created, expires (os dois últimos servem para determinar o tempo que o produto estará disponível). Todas as informações inseridas são de produtos que estão disponíveis no site por enquanto. se no futuro mais produtos novos chegarem, a tabela receberá informações dos novos produtos.

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


