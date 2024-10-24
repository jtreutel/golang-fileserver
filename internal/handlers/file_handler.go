package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Set the file storage directory
const fileStorageDir = "uploaded_files"

// Ensure storage directory exists
func init() {
	if _, err := os.Stat(fileStorageDir); os.IsNotExist(err) {
		os.Mkdir(fileStorageDir, os.ModePerm)
	}
}

// Handle file upload and deletion
func FileOperations(w http.ResponseWriter, r *http.Request) {
	// Extract file name from the URL path
	fileName := strings.TrimPrefix(r.URL.Path, "/files/")
	if fileName == "" {
		http.Error(w, "File name is required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPost:
		uploadFile(w, r, fileName)
	case http.MethodDelete:
		deleteFile(w, fileName)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Upload a file to the server
func uploadFile(w http.ResponseWriter, r *http.Request, fileName string) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Could not parse multipart form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Could not read uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create the file in the storage directory
	out, err := os.Create(filepath.Join(fileStorageDir, fileName))
	if err != nil {
		http.Error(w, "Could not create file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Copy the uploaded file to the server's storage
	if _, err := io.Copy(out, file); err != nil {
		http.Error(w, "Could not save file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "File uploaded successfully: %s", fileName)
}

// Delete a file from the server
func deleteFile(w http.ResponseWriter, fileName string) {
	filePath := filepath.Join(fileStorageDir, fileName)

	if err := os.Remove(filePath); err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
		} else {
			http.Error(w, "Could not delete file", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File deleted successfully: %s", fileName)
}

// List all uploaded files
func ListFiles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Open the directory and list files
	files, err := os.ReadDir(fileStorageDir)
	if err != nil {
		http.Error(w, "Could not read directory", http.StatusInternalServerError)
		return
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	// Write the list of files as the response
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", fileNames)
}
