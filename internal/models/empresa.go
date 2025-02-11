package models

type Empresa struct {
	ID                      int     `json:"id"`
	RazonSocial             *string `json:"razon_social"`
	NumeroIdentificador     *string `json:"numero_identificador"`
	TipoNumeroIdentificador *string `json:"tipo_numero_identificador"`
}
