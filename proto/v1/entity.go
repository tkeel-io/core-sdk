package v1

// Create Entity Request.
type CreateEntityRequest struct {
	Id         string                 `json:"id,omitempty"`
	Type       string                 `json:"type,omitempty"`
	From       string                 `json:"from,omitempty"`
	Owner      string                 `json:"owner,omitempty"`
	Source     string                 `json:"source,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// Update Entity Request.
type UpdateEntityRequest struct {
	Id         string                 `json:"id,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Owner      string                 `json:"owner,omitempty"`
	Source     string                 `json:"source,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// Get Entity Request.
type GetEntityRequest struct {
	Id     string `json:"id,omitempty"`
	Type   string `json:"type,omitempty"`
	Owner  string `json:"owner,omitempty"`
	Source string `json:"source,omitempty"`
}

// Delete Entity Request.
type DeleteEntityRequest struct {
	Id     string `json:"id,omitempty"`
	Type   string `json:"type,omitempty"`
	Owner  string `json:"owner,omitempty"`
	Source string `json:"source,omitempty"`
}

type DeleteEntityResponse struct {
	Id     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}

// Update(upsert) Entity Properties Request.
type UpdateEntityPropsRequest struct {
	Id         string                 `json:"id,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Owner      string                 `json:"owner,omitempty"`
	Source     string                 `json:"source,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// Patch Entity Properties Request.
type PatchEntityPropsRequest struct {
	Id         string      `json:"id,omitempty"`
	Type       string      `json:"type,omitempty"`
	Owner      string      `json:"owner,omitempty"`
	Source     string      `json:"source,omitempty"`
	Properties []PatchData `json:"properties,omitempty"`
}

// Get Entity Properties Request.
type GetEntityPropsRequest struct {
	Id           string   `json:"id,omitempty"`
	Type         string   `json:"type,omitempty"`
	Owner        string   `json:"owner,omitempty"`
	Source       string   `json:"source,omitempty"`
	PropertyKeys []string `json:"property_keys,omitempty"`
}

// Remove Entity Properties Request.
type RemoveEntityPropsRequest struct {
	Id           string   `json:"id,omitempty"`
	Type         string   `json:"type,omitempty"`
	Owner        string   `json:"owner,omitempty"`
	Source       string   `json:"source,omitempty"`
	PropertyKeys []string `json:"property_keys,omitempty"`
}

// Update(upsert) Entity Configs Request.
type UpdateEntityConfigsRequest struct {
	Id      string   `json:"id,omitempty"`
	Type    string   `json:"type,omitempty"`
	Owner   string   `json:"owner,omitempty"`
	Source  string   `json:"source,omitempty"`
	Configs []Config `json:"configs,omitempty"`
}

// Patch Entity Configs Request.
type PatchEntityConfigsRequest struct {
	Id      string            `json:"id,omitempty"`
	Type    string            `json:"type,omitempty"`
	Owner   string            `json:"owner,omitempty"`
	Source  string            `json:"source,omitempty"`
	Configs []ConfigPatchData `json:"configs,omitempty"`
}

// Get Entity Configs Request.
type GetEntityConfigsRequest struct {
	Id           string   `json:"id,omitempty"`
	Type         string   `json:"type,omitempty"`
	Owner        string   `json:"owner,omitempty"`
	Source       string   `json:"source,omitempty"`
	PropertyKeys []string `json:"property_keys,omitempty"`
}

// Remove Entity Configs Request.
type RemoveEntityConfigsRequest struct {
	Id           string   `json:"id,omitempty"`
	Type         string   `json:"type,omitempty"`
	Owner        string   `json:"owner,omitempty"`
	Source       string   `json:"source,omitempty"`
	PropertyKeys []string `json:"property_keys,omitempty"`
}

// List Entity Request.
type ListEntityRequest struct {
	Type      string             `json:"type,omitempty"`
	Owner     string             `json:"owner,omitempty"`
	Source    string             `json:"source,omitempty"`
	Query     string             `json:"query,omitempty"`
	Page      *Pager             `json:"page,omitempty"`
	Condition []*SearchCondition `json:"condition,omitempty"`
}

// List Entity Response.
type ListEntityResponse struct {
	Total  int64             `json:"total,omitempty" mapstructure:"total"`
	Page   int64             `json:"page,omitempty" mapstructure:"page"`
	Limit  int64             `json:"limit,omitempty" mapstructure:"limit"`
	Offset int64             `json:"offset,omitempty" mapstructure:"offset"`
	Items  []*EntityResponse `json:"items,omitempty" mapstructure:"items"`
}

// Entity Response.
type EntityResponse struct {
	Id         string                 `json:"id,omitempty" mapstructure:"id"`
	Type       string                 `json:"type,omitempty" mapstructure:"type"`
	Owner      string                 `json:"owner,omitempty" mapstructure:"owner"`
	Source     string                 `json:"source,omitempty" mapstructure:"source"`
	Mappers    []*Mapper              `json:"mappers,omitempty" mapstructure:"mappers"`
	Configs    map[string]Config      `json:"configs,omitempty" mapstructure:"configs"`
	Properties map[string]interface{} `json:"properties,omitempty" mapstructure:"properties"`
}
