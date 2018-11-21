package serializers

// swagger:response swagRouteResponseWrap
type swagRouteResponseWrap struct {
	//in:body
	Body SwagRouteResponse
}

type SwagRouteResponse struct {
	//url路径
	Url string `json:"url"`
}
