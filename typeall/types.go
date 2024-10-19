package typeall

type PrintDataPDFToImageReq struct {
	ID                    *string `json:"id"`
	FilePDFUrl            *string `json:"filePDFUrl"`
	FilePDFImageUrl       *string `json:"filePDFImageUrl"`
	FilePDFImageUploadUrl *string `json:"filePDFImageUploadUrl"`
}

type PrintDataImageFromPDFResp struct {
	ID              *string `json:"id"`
	Status          *int    `json:"status"` // 1：成功，0：失败
	Message         *string `json:"message"`
	FilePDFImageUrl *string `json:"filePDFImageUrl"`
}
