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
│   ├── acis/      
│   ├── config/               # structs and data types for ACIS
│   │   └── config.go         # will read env vars and turn them into a typed go objects/whatever we need them for
│   └── http/                 
│       ├── handlers/         # basically whenever a route is hit, what function handles it
│       │   └── health.go     # handler for our first endpoint
│       └── router.go         # define all of our endpoints and routes
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```
