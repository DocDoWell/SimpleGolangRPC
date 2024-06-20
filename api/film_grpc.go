package api

import (
	"context"
	"errors"
	"movies/gapi"
)

func (server * grpcServer) GetRatedFilm(ctx context.Context, req *gapi.GetFilmRequest) (*gapi.GetFilmResponse, error) {
	actor := req.Name

	film, ok := server.ratedFilmCollection.collection[actor]

   if ok {
	   return &gapi.GetFilmResponse{Film:&gapi.Film{Id:film.Id, Title: film.Title, Rating: film.Rating}}, nil
   }

   return &gapi.GetFilmResponse{}, errors.New("a movie starring this actor does not belong to our comprehensive collection - try gain")

}


