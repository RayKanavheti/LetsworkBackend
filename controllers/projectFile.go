package controllers

import (
	"encoding/json"
	// "fmt"
	"io"
	// "log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/raykanavheti/LetsworkBackend/controllers/util"
	"github.com/raykanavheti/LetsworkBackend/models"
)

// ProjectFileController controller
type ProjectFileController struct{}

var docs []models.ProjectFile

// CreateProjectFiles creates a new education for a user
func (fileCntrl *ProjectFileController) CreateProjectFiles(w http.ResponseWriter, r *http.Request) {
	responseWriter := util.GetResponseWriter(w, r)
	vars := mux.Vars(r)
	ProjIDRaw := vars["id"]
	ProjID, _ := strconv.Atoi(ProjIDRaw)
	//parse the multipart form in the request
	err := r.ParseMultipartForm(200000)
	if err != nil {
		mapError := map[string]string{"message": err.Error()}
		errj, _ := json.Marshal(mapError)
		responseWriter.WriteHeader(400)
		responseWriter.Write(errj)
	}
	m := r.MultipartForm
	//get the *fileheaders
	files := m.File["files"]
	for i, _ := range files {
		//for each fileheader, get a handle to the actual file
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			// log.Fatal(err)
			// fmt.Fprintln(w, err)
			// return
			mapError := map[string]string{"message": err.Error()}
			errj, _ := json.Marshal(mapError)
			responseWriter.WriteHeader(400)
			responseWriter.Write(errj)
		}
		//create destination file making sure the path is writeable.
		// dst, err := os.Create("/tmp/" + files[i].Filename)
		dst, err := os.Create("./static/userfiles/" + files[i].Filename)
		defer dst.Close()
		if err != nil {
			// fmt.Fprintf(w, "Unable to create file for writing. Check your write access priviledge")
			// return
			mapError := map[string]string{"message": err.Error()}
			errj, _ := json.Marshal(mapError)
			responseWriter.WriteHeader(400)
			responseWriter.Write(errj)
		}
		//copy the uploaded file to the destination file
		_, err1 := io.Copy(dst, file)
		if err1 != nil {
			// fmt.Fprintln(w, err1)
			// return
			mapError := map[string]string{"message": err1.Error()}
			errj, _ := json.Marshal(mapError)
			responseWriter.WriteHeader(400)
			responseWriter.Write(errj)
		}
		// fmt.Fprintf(w, "Files uploaded successfully: ")
		// fmt.Fprintf(w, files[i].Filename+"\n")
		docs = append(docs, models.ProjectFile{Path: r.Host + "/static/userfiles/" + files[i].Filename, ProjectID: ProjID})

	}
	pathArray, err2 := models.CreateFiles(docs)
	if err2 == nil {
		uj, _ := json.Marshal(pathArray)
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(201)
		responseWriter.Write(uj)
	} else {
		mapError := map[string]string{"message": err2.Error()}
		err2, _ := json.Marshal(mapError)
		responseWriter.WriteHeader(400)
		responseWriter.Write(err2)
	}
}
