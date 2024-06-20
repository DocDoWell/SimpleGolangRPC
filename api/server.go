package api

import (
	"movies/gapi"

	"google.golang.org/grpc"
)

type grpcServer struct {
	ratedFilmCollection RatedFilmCollection
	gapi.UnimplementedRatedFilmServiceServer
}

func NewRPCServer() *grpc.Server {
	
	ratedFilmCollection := RatedFilmCollection{collection: make(map[string]Film)} 
	ratedFilmCollection.instantiatecollection()
	
	srv := grpcServer{
		ratedFilmCollection: ratedFilmCollection,
	}

	gsrv := grpc.NewServer()
	gapi.RegisterRatedFilmServiceServer(gsrv, &srv)
	return gsrv
}

