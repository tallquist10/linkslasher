package links

type Link struct {
	Original string `json:"original" binding:"required"`
	Hash     string `json:"hash"`
}

type LinksApiRequest struct {
	Method string
	Params map[string]string
}
