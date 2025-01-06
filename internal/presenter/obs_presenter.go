package presenter

type FileUploadSuccessResponse struct {
	Filename string `json:"filename"`
	FileType string `json:"fileType"`
	Url      string `json:"url"`
	Key      string `json:"key"`
}

type FileLengthResponse struct {
	FileName string `json:"fileName"`
	Url      string `json:"url"`
	Length   int    `json:"length"`
}
