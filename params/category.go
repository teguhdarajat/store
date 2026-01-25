package params

type Category struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Categories struct {
	Categories []Category `json:"categories"`
}
