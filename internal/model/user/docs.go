package user

// swagger:parameters loginReqWrap
type loginReqWrap struct {
	// in:body
	Body LoginReq
}

// swagger:response userRespWrap
type userRespWrap struct {
	// in: body
	Body UserResp
}
