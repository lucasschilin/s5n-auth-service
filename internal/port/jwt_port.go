package port

type JWT interface {
	GenerateToken(claims map[string]interface{}) (string, error)
}
