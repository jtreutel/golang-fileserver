# Simple Golang Fileserver

This is a super simple fileserver written to help me learn Golang.

### Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   └── handlers/
│       └── file_handler.go
└── go.mod
```


### Local Testing

**Setting up the environment**
1. Build the Docker image: `docker build -t golang-filestorage .`
2. Create a local file upload directory if you haven't already: `mkdir uploaded_files`
3. Run the container: `docker run golang-filestorage -d -p 8080:8080 -v $(pwd)/uploaded_files:/root/uploaded_files`
   - Runs the container in the background with port mapping 8080->8080 and a local volume `uploaded_files`

**Uploading a file**  
`curl -X POST -F "file=@path/to/file/FILENAME.txt" http://localhost:8080/files/FILENAME.txt`

**Listing files**  
`curl http://localhost:8080/files`

**Deleting a file**  
`curl -X DELETE http://localhost:8080/files/FILENAME.txt`


### Screenshot of local testing

<img src="example_usage.png" alt="Example usage" width="800"/>
