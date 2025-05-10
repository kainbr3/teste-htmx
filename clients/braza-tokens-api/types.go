package brazatokensapi

import "time"

type OperationTypesResponse struct {
	ID        string    `json:"id"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

type OperationDomainsResponse struct {
	ID        string    `json:"id"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

type BlockchainsResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Abbr      string    `json:"abbr"`
	MainToken string    `json:"main_token"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TokensResponse struct {
	ID           string              `json:"id"`
	IsActive     bool                `json:"is_active"`
	CreatedAt    time.Time           `json:"created_at"`
	UpdatedAt    time.Time           `json:"updated_at"`
	Name         string              `json:"name"`
	Abbr         string              `json:"abbr"`
	Contract     string              `json:"contract"`
	Address      string              `json:"address"`
	Precision    int                 `json:"precision"`
	Type         string              `json:"type"`
	BlockchainID string              `json:"blockchain_id"`
	Blockchain   BlockchainsResponse `json:"blockchain"`
}

type PaginatedOperationsResponse struct {
	TotalCount   int                 `json:"total_count"`
	TotalPages   int                 `json:"total_pages"`
	CurrentPage  int                 `json:"current_page"`
	NextPage     int                 `json:"next_page"`
	PreviousPage int                 `json:"previous_page"`
	Data         []OperationResponse `json:"data"`
}

type OperationResponse struct {
	ID               string    `json:"id"`
	Type             string    `json:"type"`
	Domain           string    `json:"domain"`
	Amount           string    `json:"amount"`
	Operator         string    `json:"operator"`
	Token            string    `json:"token"`
	FireblocksStatus string    `json:"fireblocks_status"`
	BlockchainStatus string    `json:"blockchain_status"`
	FireblocksID     string    `json:"fireblocks_id"`
	TransactionHash  string    `json:"transaction_hash"`
	TransactionLink  string    `json:"transaction_link"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type TransactionAsset struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	TokenID   string          `json:"token_id"`
	Token     *TokensResponse `json:"token"`
	IsActive  bool            `json:"is_active"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type Transaction struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	Domain          string    `json:"domain"`
	Amount          string    `json:"amount"`
	Status          string    `json:"status"`
	ExternalId      string    `json:"external_id"`
	FireblocksId    string    `json:"fireblocks_id"`
	TransactionHash string    `json:"transaction_hash"`
	TransactionLink string    `json:"transaction_link"`
	CallbackURL     string    `json:"callback_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
