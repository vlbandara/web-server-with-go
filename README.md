# Go Web Server with Static Files

This repository contains a basic web server written in Go. The server includes advanced features such as routing with `gorilla/mux`, middleware for logging requests, serving JSON responses, and static file handling.

## Features

- Simple HTTP server using Go's net/http package
- Routing with `gorilla/mux`
- Middleware for logging HTTP requests
- JSON API endpoints
- Serving static files

## Prerequisites

- Go (version 1.16 or higher)
- Git

## Setup

1. Clone the repository to your local machine:

    ```sh
    git clone https://github.com/yourusername/web-server-with-go.git
    cd web-server-with-go
    ```

2. Initialize the Go module:

    ```sh
    go mod init example.com/webserver
    ```

    Replace `example.com/webserver` with your module path if necessary.

3. Install the dependencies:

    ```sh
    go get github.com/gorilla/mux@latest
    ```

## Running the Server

1. Make sure you are in the project directory.
2. Run the server:

    ```sh
    go run main.go
    ```

3. The server will start on port `8080` by default. You can access it at `http://localhost:8080`.

## Endpoints

### Root Endpoint

- **URL**: `/`
- **Method**: `GET`
- **Response**: `"Hello, World!"`

### Greet Endpoint

- **URL**: `/greet`
- **Method**: `GET`
- **Query Parameter**: `name` (optional)
- **Response**: `"Hello, [name]!"` (defaults to "Guest" if no name is provided)

### JSON Endpoint

- **URL**: `/api/json`
- **Method**: `GET`
- **Response**: JSON object with a message

    ```json
    {
        "message": "Hello, JSON World!"
    }
    ```

### Static Files

Static files are served from the `static` directory. You can access them at the `/static` URL path.

Example: Accessing a file named `index.html` in the `static` directory:

- **URL**: `/static/index.html`

### Example Static File

Here is an example of a simple, attractive HTML file you can place in the `static` directory (`static/index.html`):

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome to My Go Server</title>
    <style>
        body {
            font-family: 'Arial', sans-serif;
            background-color: #f0f0f0;
            margin: 0;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            color: #333;
        }
        .container {
            text-align: center;
            background: white;
            padding: 30px;
            border-radius: 10px;
            box-shadow: 0 4px 8px rgba(0,0,0,0.1);
        }
        h1 {
            color: #4CAF50;
        }
        button {
            background-color: #4CAF50;
            color: white;
            border: none;
            padding: 10px 20px;
            border-radius: 5px;
            cursor: pointer;
            font-size: 16px;
        }
        button:hover {
            background-color: #45a049;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Welcome to My Go Server</h1>
        <p>This is a simple static file served by your Go web server.</p>
        <button onclick="showMessage()">Click Me</button>
        <p id="message" style="display:none; margin-top:20px;">Hello, Go Developer!</p>
    </div>
    <script>
        function showMessage() {
            document.getElementById('message').style.display = 'block';
        }
    </script>
</body>
</html>


Check the tutorial on my blog - https://www.aggrandizer.info/2024/06/simple-web-server-with-go.html
