package domain

import "time"

// Article is representing the Article data struct
type Article struct {
	ID          *string       `json:"id,omitempty"`
	CityID      *string       `json:"city_id,omitempty"`
	Content     *[]*Section   `json:"content,omitempty"`
	Children    *[]*Children  `json:"children,omitempty"`
	BorderTowns *[][3]*string `json:"border_towns,omitempty"`
	CreatedAt   *time.Time    `json:"created_at,omitempty"`
	UpdatedAt   *time.Time    `json:"updated_at,omitempty"`
}

// Section ...
type Section struct {
	Type  *string `json:"type,omitempty"`
	Title *string `json:"title,omitempty"`
	Text  *string `json:"text,omitempty"`
}

// Children models the data of an illustrious son of a city
type Children struct {
	ID          *string       `json:"id,omitempty"`
	URL         *string       `json:"url,omitempty"`
	Name        *string       `json:"name,omitempty"`
	ShortDesc   *string       `json:"short_desc,omitempty"`
	Biography   *string       `json:"biography,omitempty"`
	Professions *[]*string    `json:"professions,omitempty"`
	Parents     *[2]*Children `json:"parents,omitempty"`
	DateBirth   *string       `json:"date_birth,omitempty"`
	DateDeath   *string       `json:"date_death,omitempty"`
}
