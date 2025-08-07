package dto
type Auth struct {
	User string `json:"name" bilding:'required`
	Token string `json:"token" bilding:"required"`
}
// мне было лень делать зависимоть от абстракций тесты накатывать сюда тоже в падлу так что вот так