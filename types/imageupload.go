package types

type ImageUploadRsp struct {
	ImageUploadRspData
	Data  CommonData `json:"data"`
	Extra ExtraData  `json:"extra"`
}

type ImageUploadRspData struct {
	Height  int    `json:"height"`
	ImageId string `json:"image_id"`
	Md5     string `json:"md5"`
	Width   int    `json:"width"`
}
