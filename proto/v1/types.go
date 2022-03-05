package v1

type PatchData struct {
	Path     string      `json:"path"`
	Value    interface{} `json:"value"`
	Operator string      `json:"operator"`
}

type ConfigPatchData struct {
	Path     string `json:"path"`
	Value    Config `json:"value"`
	Operator string `json:"operator"`
}

type Config struct {
	ID                string                 `json:"id" mapstructure:"id"`
	Type              string                 `json:"type" mapstructure:"type"`
	Name              string                 `json:"name" mapstructure:"name"`
	Weight            int                    `json:"weight" mapstructure:"weight"`
	Enabled           bool                   `json:"enabled" mapstructure:"enabled"`
	EnabledSearch     bool                   `json:"enabled_search" mapstructure:"enabled_search"`
	EnabledTimeSeries bool                   `json:"enabled_time_series" mapstructure:"enabled_time_series"`
	Description       string                 `json:"description" mapstructure:"description"`
	Define            map[string]interface{} `json:"define" mapstructure:"define"`
	LastTime          int64                  `json:"last_time" mapstructure:"last_time"`
}

type Pager struct {
	Limit   int64  `json:"limit,omitempty"`
	Offset  int64  `json:"offset,omitempty"`
	Sort    string `json:"sort,omitempty"`
	Reverse bool   `json:"reverse,omitempty"`
}

type SearchCondition struct {
	Field    string      `json:"field,omitempty"`
	Operator string      `json:"operator,omitempty"`
	Value    interface{} `json:"value,omitempty"`
}
