package back

type Character struct {
	Id          int    `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
}
