package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
)

func (server *Server) uploadFileHandler(c *gin.Context) {
	s3Config := &aws.Config{
		Region:      aws.String("ap-southeast-2"),
		Credentials: credentials.NewStaticCredentials("AKIAXS2AXJQN6D7EYJK5", "PwZWdvNT3LWttTsj+aWyB4uX7nRNEyhr9PyFOufk", ""),
	}
	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)
	downloader := s3manager.NewDownloader(s3Session)

	file, err := c.FormFile("file")
	f, err := file.Open()
	if err != nil {
		// Handle error
	}
	defer f.Close()

	// Get the file size and read the entire file into a byte slice
	size := file.Size
	fileBytes := make([]byte, size)
	io.ReadFull(f, fileBytes)

	input := &s3manager.UploadInput{
		Bucket:      aws.String("simple-bank"),  // bucket's name
		Key:         aws.String("khai/txt"),     // files destination location
		Body:        bytes.NewReader(fileBytes), // content of the file
		ContentType: aws.String("text/plain"),   // content type
	}
	output, err := uploader.UploadWithContext(context.Background(), input)
	fmt.Println(*output)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	localFilePath := "file.txt"
	buf := &aws.WriteAtBuffer{}
	// Download the file from S3 and save the content to the local file
	downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String("simple-bank"),
		Key:    aws.String("khai/txt"),
	})
	err = ioutil.WriteFile(localFilePath, buf.Bytes(), 0644)
	fmt.Println(*buf)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("File %s uploaded successfully", file.Filename),
	})
}
