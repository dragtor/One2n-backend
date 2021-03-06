package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dragtor/One2n-backend/backend/controller"
	"github.com/dragtor/One2n-backend/backend/pkg"
	"github.com/gorilla/mux"
)

var (
	servicePort  *string
	awsAccessKey *string
	awsSecret    *string
	token        *string
	region       *string
)

func init() {
	servicePort = flag.String("p", strings.TrimSpace(os.Getenv("SERVICE_PORT")), "service port")
	awsAccessKey = flag.String("ak", strings.TrimSpace(os.Getenv("AWS_S3_ACCESS_KEY")), "aws s3 access key")
	awsSecret = flag.String("sk", strings.TrimSpace(os.Getenv("AWS_S3_SECRET_KEY")), "aws s3 secret key")
	token = flag.String("tok", strings.TrimSpace(os.Getenv("AWS_S3_TOKEN")), "aws s3 token")
	region = flag.String("reg", strings.TrimSpace(os.Getenv("AWS_S3_REGION")), "aws s3 region")
	flag.Parse()
}

type App struct {
	AppCntrl *controller.Controller
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type HttpResponseStruct struct {
	Contents []string `json:"contents"`
}

func convertToHttpResponse(resp *controller.ControllerResponse) *HttpResponseStruct {
	return &HttpResponseStruct{
		Contents: resp.LsDir,
	}
}

func validateRequest(param string) error {
	return nil
}

func (app *App) listBucketContent(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request to fetch query")
	param := mux.Vars(r)["param"]

	err := validateRequest(param)
	if err != nil {
		log.Printf("Error : Failed to decode request data ")
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	resp, err := app.AppCntrl.CommandS3ls(param)
	if err != nil {
		log.Printf("Error : Failed to Fetch Info")
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	httpresp := convertToHttpResponse(resp)
	respondWithJSON(w, http.StatusOK, httpresp)
}

func main() {
	log.Printf("Initializing server \n")
	r := mux.NewRouter()
	s3iter, err := pkg.NewS3Service(*awsAccessKey, *awsSecret, *token, *region)
	if err != nil {
		fmt.Errorf("error")
		// return err
		return
	}
	controller := controller.NewController(s3iter)
	t := App{AppCntrl: controller}
	r.HandleFunc("/list-bucket-content/{param:.*}", t.listBucketContent)
	log.Printf("Server listening on port : %s", *servicePort)
	http.ListenAndServe(fmt.Sprintf(":%s", *servicePort), r)
}
