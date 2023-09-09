package article

import (
	"nossobr/domain/people"
	"time"
)

// Article is representing the Article data struct
type Article struct {
	ID          *string           `json:"id,omitempty"`
	CityID      *string           `json:"city_id,omitempty"`
	Content     *[]*Section       `json:"content,omitempty"`
	Children    *[]*people.People `json:"children,omitempty"`
	BorderTowns *[][3]*string     `json:"border_towns,omitempty"`
	CreatedAt   *time.Time        `json:"created_at,omitempty"`
	UpdatedAt   *time.Time        `json:"updated_at,omitempty"`
}

// Section ...
type Section struct {
	Type  *string `json:"type,omitempty"`
	Title *string `json:"title,omitempty"`
	Text  *string `json:"text,omitempty"`
}
