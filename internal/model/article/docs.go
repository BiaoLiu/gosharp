package article

// swagger:parameters articleReqWrap
type articleReqWrap struct {
	// in:body
	Body ArticleReq
}

// swagger:response articleRespWrap
type articleRespWrap struct {
	// in: body
	Body ArticleResp
}
