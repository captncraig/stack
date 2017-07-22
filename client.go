package stack

import (
	"context"
	"encoding/json"
	"net/http"

	"fmt"

	"golang.org/x/oauth2"
)

// Client is the main client for all api endpoints. It has children in a heirarchy that fiollow the headings on the documentation page (https://api.stackexchange.com/docs)
type Client interface {
	Network() Network
	Site(string) SingleSite
}

type Network interface {
	Sites() NetworkSites
}

type NetworkSites interface {
	GetAllSites(ctx context.Context) ([]*Site, *Response, error)
}
type SingleSite interface {
}

type client struct {
	hc *http.Client
}

type siteClient struct {
	*client
	site string
}

func (c *client) Site(name string) SingleSite {
	return &siteClient{client: c, site: name}
}

func (c *client) Network() Network {
	return c
}

func (c *client) Sites() NetworkSites {
	return c
}

func NewClient() Client {
	return create(http.DefaultClient)
}

func NewAuthenticatedClient(accessToken string) Client {
	return create(oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})))
}

func create(hc *http.Client) Client {
	return &client{hc: hc}
}

const baseURL = "https://api.stackexchange.com/2.2"

func (c *client) do(req *http.Request, target interface{}) (*Response, error) {
	resp, err := c.hc.Do(req)
	if err != nil {
		return nil, err
	}
	// TODO: possibly check status code. Return smoother errors for go std-lib errors
	defer resp.Body.Close()
	r := &Response{}
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(r); err != nil {
		return nil, err
	}

	if r.ErrorID != 0 {
		return nil, fmt.Errorf("Error %d: %s. %s", r.ErrorID, r.ErrorName, r.ErrorMsg)
	}
	if err = json.Unmarshal([]byte(r.Items), target); err != nil {
		return nil, err
	}
	//clear raw message
	r.Items = nil
	return r, nil
}

func (c *client) GetAllSites(ctx context.Context) ([]*Site, *Response, error) {
	const url = baseURL + "/sites?pagesize=10000"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	sites := []*Site{}
	resp, err := c.do(req, &sites)
	if err != nil {
		return nil, nil, err
	}
	return sites, resp, err
}
