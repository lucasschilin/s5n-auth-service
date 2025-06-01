package port

type JWT interface {
	GenerateToken(claims map[string]interface{}) (string, error)
	ValidateToken(token string) (map[string]interface{}, error)
}
