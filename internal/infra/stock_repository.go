package infra

import (
	"database/sql"
	"fmt"

	"github.com/agatha-katherine/atstock-api/internal/domain"
)

// StockRepository gerencia a comunicação da tabela stock_items com o banco
type StockRepository struct {
	DB *sql.DB
}

// InserirItem adiciona um novo item com serial único no Supabase
func (r *StockRepository) InserirItem(item domain.StockItemInput) error {
	// A query SQL. O product_id ficará nulo por enquanto para facilitar o seu teste inicial.
	query := `
		INSERT INTO stock_items (serial_number, condition, status) 
		VALUES ($1, $2, 'em_estoque')
	`

	_, err := r.DB.Exec(query, item.SerialNumber, item.Condition)
	if err != nil {
		return fmt.Errorf("erro ao inserir no banco: %w", err)
	}

	return nil
}

// BuscarPorSerial faz a consulta no banco unindo a peça ao seu modelo
func (r *StockRepository) BuscarPorSerial(serial string) (*domain.StockItemOutput, error) {
	// Usamos LEFT JOIN e COALESCE para evitar erros caso a peça ainda não tenha um produto vinculado
	query := `
		SELECT 
			si.serial_number, 
			si.condition, 
			si.status, 
			COALESCE(p.name, 'Produto não vinculado') AS product_name, 
			COALESCE(p.category, 'Sem categoria') AS product_model
		FROM stock_items si
		LEFT JOIN products p ON si.product_id = p.id
		WHERE si.serial_number = $1;
	`

	var item domain.StockItemOutput

	// Executa a query e mapeia o resultado para a struct
	err := r.DB.QueryRow(query, serial).Scan(
		&item.SerialNumber,
		&item.Condition,
		&item.Status,
		&item.ProductName,
		&item.ProductModel,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// Retorna nulo se o serial não existir no banco
			return nil, nil
		}
		return nil, fmt.Errorf("erro ao buscar item: %w", err)
	}

	return &item, nil
}
