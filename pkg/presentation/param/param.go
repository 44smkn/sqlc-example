package param

type PostChairParam struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Thumbnail   string   `json:"thumbnail"`
	Height      int32    `json:"height"`
	Width       int32    `json:"width"`
	Depth       int32    `json:"depth"`
	Color       string   `json:"color"`
	Features    []string `json:"feature"`
	Kind        string   `json:"kind"`
	Popularity  int32    `json:"popularity"`
	Stock       int32    `json:"stock"`
}
