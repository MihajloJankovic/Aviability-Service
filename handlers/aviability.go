package handlers

import (
	"context"
	"errors"
	protos "github.com/MihajloJankovic/Aviability-Service/protos/main"
	"log"
	"strings"
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

// isValidUID checks if the UID is valid.
func isValidUID(uid string) bool {
	return len(uid) > 0
}

// isValidDateRange checks if the date range is valid.
func isValidDateRange(from, to string) bool {
	// Add your date range validation logic here
	return true
}

// isValidPrice checks if the price is valid.
func isValidPrice(price int32) bool {
	// Add your price validation logic here
	return price >= 0
}

// trimString trims leading and trailing whitespaces from a string.
func trimString(str string) string {
	return strings.TrimSpace(str)
}
func (s myAviabilityServer) GetAccommodationCheck(xtx context.Context, in *protos.CheckRequest) (*protos.CheckSet, error) {
	// Validate UID
	if !isValidUID(in.GetId()) {
		return nil, errors.New("Invalid UID")
	}

	// Validate date range
	if !isValidDateRange(in.GetFrom(), in.GetTo()) {
		return nil, errors.New("Invalid date range")
	}

	// Trim input data
	in.Id = trimString(in.GetId())
	in.From = trimString(in.GetFrom())
	in.To = trimString(in.GetTo())

	// Perform the actual logic
	out, err := s.repo.GetAccommodationCheck(xtx, in)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return out, nil
}
func (s myAviabilityServer) GetAllforAccomendation(ctx context.Context, in *protos.GetAllRequest) (*protos.DummyLista3, error) {
	// Validate UID
	if !isValidUID(in.GetId()) {
		return nil, errors.New("Invalid UID")
	}

	// Trim input data
	in.Id = trimString(in.GetId())

	// Perform the actual logic
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
	// Validate UID
	if !isValidUID(in.GetUid()) {
		return nil, errors.New("Invalid UID")
	}

	// Trim input data
	in.Uid = trimString(in.GetUid())

	// Perform the actual logic
	out, err := s.repo.DeleteByUser(ctx, in)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return out, nil
}
func (s myAviabilityServer) SetAccommodationAviability(xtx context.Context, in *protos.CheckSet) (*protos.Emptyb, error) {
	// Validate price and number of people
	if in.GetPricePerPerson() != 0 && in.GetNumberOfPeople() != 0 {
		// Validate price
		if !isValidPrice(in.GetPricePerPerson()) {
			return nil, errors.New("Invalid price per person")
		}

		// Trim input data
		in.Uid = trimString(in.GetUid())
		in.From = trimString(in.GetFrom())
		in.To = trimString(in.GetTo())

		bb, err := s.repo.SetAccommodationAviability(xtx, in)
		if err != nil {
			s.logger.Println(err)
			return nil, err
		}
		return bb, nil
	} else {
		if in.GetPriceHole() != 0 {
			// Validate price
			if !isValidPrice(in.GetPriceHole()) {
				return nil, errors.New("Invalid price for the entire hole")
			}

			// Trim input data
			in.Uid = trimString(in.GetUid())
			in.From = trimString(in.GetFrom())
			in.To = trimString(in.GetTo())

			bb, err := s.repo.SetAccommodationAviability(xtx, in)
			if err != nil {
				s.logger.Println(err)
				return nil, err
			}
			return bb, nil
		} else {
			err := errors.New("Invalid data, not possible")
			return nil, err
		}
	}
}
func (s myAviabilityServer) GetallbyIDandPrice(ctx context.Context, in *protos.PriceAndIdRequest) (*protos.DummyLista3, error) {
	// Validacija UID
	if !isValidUID(in.GetId()) {
		return nil, errors.New("Invalid UID")
	}

	// Validacija cene
	if !isValidPrice(int32(in.GetMinPrice())) || !isValidPrice(int32(in.GetMaxPrice())) {
		return nil, errors.New("Invalid price range")
	}

	// Dohvati podatke iz baze pomoću repo funkcije
	out, err := s.repo.GetallbyIDandPrice(ctx, in)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	// Ako nema rezultata, možete vratiti odgovarajuću grešku ili praznu listu
	if out == nil || len(out.Dummy) == 0 {
		return &protos.DummyLista3{
			Dummy: []*protos.CheckSet{},
		}, nil
	}

	return out, nil
}
