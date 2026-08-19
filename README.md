# go-acis

## Project structure

```text
go-acis/
├── .github/                  
│   └── workflows/            
│       └── ci.yml
├── cmd/                      
│   └── go-acis/              
│       └── main.go           # entry point for the api / what actually starts our API
├── internal/                 # private business logic and app implementation.
│   ├── acis/                 # ACIS logic, API client and request/response types
│   │   ├── client.go         # reusable HTTP client for communicating with ACIS
│   │   ├── payload.go        # ACIS request types
│   │   └── response.go       # ACIS response types
│   ├── config/               # cant think of how exactly we will use this but seems like mainstream go APIs all have this folder
│   │   └── config.go         # will read env vars and turn them into a typed go objects/whatever we need them for
│   └── http/                 
│       ├── handlers/         # our handlers manage the http request coming TO the API
│       │   └── health.go     # handler for our first endpoint
│       └── router.go         # define all of our endpoints and routes
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```
