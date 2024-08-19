package domain

type LinksOut struct {
	FullLink  string `json:"full_link"`
	ShortLink string `json:"short_link"`
}

type LinksIn struct {
	FullLink string `json:"full_link"`
}
