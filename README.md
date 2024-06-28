# Web Gallery

Web Gallery is a simple Go web application that serves images from a mounted directory. It resizes the images on the fly and displays them in a web gallery format.

## Features

- Serves images from a specified directory.
- Resizes images on the fly using bimg.
- Organizes images by year and month.
- Uses HTML templates to render the gallery.

## Requirements
- Go 1.16 or later
- pkg-config and libvips-dev for building the project
- libvips for running the project

# Installation
Before building the project, you need to install pkg-config and libvips-dev on your development machine.

```
sudo apt update
sudo apt install -y pkg-config libvips-dev
```

# Building the Project

1. Clone the repository:
```
git clone https://github.com/yourusername/webgallery.git
cd webgallery
```

2. Initialize the Go module:
```
go mod init webgallery
go mod tidy
```

3. Build the project:
```
go build -o bin/webgallery ./cmd/webgallery
```

4. Copy the necessary files to the target machine:
```
scp -r bin/webgallery templates static username@targetmachine:/path/to/webgallery
```

5. Run the application:
```
cd /path/to/webgallery
./webgallery
```

Open your web browser and navigate to http://localhost:8080 to view the gallery.

## Directory Structure
```
webgallery/
├── bin/
│   └── webgallery          # Compiled binary
├── cmd/
│   └── webgallery/
│       └── main.go         # Entry point of the application
├── internal/
│   └── gallery/
│       ├── handler.go      # HTTP handlers
│       └── gallery.go      # Image processing and loading
├── templates/
│   └── gallery.html        # HTML template for the gallery
├── static/
│   └── images/             # Directory for storing images
├── go.mod
└── go.sum
```

## Folder Pattern for Storing Images
The application expects images to be stored in a specific folder pattern within the mounted directory. The pattern should be year/month/image. For example:

```
/mnt/external/
├── 2023/
│   ├── 01-Jan/
│   │   ├── image1.jpg
│   │   ├── image2.jpg
│   ├── 02-Feb/
│       ├── image3.jpg
│       ├── image4.jpg
└── 2024/
    ├── 01-Jan/
    │   ├── image5.jpg
    │   ├── image6.jpg
    ├── 02-Feb/
        ├── image7.jpg
        ├── image8.jpg
```

## Configuration
The application expects images to be in a directory mounted at */mnt/external.*

You can change this path and the server port in main.go:

```go
const imageDir = "/mnt/external" // Path to the mounted directory
const port = 8080
```