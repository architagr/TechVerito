package persistance

import (
	"movie_booking/enums"
	"movie_booking/errors"
	"movie_booking/models"
)

type IMoviePersistance interface {
	GetAllMovies() []models.MovieModel
	GetMovie(id models.MovieIdModel) (models.MovieModel, error)
	GetAllMoviesByLanguage(lang enums.LanguageEnum) ([]models.MovieModel, error)
}

type MoviePersistance struct {
	movies []models.MovieModel
}

func (repo *MoviePersistance) GetAllMovies() []models.MovieModel {
	return repo.movies
}
func (repo *MoviePersistance) GetMovie(id models.MovieIdModel) (models.MovieModel, error) {

	for _, movie := range repo.movies {
		if movie.Id == id {
			return movie, nil
		}
	}
	return models.MovieModel{}, &errors.MovieNotFoundError{
		Id: id,
	}
}
func (repo *MoviePersistance) GetAllMoviesByLanguage(lang enums.LanguageEnum) (movies []models.MovieModel, err error) {
	movies = make([]models.MovieModel, 0, len(repo.movies))
	for _, movie := range repo.movies {
		if Contains(movie.Languages, lang) {
			movies = append(movies, movie)
		}
	}
	if len(movies) == 0 {
		err = &errors.MovieInLanguageNotFoundError{
			Lang: lang,
		}
	}
	return
}

func Contains(list []enums.LanguageEnum, val enums.LanguageEnum) (ok bool) {
	ok = false
	for _, v := range list {
		if v == val {
			ok = true
			break
		}
	}
	return
}

func InitMoviePersistance() IMoviePersistance {
	return &MoviePersistance{
		movies: []models.MovieModel{
			{
				Id:   models.MovieIdModel(1),
				Name: "Jawan",
				Languages: []enums.LanguageEnum{
					enums.LANGUAGE_ENGLISH,
					enums.LANGUAGE_HINDI,
				},
			},
		},
	}
}
