package domain

import "time"

type Base struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type Namespace struct {
	Base
	Environment string `json:"environment"`
	Name        string `json:"namespace"`
}

type Variable struct {
	Base
	Environment string     `json:"environment"`
	NamespaceID string     `json:"namespace_id"`
	NameSpace   *Namespace `json:"namespace"`
	Key         string     `json:"key"`
	Value       string     `json:"value"`
}
