package service

import (
	"context"
	"gosharp/internal/dao"
)

// Service service.
type Service struct {
	dao *dao.Dao
}

// New new a service and return.
func New() (s *Service) {
	s = &Service{
		dao: dao.New(),
	}
	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}
