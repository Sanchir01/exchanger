package exchanger

type Services struct {
	repository Repositories
}
type Repositories interface {
}

func NewService(repo Repositories) *Services {
	return &Services{repository: repo}
}
