package main

type Cacabox struct{}

func (s *Cacabox) Name() string {

	return "Name"
}

func (s *Cacabox) Description() string {

	return "Description"
}

func (s *Cacabox) Version() string {

	return "1.0.0"
}

type Vroum struct{}

func (s *Vroum) Name() string {

	return "Vroum"
}

func (s *Vroum) Description() string {

	return "Vroum Desc"
}

func (s *Vroum) Version() string {

	return "1.0.0"
}
