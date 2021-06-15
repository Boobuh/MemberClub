package storage

type Storage struct {
	Data Members
}

func NewStorage() *Storage {
	data := Members{}
	return &Storage{Data: data}
}

func (s *Storage) GetAll() Members {
	return s.Data
}
func (s *Storage) Save(member Member) error {
	err := member.Validate()
	if err != nil {
		return err
	}
	member.SetRegistrationDate()
	s.Data = append(s.Data, member)
	return nil
}
