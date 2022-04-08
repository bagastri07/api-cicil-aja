package model

type DataResponse struct {
	Data interface{} `json:"data"`
}

type MessageResponse struct {
	Message interface{} `json:"messages"`
}

type MessageDataResponse struct {
	Message interface{} `json:"messages"`
	Data    interface{} `json:"data"`
}

type OauthResp struct {
}
