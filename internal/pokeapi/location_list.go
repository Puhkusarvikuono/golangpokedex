// internal/pokeapi/location_list.go
package pokeapi

type LocationAreaListResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// If url == nil, fetch the first page.
func (c *Client) ListLocationAreas(url *string) (LocationAreaListResponse, error)
