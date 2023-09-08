package people

// People models the data of an illustrious son of a city
type People struct {
	ID          *string     `json:"id,omitempty"`
	URL         *string     `json:"url,omitempty"`
	Name        *string     `json:"name,omitempty"`
	ShortDesc   *string     `json:"short_desc,omitempty"`
	Biography   *string     `json:"biography,omitempty"`
	Professions *[]*string  `json:"professions,omitempty"`
	Parents     *[2]*People `json:"parents,omitempty"`
	DateBirth   *string     `json:"date_birth,omitempty"`
	DateDeath   *string     `json:"date_death,omitempty"`
}
