package api

type RatedFilmCollection struct {
	collection    map[string]Film // increment book id
}

func (r *RatedFilmCollection) instantiatecollection() {
	r.collection["Gary Oldman"] = Film{Id:"1",Title:"Leon",Rating:100}
	r.collection["Anthony Hopkins"] = Film{Id:"2",Title:"The Silence Of The Lambs",Rating:90}
	r.collection["Robert De Niro"] = Film{Id:"3",Title:"Goodfellas",Rating:98}
	r.collection["Al Pacino"] = Film{Id:"4",Title:"The Godfather",Rating:100}
	r.collection["Toni Collette"] = Film{Id:"5",Title:"Knives Out",Rating:85}
	r.collection["Jamie Lee Curtis"] = Film{Id:"6",Title:"Halloween",Rating:91}
}


