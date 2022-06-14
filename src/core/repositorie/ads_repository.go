package repositories

type IAdsRepository interface {
	Index() error
	Show() error
	Delete() error
	Update() error
}
