// shippy-user-service/token_service.go

package main

type Authable interface {
	Decode(token string) (interface{}, error)
	Encode(data interface{}) (string, error)
}

type TokenService struct {
	repo repository
}

// Decode - Decode the token
func (srv *TokenService) Decode(token string) (interface{}, error) {
	return "", nil
}

// Encode - Encode the token
func (srv *TokenService) Encode(data interface{}) (string, error) {
	return "", nil
}
