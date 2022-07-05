package espn

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestBasePaths(t *testing.T) {

	var tests = []struct {
		year           int
		wantLeaguePath string
	}{
		{2022, "/ffl/seasons/2022/segments/0/leagues/12345"},
		{2017, "/ffl/leagueHistory/12345?seasonId=2017"},
	}

	expectedEspnUrl := "https://fantasy.espn.com/apis/v3/games"

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.year), func(t *testing.T) {
			expectedBasePath := fmt.Sprintf("/ffl/seasons/%d", tt.year)
			client := newPublicClient(GameTypeNfl, "12345", tt.year)

			if client.espnUrl != expectedEspnUrl {
				t.Errorf("expected ESPN URL to be %s, got %s", expectedEspnUrl, client.espnUrl)
			}

			if client.basePath != expectedBasePath {
				t.Errorf("expected base path to be %s, got %s", expectedBasePath, client.basePath)
			}

			if client.baseLeaguePath != tt.wantLeaguePath {
				t.Errorf("expected league path to be %s, got %s", tt.wantLeaguePath, client.baseLeaguePath)
			}
		})
	}
}

func TestBaseRequest(t *testing.T) {
	var tests = []struct {
		name   string
		espnS2 string
		swid   string
		filter string
	}{
		{"publicNoFilter", "", "", ""},
		{"publicFilter", "", "", "{\"foo\": \"bar\"}"},
		{"privateNoFilter", "private_espn_s2", "private_swid", ""},
		{"privateFilter", "private_espn_s2", "private_swid", "{\"foo\": \"bar\"}"},
	}

	expectedPath := "/ffl/seasons/2022/segments/0/leagues/12345"
	expectedQueryParams := map[string][]string{
		"view": {"mView1", "mView2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("starting request check")
				if r.URL.Path != expectedPath {
					t.Errorf("expected request to %s, got %s", expectedPath, r.URL.Path)
				}

				if fmt.Sprintf("%s", r.URL.Query()) != fmt.Sprintf("%s", expectedQueryParams) {
					t.Errorf("expected query params to be %s, got %s", expectedQueryParams, r.URL.Query())
				}

				if h := r.Header[http.CanonicalHeaderKey("x-fantasy-filter")]; h != nil {
					t.Errorf("expected header x-fantasy-filter to be nil, got %s", h)
				}

				expectedCookies := map[string]string{
					"espn_s2": tt.espnS2,
					"SWID":    tt.swid,
				}

				for _, c := range r.Cookies() {
					fmt.Printf("%s: %s\n", c.Name, c.Value)
				}

				for k, v := range expectedCookies {
					cookieVal, err := r.Cookie(k)

					if v == "" && err == nil {
						t.Errorf("expected cookie %s to be nil, got %s", k, cookieVal)
					} else if v != "" && err != nil {
						t.Errorf("expected cookie %s to be %s, got error %s", k, v, err)
					} else if v != "" && v != cookieVal.Value {
						t.Errorf("expected cookie %s to be %s, got %s", k, v, cookieVal.Value)
					}
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{}`))
			}))
			defer server.Close()

			var client *espnClient
			if tt.espnS2 == "" && tt.swid == "" {
				client = newPublicClient(GameTypeNfl, "12345", 2022)
			} else {
				client = newPrivateClient(GameTypeNfl, "12345", 2022, tt.espnS2, tt.swid)
			}

			// Point client at the test server and copy over any cookies set for the ESPN URL
			client.espnUrl = server.URL
			serverUrl, _ := url.Parse(server.URL)
			espnUrl, _ := url.Parse(EspnBaseUrl)

			if client.HTTPClient.Jar != nil {
				cookies := make([]*http.Cookie, 0)
				for _, c := range client.HTTPClient.Jar.Cookies(espnUrl) {
					cookies = append(cookies, &http.Cookie{
						Name:  c.Name,
						Value: c.Value,
					})
				}
				client.HTTPClient.Jar.SetCookies(serverUrl, cookies)
			}

			res := LeagueInfoResponseJson{}
			err := client.getLeagueInternal([]string{"mView1", "mView2"}, "", &res)

			if err != nil {
				t.Errorf("Error in getLeagueInternal: %s", err)
			}
		})
	}
}
