package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tsotosa/atmm/config"
	"github.com/tsotosa/atmm/gconst"
	"github.com/tsotosa/atmm/model"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func searchTvShow(query string) (response model.TheMovieDbSearchTvResponse, err error) {
	//fmt.Printf("\nsearching for tv show with query:%#v\n", query)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, config.Conf.TheMovieDbBaseApiUrlV3+"search/tv", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("api_key", config.Conf.ApiKey)
	q.Add("query", query)
	q.Add("language", "en-US")
	q.Add("page", "1")

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result model.TheMovieDbSearchTvResponse
	if err := json.Unmarshal(responseBody, &result); err != nil {
		fmt.Println("Failed to unmarshall JSON")
		fmt.Println(err)
	}

	return result, nil
}

func getTvShowEpisodeDetails(tvId, seasonNumber, episodeNumber string) (response model.TheMovieDbTvShowEpisodeDetails, err error) {
	//fmt.Printf("\nsearching for tv show episode with query:%#v %#v %#v\n", tvId, seasonNumber, episodeNumber)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/tv/%s/season/%s/episode/%s",
			config.Conf.TheMovieDbBaseApiUrlV3, tvId, seasonNumber, episodeNumber), nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("api_key", config.Conf.ApiKey)
	//q.Add("query", "query")
	q.Add("language", "en-US")
	//q.Add("page", "1")

	// assign encoded query string to http request
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return model.TheMovieDbTvShowEpisodeDetails{}, errors.New(string(rune(resp.StatusCode)))
	}

	var result model.TheMovieDbTvShowEpisodeDetails
	if err := json.Unmarshal(responseBody, &result); err != nil {
		fmt.Println("Failed to unmarshall JSON")
		fmt.Println(err)
	}

	return result, nil
}

func GetGithubReleases(page int) ([]model.GithubRelease, error) {
	var res []model.GithubRelease

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%srepos/%s/%s/releases", gconst.GithubApiBaseUrl, "TsotosA", "atmm"), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", gconst.GithubJsonAcceptHeader)
	if config.Conf.GithubUsername != "" && config.Conf.GithubPersonalToken != "" {
		req.SetBasicAuth(config.Conf.GithubUsername, config.Conf.GithubPersonalToken)
	}
	resp, err := client.Do(req)
	if err != nil {
		zap.S().Warnf("could not complete api call with error: %v", err)
		return res, err
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zap.S().Warnf("could not read response body with error: %v", err)
		return res, err
	}

	if resp.StatusCode == http.StatusNotFound {
		zap.S().Warnf("github relied with: %v", resp.StatusCode)
		return res, errors.New(fmt.Sprintf("github replied with %v", resp.StatusCode))
	}
	if err := json.Unmarshal(responseBody, &res); err != nil {
		zap.S().Warnf("could not unmarshall response body with error: %v", err)
		return res, err
	}

	return res, nil
}

func DownloadUrlToLocation(filename, url, location string) error {

	zap.S().Debugf("downloading update binary from url:%s", url)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		zap.S().Warnf("could not download update binary with err: %v", err)
		return err
	}
	req.Header.Set("Accept", gconst.HttpContentTypeOctetStream)
	if config.Conf.GithubUsername != "" && config.Conf.GithubPersonalToken != "" {
		req.SetBasicAuth(config.Conf.GithubUsername, config.Conf.GithubPersonalToken)
	}
	res, err := client.Do(req)
	if err != nil {
		zap.S().Warnf("could not complete api call with error: %v", err)
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		zap.S().Warnf("received %v intead of the expected 200", res.StatusCode)
		return errors.New("received non 200 response code")
	}
	err = os.MkdirAll(location, 0777)
	if err != nil {
		zap.S().Warnf("could not create the required directories with error: %v", err)
		return err
	}
	file, err := os.Create(fmt.Sprintf("%s%s", location, filename))
	if err != nil {
		zap.S().Warnf("could not create empty file with error: %v", err)
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, res.Body)
	if err != nil {
		zap.S().Warnf("could not copy with error: %v", err)
		return err
	}
	err = file.Chmod(0777)
	if err != nil {
		zap.S().Warnf("could not change permissions to new binary: %v", err)
		return err
	}
	return nil
}

func searchMovie(query string) (response model.TheMovieDbSearchMovieResponse, err error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, config.Conf.TheMovieDbBaseApiUrlV3+"search/movie", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("api_key", config.Conf.ApiKey)
	q.Add("query", query)
	q.Add("language", "en-US")
	q.Add("page", "1")

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result model.TheMovieDbSearchMovieResponse
	if err := json.Unmarshal(responseBody, &result); err != nil {
		fmt.Println("Failed to unmarshall JSON")
		fmt.Println(err)
	}

	return result, nil
}
