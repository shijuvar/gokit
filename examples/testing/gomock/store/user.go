package store

//go:generate mockgen -destination=../mocks/mock_store.go -package=mocks github.com/shijuvar/gokit/examples/testing/gomock/store UserStore

type UserStore interface {
	AddUser(string) error
}
