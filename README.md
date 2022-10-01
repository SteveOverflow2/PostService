# Globe Protocol Go Template

## Project
This project is a template which is meant as a simple version of what future project structures should look like.
Within new projects the following things should be changed:

## Change / Delete
- New go.mod
- Project name 
- Package names upon change/deletion of examples. (Don't forget import names)
- .env file variables (Name, port, any new ones created)
- Dockerfile port.
- Routes in pkg\http\rest\router.go
- Handlers in pkg\http\rest\handlers
- Logic in pkg\post
- Storage in pkg\storage
- Pipeline setup (ask DevOps)
- Azure-pipelines YAML (ask DevOps)

## Flow
1. Project starts in main.go
2. In pkg\config\config.go, it loads all configuration set up in pkg\config.
This config based on .env file. Should something be excluded/forgotten/mistaken, the terminal will provide a notification.
3. An HTTP server is created* in pkg\http\rest\server.go, using HTTP Config, a new PostService instance and the Gorilla Mux router
4. Routes are defined in pkg\http\rest\router.go
5. Interaction to the afformentioned routes will route to methods in pkg\http\rest\handlers\post\handler.go
6. The handlers call to pkg\post\interface.go. This class contains interfaces for pkg\post\logic.go & pkg\storage\example_database\post\post.go
7. pkg\post\logic.go implements the interface from 6, and therefore will process the calls from 5. 
8. If needed, pkg\post\logic.go will call upon pkg\post\interface.go, to go to pkg\storage\example_database\post\post.go. Here database calls are made.

*it's dependencies are created in the init() function

## Other Details
- Use pkg\post\interface.go to generate mock classes. 
- Remember to write tests
- Make new packages if adding a new flow 'domain', such as pkg\http\rest\handlers\coupons, pkg\coupons & pkg\storage\coupons
- Use the pkg\http\rest\handlers\response.go package in the Handlers to response
- Use the pkg\util\error.go package in the Handlers for errors where necessary
- CORS has been configured in the router using pkg\http\rest\middleware.go

### Base .env file:
```
    # Base env fields:
    APP_NAME=Go Template
    APP_VERSION=v1

    HTTP_HOST=localhost
    HTTP_PORT=8080
    HTTP_ALLOWED_ORIGIN=http://localhost:3000
    HTTP_WRITE_TIME_OUT=150s
    HTTP_READ_TIME_OUT=150s
    HTTP_IDLE_TIME_OUT=150

    # When other env fields are missing the app will notify you on runtime
```