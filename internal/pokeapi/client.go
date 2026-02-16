// internal/pokeapi/client.go
package pokeapi
import (
	"io"
	"log"
	"net/http"
	"time"
	"encoding/json"
	"github.com/Puhkusarvikuono/golangpokedex/internal/pokecache"
)

type Client struct {
	baseURL			string
	httpClient	*http.Client
	cache				pokecache.Cache
}

func NewClient(cacheInterval time.Duration) *Client {
	myClient := Client{
		baseURL: "https://pokeapi.co/api/v2/location-area/",
		httpClient: &http.Client{Timeout: 5 * time.Second},
		cache: pokecache.NewCache(cacheInterval),
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

	val, ok := c.cache.Get(target)
	if ok {
		err := json.Unmarshal(val, &locResponse)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locResponse, nil
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

	c.cache.Add(target, body)
	
	err = json.Unmarshal(body, &locResponse)
	if err != nil {
		return locResponse, err
	}

	return locResponse, nil

}
