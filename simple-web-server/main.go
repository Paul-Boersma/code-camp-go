package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// File Servers provide a central location to store and share files for a network (static assets).
	// Protocols: SBM (Server Message Block), NFS (Network File Server), FTP (File Transfer Protocol), SFTP (Secure FTP)
	// A database is no File Server, as databases only deal with structured data accessed by a query.
	// Indicate the directory with files to be served
	// - NAS (Network attached storage) is dedicated file server hardware
	// - DMS (Document Management System) is a file server dedicated to storing documents (Word, PDF, but not blob data or videos)
	// If present, http.FileServer serves index.html by default, otherwise it shows the directory.
	// The FileServer is used to serve Front End files (HTML, CSS, JavaScript) instead of plain text when handling (dynamic) requests.
	fileServer := http.FileServer(http.Dir("static/"))

	// To access the root directory of the file server, provide the root URL.
	// With the go std http package, to access any file or directory within the root folder, provide the root URL: '/folder/somefile.html
	// With the gorilla mux package, to acces a file or directory provide the full path from the current working directory, regardless of the path requested
	r.Handle("/", fileServer)
	// ! Required when using Mux
	r.HandleFunc("/folder/", handleFolder)

	http.ListenAndServe(":8080", r)
}

func handleFolder(rw http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./static/folder/index.html"))
	tmpl.Execute(rw, nil)
}
