# Custom Web Server Implementation in Go

This is a custom web server implemented in Go language. It's a TCP server where communication at the application layer is implemented following the HTTP protocol and all its rules. The architecture of the web server follows the principles of the Clean Architecture, and adheres to its rules. Additional packages are minimized, and the entire communication, as well as routing, is done manually. The only additional library used is Gorm as an ORM tool for integration with the MySQL database.

## Clone the project

```bash
$ git clone https://github.com/perasd9/MT-GoWebServer.git
```

For starting backend you need type in console 'go mod tidy' to get all packages.
After that you just can start backend with 'go run main.go'.
