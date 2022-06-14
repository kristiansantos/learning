package repositories

type IPlacementRepository interface {
	Index() error
	Show() error
	Delete() error
	Update() error
}
