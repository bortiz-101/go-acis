# go-acis

A modern ACIS web application with a Go API and Angular frontend.

```text
go-acis/                              
├── .github/                         
│   └── workflows/                    
│       └── ci.yml                   
├── api/                              # where our go API code will live
│   ├── cmd/                          
│   │   └── server/                   
│   │       └── main.go               # entry point for the api / what actually starts our API 
│   ├── internal/                     
│   │   ├── config/                   # Application configuration package.
│   │   │   └── config.go             # will read env vars and turn them into a typed go objects/whatever we need them for
│   │   └── http/                    
│   │       ├── handlers/             # basically whenever a route is hit, what function handles it
│   │       │   └── health.go          # handler for our first endpoint
│   │       └── router.go              # define all of our endpoints and routes
│   ├── Dockerfile                    
│   ├── go.mod                        
│   └── go.sum                        
├── frontend/                         # where angular frontend will live. i want this to be as bare as possible until backend logic is polished
│   └── .gitkeep                      
├── .gitignore                       
├── compose.yaml                      
├── LICENSE                           
└── README.md                        
```
