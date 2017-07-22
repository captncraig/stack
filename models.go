package stack

import "encoding/json"

type Response struct {
	Backoff         int             `json:"backoff"`
	ErrorID         int             `json:"error_id"`
	ErrorMsg        string          `json:"error_message"`
	ErrorName       string          `json:"error_name"`
	HasMore         bool            `json:"has_more"`
	QuotaMax        int             `json:"quota_max"`
	QuotaRemnaining int             `json:"quota_remaining"`
	Items           json.RawMessage `json:"items,omitempty"`
}

type Site struct {
	Aliases               []string            `json:"aliases"`
	ApiSiteParameter      string              `json:"api_site_parameter"`
	Audience              string              `json:"audience"`
	ClosedBetaData        uint64              `json:"closed_beta_date"`
	FaviconURL            string              `json:"favicon_url"`
	HighResolutionIconURL string              `json:"high_resolution_icon_url"`
	IconURL               string              `json:"icon_url"`
	LaunchDate            uint64              `json:"launch_date"`
	LogoURL               string              `json:"logo_url"`
	MarkdownExtensions    []MarkdownExtension `json:"markdown_extensions"`
	Name                  string              `json:"name"`
	OpenBetaDate          uint64              `json:"open_beta_date"`
	RelatedSites          []*RelatedSite      `json:"related_sites"`
	SiteState             SiteState           `json:"site_state"`
	SiteType              SiteType            `json:"site_type"`
	SiteURL               string              `json:"site_url"`
	TwitterAccount        string              `json:"twitter_account"`
}

type RelatedSite struct {
	ApiSiteParameter string `json:"api_site_parameter"`
	Name             string `json:"name"`
	Relation         Relation
	SiteURL          string `json:"site_url"`
}

// enum types

type SiteState string
type SiteType string
type MarkdownExtension string
type Relation string
