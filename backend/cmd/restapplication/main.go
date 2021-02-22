package main

import (
	"flag"
	"fmt"
	"go/constant"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var (
	servicePort *string
	accessKey   *string
	secretKey   *string
	token       *string
	region      *string
)

func init() {
	servicePort = flag.String("p", strings.TrimSpace(os.Getenv("SERVICE_PORT")), "service port")
	accessKey = flag.String("ak", strings.TrimSpace(os.Getenv("AWS_S3_ACCESS_KEY")), "aws s3 access key")
	secretKey = flag.String("sk", strings.TrimSpace(os.Getenv("AWS_S3_SECRET_KEY")), "aws s3 secret key")
	token = flag.String("tok", strings.TrimSpace(os.Getenv("AWS_S3_TOKEN")), "aws s3 token")
	region = flag.String("reg", strings.TrimSpace(os.Getenv("AWS_S3_REGION")), "aws s3 region")
	flag.Parse()
}

type App struct {
	// S3Access *
}

func (app *App) listBucketContent(w http.ResponseWriter, r *http.Request) {

}

func main() {
	log.Printf("Initializing server\n")
	r := mux.NewRouter()

	// s3Client, err := s3client.S3Service(accessKey, secretKey, token, region)
	// if err != nil {
	// 	log.Fatal("Falied to connect aws s3")
	// }
	constants.
	t := App{}
	r.HandleFunc("/list-bucket-content/{param:.*}", t.listBucketContent)
	log.Printf("Server listening on port : %s", *servicePort)
	http.ListenAndServe(fmt.Sprintf(":%s", *servicePort), r)
}
