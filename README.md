# 📦 AtStock API (Core)

API de alta performance desenvolvida em **Golang** para o ecossistema AtStock, um sistema inteligente de gestão de stock focado na rastreabilidade de componentes para assistências técnicas.

Esta aplicação atua como o motor principal do sistema, validando regras de negócio e gerenciando a comunicação com a base de dados na nuvem.

## 🚀 Tecnologias e Arquitetura

O projeto foi construído seguindo os princípios da **Clean Architecture** (Arquitetura Limpa), garantindo o desacoplamento, a facilidade de manutenção e a escalabilidade.

* **Linguagem:** Golang (Go)
* **Base de Dados:** PostgreSQL (Hospedado no Supabase)
* **Drivers & Libs:** `lib/pq` (Driver Postgres) e `godotenv` (Gestão de variáveis de ambiente)

### 📂 Estrutura de Diretórios
* `/cmd/api`: Ponto de entrada da aplicação (`main.go`) e definição das rotas HTTP.
* `/internal/domain`: Entidades centrais do sistema (Ex: `StockItemInput`, `StockItemOutput`).
* `/internal/infra`: Camada de infraestrutura, responsável pela conexão com o banco e queries SQL.

## 🔌 Endpoints Disponíveis

A API expõe as seguintes rotas REST:

* `GET /health`: Rota de verificação de integridade (Health Check). Retorna o status operacional do servidor.
* `POST /entrada`: Regista a entrada de um novo componente no stock.
    * **Payload esperado:** `{"serial_number": "SN-123", "condition": "novo"}`
* `GET /item?serial={serial_number}`: Consulta detalhada de um item no stock via Número de Série. Realiza um `LEFT JOIN` relacional com a tabela de produtos para retornar o nome e o modelo.

## 🛠️ Como Executar o Projeto Localmente

1. Clone este repositório:
   ```bash
   git clone [https://github.com/katdott/atstock-api.git](https://github.com/katdott/atstock-api.git)
2. Crie um ficheiro .env na raiz do projeto com a sua connection string da base de dados:
   ```bash
   DATABASE_URL=postgresql://[usuario]:[senha]@[host]:[porta]/postgres
3. Baixe as dependências e inicie o servidor:
   ```bash
    go mod tidy
    go run ./cmd/api/main.go
  O servidor iniciará na porta 8080 (http://localhost:8080).

  Desenvolvido por Agatha Katherine.
