package utils

// import (
// 	"bytes"
// 	"fmt"
// 	"furnirobackend/internal/presenter"
// 	"mime/multipart"
// 	"net/http"
// 	"sync"

// 	obs "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
// )

// const (
// 	chunkSize = 40 * 1024 * 1024 //40 MB per part
// 	workers   = 10
// )

// func UploadFile(objDir string, file *multipart.FileHeader) (*presenter.FileUploadSuccessResponse, error) {
// 	obsClient, buffer, err := UploadFileHandler(file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	key := objDir + "/" + file.Filename

// 	input := &obs.InitiateMultipartUploadInput{}
// 	input.Bucket = config.BucketName
// 	input.Key = key
// 	output, err := obsClient.InitiateMultipartUpload(input)
// 	if err != nil {
// 		return nil, err
// 	}
// 	uploadID := output.UploadId

// 	var parts []obs.Part
// 	var wg sync.WaitGroup
// 	partNumber := 1
// 	start := 0

// 	for start < buffer.Len() {
// 		end := start + chunkSize
// 		if end > buffer.Len() {
// 			end = buffer.Len()
// 		}

// 		wg.Add(1)
// 		go func(start, end, partNumber int) {
// 			defer wg.Done()

// 			partInput := &obs.UploadPartInput{
// 				Bucket:     config.BucketName,
// 				Key:        key,
// 				UploadId:   uploadID,
// 				PartNumber: partNumber,
// 				Body:       bytes.NewReader(buffer.Bytes()[start:end]),
// 			}

// 			partOutput, err := obsClient.UploadPart(partInput)
// 			if err != nil {
// 				fmt.Printf("Error uploading part %d: %v\n", partNumber, err)
// 				return
// 			}

// 			parts = append(parts, obs.Part{PartNumber: partNumber, ETag: partOutput.ETag})

// 		}(start, end, partNumber)

// 		start = end
// 		partNumber++

// 		if partNumber%workers == 0 {
// 			wg.Wait()
// 		}
// 	}

// 	wg.Wait()
// 	completeInput := &obs.CompleteMultipartUploadInput{
// 		Bucket:   config.BucketName,
// 		Key:      key,
// 		UploadId: uploadID,
// 		Parts:    parts,
// 	}

// 	_, err = obsClient.CompleteMultipartUpload(completeInput)
// 	if err != nil {
// 		fmt.Printf("Error completing multipart upload: %v\n", err)
// 		return nil, err
// 	}

// 	fileURL := fmt.Sprintf("https://%s.%s/%s", config.BucketName, config.Endpoint, key)
// 	response := &presenter.FileUploadSuccessResponse{
// 		Filename: file.Filename,
// 		FileType: http.DetectContentType(buffer.Bytes()),
// 		Url:      fileURL,
// 		Key:      key,
// 	}
// 	return response, nil
// }
