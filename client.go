package espn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"time"
)

type espnClient struct {
	HTTPClient *http.Client

	baseUrl       string
	baseLeagueUrl string
}

func newPublicClient(gameType GameType, leagueId string, year int) *espnClient {
	baseUrl := EspnBaseUrl + "/" + gameType.String() + "/seasons/" + strconv.Itoa(year)
	var baseLeagueUrl string
	if year < 2018 {
		baseLeagueUrl = EspnBaseUrl + "/" + gameType.String() + "/leagueHistory/" + leagueId + "?seasonId=" + strconv.Itoa(year)
	} else {
		baseLeagueUrl = baseUrl + "/segments/0/leagues/" + leagueId
	}
	return &espnClient{
		baseUrl:       baseUrl,
		baseLeagueUrl: baseLeagueUrl,

		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func newPrivateClient(gameType GameType, leagueId string, year int, espnS2 string, swid string) *espnClient {
	jar, _ := cookiejar.New(nil)
	url, _ := url.Parse(EspnBaseUrl)
	cookies := make([]*http.Cookie, 2)
	cookies[0] = &http.Cookie{
		Name:  "espn_s2",
		Value: espnS2,
	}
	cookies[1] = &http.Cookie{
		Name:  "SWID",
		Value: swid,
	}
	jar.SetCookies(url, cookies)

	client := newPublicClient(gameType, leagueId, year)
	client.HTTPClient.Jar = jar
	return client
}

func (c *espnClient) sendRequest(req *http.Request, v interface{}) error {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}

func (c *espnClient) getLeagueInternal(views []string, filter string, v interface{}) error {
	req, err := http.NewRequest("GET", c.baseLeagueUrl, nil)
	if err != nil {
		fmt.Printf("error in espn request: %v", err)
		return err
	}

	if filter != "" {
		req.Header.Add("x-fantasy-filter", filter)
	}

	if len(views) > 0 {
		q := req.URL.Query()
		for _, v := range views {
			q.Add("view", v)
		}
		req.URL.RawQuery = q.Encode()
	}
	err = c.sendRequest(req, &v)
	if err != nil {
		fmt.Printf("error in espn request: %v", err)
		return err
	}
	return nil
}
