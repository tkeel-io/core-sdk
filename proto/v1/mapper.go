package v1

type Mapper struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	TqlText     string `json:"tql_text,omitempty"`
	Description string `json:"description,omitempty"`
}

// Append Mapper Request.
type AppendMapperRequest struct {
	Type     string  `json:"type,omitempty"`
	Owner    string  `json:"owner,omitempty"`
	Source   string  `json:"source,omitempty"`
	EntityId string  `json:"entity_id,omitempty"`
	Mapper   *Mapper `json:"mapper,omitempty"`
}

// Get Mapper Request.
type GetMapperRequest struct {
	Id       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Owner    string `json:"owner,omitempty"`
	Source   string `json:"source,omitempty"`
	EntityId string `json:"entity_id,omitempty"`
}

// List Mapper Request.
type ListMapperRequest struct {
	Type     string `json:"type,omitempty"`
	Owner    string `json:"owner,omitempty"`
	Source   string `json:"source,omitempty"`
	EntityId string `json:"entity_id,omitempty"`
}

// Remove Mapper Request.
type RemoveMapperRequest struct {
	Id       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Owner    string `json:"owner,omitempty"`
	Source   string `json:"source,omitempty"`
	EntityId string `json:"entity_id,omitempty"`
}

// Append Mapper Response.
type AppendMapperResponse struct {
	Type     string  `json:"type,omitempty"`
	Owner    string  `json:"owner,omitempty"`
	Source   string  `json:"source,omitempty"`
	EntityId string  `json:"entity_id,omitempty"`
	Mapper   *Mapper `json:"mapper,omitempty"`
}

// Remove Mapper Response.
type RemoveMapperResponse struct {
	Id       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Owner    string `json:"owner,omitempty"`
	Source   string `json:"source,omitempty"`
	EntityId string `json:"entity_id,omitempty"`
}

// Get Mapper Response.
type GetMapperResponse struct {
	Type     string  `json:"type,omitempty"`
	Owner    string  `json:"owner,omitempty"`
	Source   string  `json:"source,omitempty"`
	EntityId string  `json:"entity_id,omitempty"`
	Mapper   *Mapper `json:"mapper,omitempty"`
}

// List Mapper Response.
type ListMapperResponse struct {
	Type     string    `json:"type,omitempty"`
	Owner    string    `json:"owner,omitempty"`
	Source   string    `json:"source,omitempty"`
	EntityId string    `json:"entity_id,omitempty"`
	Mappers  []*Mapper `json:"mappers,omitempty"`
}




