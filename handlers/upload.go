package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"static-site-hosting/models"
	"static-site-hosting/services"
)

var deployments []models.Deployment

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(50 << 20) // 50MB
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Missing file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	siteName := r.FormValue("site_name")
	if siteName == "" {
		http.Error(w, "Missing site name", http.StatusBadRequest)
		return
	}

	sitePath := filepath.Join("sites", siteName)
	os.MkdirAll(sitePath, os.ModePerm)

	fileCount, err := services.ExtractZip(file, sitePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to extract zip: %v", err), http.StatusInternalServerError)
		return
	}

	deployment := models.Deployment{
		SiteName:   siteName,
		DeployedAt: time.Now(),
		FileCount:  fileCount,
	}
	deployments = append(deployments, deployment)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":     "Deployment successful",
		"site":        siteName,
		"preview_url": fmt.Sprintf("http://%s/sites/%s/", r.Host, siteName),
	})
}

func ListDeployments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deployments)
}
