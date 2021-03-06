package fedruntime

// FederatedService is the service object that the
// generated.go file will return for the _service
// query
type FederatedService struct {
	SDL string `json:"sdl"`
}

// Everything with a @key implements this
type FederatedEntity interface {
	IsFederatedEntity()
}
