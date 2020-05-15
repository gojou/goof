package service

// Service is just this struct, you know?
type Service struct {
}

type Struct interface {
	GetFencer(id string) (*Fencer, error)
	AddFencer(fencer *Fencer) (*Fencer, error)
	ListFencers() ([]Fencer, error)
}

func GetFencer(id string) (*Fencer, error) {
	return nil, nil
}
func AddFencer(fencer *Fencer) (*Fencer, error) {
	return nil, nil
}
func ListFencers() ([]Fencer, error) {
	return nil, nil
}
