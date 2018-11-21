package forms

// swagger:parameters swagRouteFormWrap
type swagRouteFormWrap struct {
	//in:body
	Body SwagRouteForm
}

type SwagRouteForm struct {
	BaseForm
	Url string `json:"url"`
}
