package stncupload

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/minio/minio-go/v6"
)

//NewFileUpload  construct s
func NewFileUpload() *fileUpload {
	return &fileUpload{}
}

type fileUpload struct{}

//UploadFileInterface interface
type UploadFileInterface interface {
	UploadFileForMinio(file *multipart.FileHeader) (string, error)
	UploadFile(filest *multipart.FileHeader, originalName string) (string, string)
	MultipleUploadFile(filest []*multipart.FileHeader, originalName string)
}

//So what is exposed is Uploader
var _ UploadFileInterface = &fileUpload{}

//	"github.com/minio/minio-go/v6"
func (fu *fileUpload) UploadFileForMinio(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", errors.New("cannot open file")
	}
	defer f.Close()

	size := file.Size
	//The image should not be more than 500KB
	fmt.Println("the size: ", size)
	if size > int64(512000) {
		return "", errors.New("sorry, please upload an Image of 500KB or less")
	}
	//only the first 512 bytes are used to sniff the content type of a file,
	//so, so no need to read the entire bytes of a file.
	buffer := make([]byte, size)
	f.Read(buffer)
	fileType := http.DetectContentType(buffer)
	//if the image is valid
	if !strings.HasPrefix(fileType, "image") {
		return "", errors.New("please upload a valid image")
	}
	filePath := FormatFile(file.Filename)

	accessKey := os.Getenv("DO_SPACES_KEY")
	secKey := os.Getenv("DO_SPACES_SECRET")
	endpoint := os.Getenv("DO_SPACES_ENDPOINT")
	ssl := true

	// Initiate a client using DigitalOcean Spaces.
	client, err := minio.New(endpoint, accessKey, secKey, ssl)
	if err != nil {
		log.Fatal(err)
	}
	fileBytes := bytes.NewReader(buffer)
	cacheControl := "max-age=31536000"
	// make it public
	userMetaData := map[string]string{"x-amz-acl": "public-read"}
	n, err := client.PutObject("chodapi", filePath, fileBytes, size, minio.PutObjectOptions{ContentType: fileType, CacheControl: cacheControl, UserMetadata: userMetaData})
	if err != nil {
		fmt.Println("the error", err)
		return "", errors.New("something went wrong")
	}
	fmt.Println("Successfully uploaded bytes: ", n)
	return filePath, nil
}

//TODO: https://github.com/gin-gonic/examples/tree/master/upload-file upload ornekleri var
//TODO: gerçek resim dosayasını tespit eden fonksiyon başka yere alınablir
//TODO: boyutlandırma https://github.com/disintegration/imaging
//https://socketloop.com/tutorials/golang-how-to-verify-uploaded-file-is-image-or-allowed-file-types
//https://www.golangprograms.com/how-to-get-dimensions-of-an-image-jpg-jpeg-png-or-gif.html
//UploadFile standart upload
func (fu *fileUpload) UploadFile(filest *multipart.FileHeader, originalName string) (filename, errorReturn string) {
	var uploadFilePath string = "public/upl/"
	// var filename string
	// var errorReturn string

	if filest != nil {
		f, err := filest.Open()
		defer f.Close()
		if err != nil {
			errorReturn = err.Error()
		}

		if filest.Header != nil {

			size := filest.Size
			var size2 = strconv.FormatUint(uint64(size), 10)
			if size > int64(1024000*5) { // 1 MB
				errorReturn = "Resim boyutu çok yüksek maximum 5 MB olmalıdır" + size2
				filename = "false"
			}

			newFileName, filename := newFileName(uploadFilePath, filest.Filename)

			out, err := os.Create(newFileName)

			defer out.Close()

			if err != nil {
				log.Fatal(err)
				errorReturn = err.Error()
				filename = "false"
			}

			_, err = io.Copy(out, f)

			if err != nil {
				log.Fatal(err)
				errorReturn = err.Error()
				filename = "false"
			}

			ret := realImage(newFileName)
			if ret != "true" {
				errorReturn = ret
				filename = "false"
				//TODO: bu kısım veritabanına gitsin daha sonra silsin gibi bişey olacak
				// errFile := os.Remove(newFileName)
				// if errFile != nil {
				// 	fmt.Println("errFile.Error()")
				// 	fmt.Println(errFile.Error())
				// 	errorReturn = errFile.Error()
				// 	return filename, errorReturn
				// }

			}

			return filename, errorReturn
		} else {
			return originalName, ""
		}
	} else {
		return originalName, ""
	}

}

func (fu *fileUpload) MultipleUploadFile(files []*multipart.FileHeader, originalName string) {
	// var uploadFilePath string = "public/upl/"

	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			// fmt.Fprintln(w, err)
			return
		}

		out, err := os.Create("tmp/" + files[i].Filename)

		defer out.Close()
		if err != nil {
			// fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
			return
		}

		_, err = io.Copy(out, file) // file not files[i] !

		if err != nil {
			// fmt.Fprintln(w, err)
			return
		}

		fmt.Println("Files uploaded successfully : ")
		fmt.Println(files[i].Filename + "\n")

	}

}

func realImage(fileName string) string {
	var errorReturn string
	// open the uploaded file
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		errorReturn = err.Error()
		// os.Exit(1)
	}

	buff := make([]byte, 512) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
	_, err = file.Read(buff)

	if err != nil {
		fmt.Println(err)
		errorReturn = err.Error()
		// os.Exit(1)
	}

	filetype := http.DetectContentType(buff)

	switch filetype {
	case "image/jpeg", "image/jpg":
		errorReturn = "true"

	case "image/gif":
		errorReturn = "true"

	case "image/png":
		errorReturn = "true"

	// case "application/pdf": // not image, but application !
	// 	fmt.Println(filetype)
	default:
		errorReturn = " gerçek bir resim dosyası değildir"
	}

	return errorReturn
}
