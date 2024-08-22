package domain

type LinksIn struct {
	FullLink string `json:"full_link"`
}

type SetLinkResponse struct {
	ShortLink string `json:"short_link"`
}

type GetLinkResponse struct {
	FullLink string `json:"full_link"`
}
