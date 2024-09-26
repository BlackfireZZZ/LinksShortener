package domain

type LinksOut struct {
	FullLink  string `json:"full_link"`
	ShortLink string `json:"short_link"`
}

type LinksIn struct {
	FullLink string `json:"full_link"`
}

type SetLinkResponse struct {
	ShortLink string `json:"short_link"`
}

type GetLinkResponse struct {
	FullLink string `json:"full_link"`
}
