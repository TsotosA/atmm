package model

type TheMovieDbSearchTvResponse struct {
	Page         int                `json:"page"`
	Results      []TheMovieDbTvShow `json:"results"`
	TotalResults int                `json:"total_results"`
	TotalPages   int                `json:"total_pages"`
}

type TheMovieDbSearchMovieResponse struct {
	Page         int               `json:"page"`
	Results      []TheMovieDbMovie `json:"results"`
	TotalResults int               `json:"total_results"`
	TotalPages   int               `json:"total_pages"`
}

type TheMovieDbTvShow struct {
	PosterPath       string  `json:"poster_path"`
	Popularity       float32 `json:"popularity"`
	Id               int     `json:"id"`
	BackdropPath     string  `json:"backdrop_path"`
	VoteAverage      float32 `json:"vote_average"`
	Overview         string  `json:"overview"`
	FirstAirDate     string  `json:"first_air_date"`
	OriginalCountry  string  `json:"original_country"`
	GenreIds         []int   `json:"genre_ids"`
	OriginalLanguage string  `json:"original_language"`
	VoteCount        int     `json:"vote_count"`
	Name             string  `json:"name"`
	OriginalName     string  `json:"original_name"`
}

type TheMovieDbMovie struct {
	PosterPath       string  `json:"poster_path"`
	Adult            bool    `json:"adult"`
	Overview         string  `json:"overview"`
	ReleaseDate      string  `json:"release_date"`
	GenreIds         []int   `json:"genre_ids"`
	Id               int     `json:"id"`
	OriginalTitle    string  `json:"original_title"`
	OriginalLanguage string  `json:"original_language"`
	Title            string  `json:"Title"`
	BackdropPath     string  `json:"backdrop_path"`
	Popularity       float32 `json:"popularity"`
	VoteCount        int     `json:"vote_count"`
	Video            bool    `json:"video"`
	VoteAverage      float32 `json:"vote_average"`
}

type TheMovieDbTvShowEpisodeDetails struct {
	AirDate       string `json:"air_date"`
	EpisodeNumber int    `json:"episode_number"`
	Name          string `json:"name"`
	Overview      string `json:"overview"`
	Id            int    `json:"id"`
	SeasonNumber  int    `json:"season_number"`
}
