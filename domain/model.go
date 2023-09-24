package domain

import "time"

// Article is representing the Article data struct
type Article struct {
	ID        *string    `json:"id,omitempty"`
	Content   *string    `json:"content,omitempty"`
	Status    *string    `json:"status,omitempty"`
	Props     []*Props   `json:"props,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Props represents the modeling of a property
type Props struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
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

// TODO: in dev.
// type City struct {
// 	ID          *string       `json:"id,omitempty"`
// 	StateID     *string       `json:"state_id,omitempty"`
// 	Name        *string       `json:"name,omitempty"`
// 	BorderTowns *[][3]*string `json:"border_towns,omitempty"`
// 	Latitude    *float64      `json:"latitude,omitempty"`
// 	Longitude   *float64      `json:"longitude,omitempty"`
// }
