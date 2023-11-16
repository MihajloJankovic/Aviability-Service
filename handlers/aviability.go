package handlers

import (
	"context"
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
func (s myAviabilityServer) GetAccommodationCheck(xtx context.Context, in *protos.CheckRequest) (*protos.Emptyb, error) {
	return nil, nil
}
func (s myAviabilityServer) GetAllforAccomendation(xtx context.Context, in *protos.Emptyb) (*protos.DummyList, error) {
	return nil, nil
}
func (s myAviabilityServer) SetAccommodationAviability(xtx context.Context, in *protos.CheckSet) (*protos.Emptyb, error) {
	return nil, nil
}
