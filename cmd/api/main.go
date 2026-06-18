package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/agatha-katherine/atstock-api/internal/domain"
	"github.com/agatha-katherine/atstock-api/internal/infra"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Carrega as variáveis de ambiente (.env)
	godotenv.Load()

	// 2. Conecta ao banco de dados Supabase
	db, err := infra.ConnectDB()
	if err != nil {
		log.Fatalf("Erro Crítico: Falha ao conectar no Supabase: %v", err)
	}
	defer db.Close()
	fmt.Println("✅ Conexão com o Supabase estabelecida!")

	// Instancia o repositório
	stockRepo := &infra.StockRepository{DB: db}

	// ==========================================
	// ROTAS DA API
	// ==========================================

	// Rota 1: Health Check (Para ver se a API está online)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "AtStock API operando normalmente.")
	})

	// Rota 2: Entrada de Estoque (POST /entrada)
	http.HandleFunc("/entrada", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// Trata a requisição de pré-teste (Preflight) que o navegador faz automaticamente
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido. Use POST.", http.StatusMethodNotAllowed)
			return
		}

		var input domain.StockItemInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		if input.SerialNumber == "" {
			http.Error(w, "O número de série é obrigatório", http.StatusBadRequest)
			return
		}

		if err := stockRepo.InserirItem(input); err != nil {
			log.Printf("Erro ao salvar item: %v\n", err)
			http.Error(w, "Erro interno ao salvar o item.", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"status":   "sucesso",
			"mensagem": fmt.Sprintf("Item %s cadastrado com sucesso!", input.SerialNumber),
		})
	})

	// Rota 3: Consulta de Item (GET /item)
	http.HandleFunc("/item", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		// Trata a requisição de pré-teste (Preflight) que o navegador faz automaticamente
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido. Use GET.", http.StatusMethodNotAllowed)
			return
		}

		serial := r.URL.Query().Get("serial")
		if serial == "" {
			http.Error(w, "O parâmetro 'serial' é obrigatório na URL", http.StatusBadRequest)
			return
		}

		item, err := stockRepo.BuscarPorSerial(serial)
		if err != nil {
			log.Printf("Erro ao buscar item: %v\n", err)
			http.Error(w, "Erro interno ao consultar o banco", http.StatusInternalServerError)
			return
		}

		if item == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"erro": "Item não encontrado no estoque",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(item)
	})

	// ==========================================
	// INICIA O SERVIDOR
	// ==========================================
	port := ":8080"
	fmt.Printf("🚀 Servidor AtStock iniciado na porta %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %v", err)
	}
}
