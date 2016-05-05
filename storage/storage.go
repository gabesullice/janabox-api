package storage

type Storage interface {
	Insert(string) (string, error)
	Get(string) (string, error)
	Update(string, string) error
	Remove(string) error
}
