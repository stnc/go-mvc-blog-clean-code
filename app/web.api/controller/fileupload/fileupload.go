package fileupload

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"stncCms/app/domain/helpers/stnchelper"
	"strconv"
	"strings"

	"github.com/h2non/filetype"
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
	UploadFile(filest *multipart.FileHeader) (string, string)
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
//UploadFile standart upload
func (fu *fileUpload) UploadFile(filest *multipart.FileHeader) (string, string) {
	var uploadFilePath string = "public/upl/"
	var filename string
	var errorReturn string

	f, err := filest.Open()
	if err != nil {
		// c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		errorReturn = err.Error()
	}
	defer f.Close()

	if filest.Header != nil {

		size := filest.Size
		var size2 = strconv.FormatUint(uint64(size), 10)
		if size > int64(1024000*5) { // 1 MB
			// return "", errors.New("sorry, please upload an Image of 500KB or less")
			errorReturn = "Resim boyutu çok yüksek maximum 5 MB olmalıdır" + size2
			filename = "false"
		}

		filenameOrg := filest.Filename

		filenameExtension := filepath.Ext(filenameOrg)

		realFilename := strings.Split(filenameOrg, ".")

		realFilenameSlug := stnchelper.GenericName(realFilename[0], 50)

		filename = realFilenameSlug + filenameExtension

		out, err := os.Create(uploadFilePath + filename)

		if err != nil {
			log.Fatal(err)
			errorReturn = err.Error()
			filename = "false"
		}
		defer out.Close()
		_, err = io.Copy(out, f)
		if err != nil {
			log.Fatal(err)
			errorReturn = err.Error()
			filename = "false"
		}

		buf, _ := ioutil.ReadFile(uploadFilePath + filename)

		if filetype.IsImage([]byte(buf)) {
			filename = realFilenameSlug + filenameExtension
			//TODO: resim boyutlandırma gelecek
			//https://github.com/disintegration/imaging
		} else {
			path := uploadFilePath + filename
			err := os.Remove(path)
			if err != nil {
				errorReturn = err.Error()
			}
			errorReturn = filenameOrg + " gerçek bir resim dosyası değildir"
			filename = "false"
		}

		return filename, errorReturn
	} else {
		return "", ""
	}

}
