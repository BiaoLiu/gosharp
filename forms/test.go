package forms

import "gosharp/utils/app"

// swagger:parameters swagRouteFormWrap
type swagRouteFormWrap struct {
	//in:body
	Body SwagRouteForm
}

type SwagRouteForm struct {
	app.BaseForm
	Url string `json:"url"`
}
