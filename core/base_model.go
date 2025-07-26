package core

type (
	BaseObject struct {
		ID string `json:"id"`

		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		DeletedAt string `json:"deleted_at,omitempty"`

		Name string   `json:"name"`
		Tags []string `json:"tags,omitempty"`
	}
)
