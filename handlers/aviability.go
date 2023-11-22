package handlers

import (
	"context"
	"errors"
	protos "github.com/MihajloJankovic/Aviability-Service/protos/main"
	"log"
)

type myAviabilityServer struct {
	protos.UnimplementedAccommodationAviabilityServer
	logger *log.Logger
	// NoSQL: injecting product repository
	repo *AviabilityRepo
}

func NewServer(l *log.Logger, r *AviabilityRepo) *myAviabilityServer {
	return &myAviabilityServer{*new(protos.UnimplementedAccommodationAviabilityServer), l, r}
}
func (s myAviabilityServer) GetAccommodationCheck(xtx context.Context, in *protos.CheckRequest) (*protos.CheckSet, error) {
	out, err := s.repo.GetAccommodationCheck(xtx, in)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return out, nil
}
func (s myAviabilityServer) GetAllforAccomendation(ctx context.Context, in *protos.GetAllRequest) (*protos.DummyLista3, error) {
	out, err := s.repo.GetAllforAccomendation(ctx, in)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	ss := new(protos.DummyLista3)
	ss.Dummy = out
	return ss, nil
}
func (s myAviabilityServer) DeleteByUser(ctx context.Context, in *protos.DeleteRequestb) (*protos.Emptyb, error) {
	out, err := s.repo.DeleteByUser(ctx, in)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return out, nil
}
func (s myAviabilityServer) SetAccommodationAviability(xtx context.Context, in *protos.CheckSet) (*protos.Emptyb, error) {

	if in.GetPricePerPerson() != 0 && in.GetNumberOfPeople() != 0 {
		in.PriceHole = 0
		bb, err := s.repo.SetAccommodationAviability(xtx, in)
		if err != nil {
			s.logger.Println(err)
			return nil, err
		}
		return bb, nil
	} else {
		if in.GetPriceHole() != 0 {
			in.PricePerPerson = 0
			in.NumberOfPeople = 0
			bb, err := s.repo.SetAccommodationAviability(xtx, in)
			if err != nil {
				s.logger.Println(err)
				return nil, err
			}
			return bb, nil
		} else {
			err := errors.New("inalid data not posible")
			return nil, err
		}
	}

}
