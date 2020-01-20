# SaaS swagger

Copyright 2019, Geeks Accelerator  
accelerator@geeksinthewoods.com.com


## Description

saas middleware to automatically generate RESTful API documentation with Swagger 2.0.


## Usage

### Start using it
1. Add comments to your API source code, [See Declarative Comments Format](https://github.com/geeks-accelerator/swag#declarative-comments-format).

2. Download [Swag](https://github.com/geeks-accelerator/swag) for Go by using:
    ```sh
    $ go get github.com/geeks-accelerator/swag/cmd/swag
    ```

3. Run the [Swag](https://github.com/geeks-accelerator/swag) in your Go project root folder which contains `main.go` file, [Swag](https://github.com/geeks-accelerator/swag) will parse comments and generate required files(`docs` folder and `docs/doc.go`).
    ```sh_ "github.com/swaggo/echo-swagger/v2/example/docs"
    $ swag init
    ```

4. Import following in your code:
    ```go
    import "geeks-accelerator/oss/saas-starter-kit/internal/mid/saas-swagger" // saas-swagger middleware
    ```

    **Canonical example:**
    
    ```go
    package main
    
    import (
        "context"
        "log"
        "net/http"
        "os"
        "os/signal"
        "syscall"
        "time"
    
        "geeks-accelerator/oss/saas-starter-kit/internal/mid"
        saasSwagger "geeks-accelerator/oss/saas-starter-kit/internal/mid/saas-swagger"
        _ "geeks-accelerator/oss/saas-starter-kit/internal/mid/saas-swagger/example/docs" // docs is generated by Swag CLI, you have to import it.
        "geeks-accelerator/oss/saas-starter-kit/internal/platform/web"
    )
    
    // @title SaaS Example API
    // @version 1.0
    // @description This is a sample server celler server.
    // @termsOfService http://geeksinthewoods.com/terms
    
    // @contact.name API Support
    // @contact.email support@geeksinthewoods.com
    // @contact.url https://gitlab.com/geeks-accelerator/oss/saas-starter-kit
    
    // @license.name Apache 2.0
    // @license.url http://www.apache.org/licenses/LICENSE-2.0.html
    
    // @host example-api.saas.geeksinthewoods.com
    // @BasePath /v1
    
    func main() {
    
        // Logging
        log := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
    
        // Configuration
        ... 
    
        // =========================================================================
        // Start API Service
        
        // Make a channel to listen for an interrupt or terminate signal from the OS.
        // Use a buffered channel because the signal package requires it.
        shutdown := make(chan os.Signal, 1)
        signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
    
        // Construct the web.App which holds all routes as well as common Middleware.
        app := web.NewApp(shutdown, log, mid.Trace(), mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics())
    
        app.Handle("GET", "/swagger/", saasSwagger.WrapHandler)
        app.Handle("GET", "/swagger/*", saasSwagger.WrapHandler)
    
        /*
            Or can use SaasWrapHandler func with configurations.
            url := saasSwagger.URL("http://localhost:1323/swagger/doc.json") //The url pointing to API definition
            e.GET("/swagger/*", saasSwagger.SaasWrapHandler(url))
        */
    
        ... 
    }
    ```

5. Run it, and browser to http://localhost:1323/swagger/index.html, you can see Swagger 2.0 Api documents.


### Dynamic Placeholders

To help ease use of the Swagger UI, dynamic placeholders have been added to the middleware. They are replaced on each 
request before the JSON is returned to the browser. These can be used in an `example` struct tag.   

1. `{RANDOM_UUID}`
    Generates a random UUID.

    Example:
    ```
    Name string `json:"name" validate:"required" example:"Company {RANDOM_UUID}"`
    ```

2. `{RANDOM_EMAIL}`
    Generate a random email address. Format will be UUID@example.com
    
    Example:
    ```
    Email string  `json:"email" validate:"required,email" example:"{RANDOM_EMAIL}"`
    ```