// internal/pokeapi/client.go
package pokeapi
import (
	"io"
	"log"
	"net/http"
	"time"
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

func (c *Client) GetLocationPage(url *string) ([]byte, error) {
	target := c.baseURL
	if url != nil {
		target = *url
	}
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
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
 	return body, nil	
}
