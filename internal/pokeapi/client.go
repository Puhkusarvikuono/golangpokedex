// internal/pokeapi/client.go
package pokeapi
import (
	"io"
	"fmt"
	"net/http"
	"strings"
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
	Results				[]Result		`json:"results"`
}

type Result struct {
	Name								string												`json:"name"`
}

type ExploreAreaResponse struct {
	PokemonEncounters			[]PokemonEncounter					`json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon			Pokemon			`json:"pokemon"`
}

type Pokemon struct {
	Name							string					`json:"name"`
	URL								string					`json:"url"`
	BaseExperience		int							`json:"base_experience"`
	Height						int							`json:"height"`
	Weight						int							`json:"weight"`

	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`

	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
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

	body, err := c.HTTPGetResponse(target)
	if err != nil {
		return locResponse, err
	}
	
	c.cache.Add(target, body)
	
	err = json.Unmarshal(body, &locResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locResponse, nil

}

func (c *Client) ExploreLocationResponse(areaName string) (ExploreAreaResponse, error) {
	locResponse := ExploreAreaResponse{}

	target := c.baseURL + areaName + "/"

	val, ok := c.cache.Get(target)
	if ok {
		err := json.Unmarshal(val, &locResponse)
		if err != nil {
			return locResponse, err
		}
		return locResponse, nil
	}

	body, err := c.HTTPGetResponse(target)
	if err != nil {
		return locResponse, err
	}

	c.cache.Add(target, body)
	
	err = json.Unmarshal(body, &locResponse)
	if err != nil {
		return ExploreAreaResponse{}, err
	}

	return locResponse, nil

}

func (c *Client) FetchPokemonData(pokemonName string) (Pokemon, error) {
	locResponse := Pokemon{}
	target := strings.TrimSuffix(c.baseURL, "/location-area/") + "/pokemon/" + pokemonName + "/"
	
	val, ok := c.cache.Get(target)
	if ok {
		err := json.Unmarshal(val, &locResponse)
		if err != nil {
			return Pokemon{}, err
		}
		return locResponse, nil
	}

	body, err := c.HTTPGetResponse(target)
	if err != nil {
		return locResponse, err
	}
	
	c.cache.Add(target, body)
	
	err = json.Unmarshal(body, &locResponse)
	if err != nil {
		return Pokemon{}, err
	}

	return locResponse, nil

}


func (c *Client) HTTPGetResponse(target string) ([]byte, error) {
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		err = fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}


