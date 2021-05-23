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

type SearchChairParam struct {
	MaxPrice  *int32   `json:"max_price"`
	MinPrice  *int32   `json:"min_price"`
	MaxHeight *int32   `json:"max_height"`
	MinHeight *int32   `json:"min_height"`
	MaxWidth  *int32   `json:"max_width"`
	MinWidth  *int32   `json:"min_width"`
	MaxDepth  *int32   `json:"max_depth"`
	MinDepth  *int32   `json:"min_depth"`
	Kind      *string  `json:"kind"`
	Color     *string  `json:"color"`
	Features  []string `json:"feature"`
}
