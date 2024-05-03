package types

type GetFruitByNameRequest struct {
	Name string `json:"name"`
}
type AddFruitRequest struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}
