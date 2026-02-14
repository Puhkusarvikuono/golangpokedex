// internal/pokeapi/client.go
package pokeapi
import (
	"io"
	"log"
	"net/http"
	"time"
	"encoding/json"
)

type Client struct {
	baseURL			string
	httpClient	*http.Client	
}

func NewClient() *Client {
	myClient := Client{
		baseURL: "https://pokeapi.co/api/v2/location-area/",
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}
	return &myClient
}

type LocationAreaResponse struct {
	Next					*string			`json:"next"`
	Previous			*string			`json:"previous"`
	Results				[]result		`json:"results"`
}

type result struct {
	Name					string			`json:"name"`
}

func (c *Client) FetchLocationResponse(url *string) (LocationAreaResponse, error) {
	locResponse := LocationAreaResponse{}
	target := c.baseURL
	if url != nil {
		target = *url
	}
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return locResponse, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locResponse, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return locResponse, err
	}
	
	err = json.Unmarshal(body, &locResponse)
	if err != nil {
		return locResponse, err
	}

	return locResponse, nil

}
