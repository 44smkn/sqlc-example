package param

type PostChairParam struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Thumbnail   string   `json:"thumbnail"`
	Height      int      `json:"height"`
	Width       int      `json:"width"`
	Depth       int      `json:"depth"`
	Color       string   `json:"color"`
	Features    []string `json:"feature"`
	Kind        string   `json:"kind"`
	Popularity  int      `json:"popularity"`
	Stock       int      `json:"stock"`
}
