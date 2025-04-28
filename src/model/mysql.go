package model

type MySqlReqArgs struct {
	Table        string         `json:"table"`
	Where        map[string]any `json:"where"`
	WhereGreater map[string]any `json:"where_greater"`
	WhereLess    map[string]any `json:"where_less"`
	WhereNot     map[string]any `json:"where_not"`
	Data         map[string]any `json:"data"`
	Limit        string         `json:"limit"`
}

type Book struct {
	Title       string `json:"titles"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	Year        int    `json:"year"`
	Description string `json:"description"`
}
