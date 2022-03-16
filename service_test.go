package main

//func TestSearchTvShow(t *testing.T) {
//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		res := TheMovieDbSearchTvResponse{
//			Page:       0,
//			TotalPages: 1,
//			Results: []TheMovieDbTvShow{{
//				PosterPath:       "",
//				Popularity:       0,
//				Id:               1234,
//				BackdropPath:     "",
//				VoteAverage:      0,
//				Overview:         "this is an overview",
//				FirstAirDate:     "",
//				OriginalCountry:  "",
//				GenreIds:         nil,
//				OriginalLanguage: "",
//				VoteCount:        0,
//				Name:             "This is the name",
//				OriginalName:     "",
//			}},
//		}
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//		_ = json.NewEncoder(w).Encode(res)
//	}))
//	t.Cleanup(func() {
//		ts.Close()
//	})
//	call, _ := http.Get(ts.URL)
//	res, _ := io.ReadAll(call.Body)
//	var parsed TheMovieDbSearchTvResponse
//	_ = json.Unmarshal(res, &parsed)
//	_ = call.Body.Close()
//	t.Logf("res: [%#v]", parsed)
//}
