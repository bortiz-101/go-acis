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

## what we have so far

- a Go server using Gin with health and station data endpoints
- an ACIS client with request/response structs and configurable URL and timeout
- a decoder for numbers, missing data, trace amounts, and accumulated values
- basic error handling and a notebook for trying requests
- GitHub Actions for linting, security checks, builds, and dependency updates

## first goal

answer one useful question from start to finish:

> how many days did Raleigh hit at least 95 degrees last summer?

start with a form for picking a station, dates, and temperature threshold. show the answer and the observations behind it, including any missing days.

## next steps

1. fix the basics: verify the ACIS URL and endpoint path, fix the timeout fallback, and correct the Docker build path.
2. validate station requests: require a station and weather element, check dates, and return clear errors.
3. test the decoder: cover normal values, missing data, trace amounts, accumulated values, and empty or unexpected rows.
4. get the example question working through the API using a known station ID and explicit dates. return the count, matching dates, and missing-day count.
5. build a simple page that submits those choices and displays the result.
6. add station search so users don't need to know station IDs.

## what comes after that

- add charts and downloadable results
- show station details, units, and available dates
- support more questions: rainfall totals, temperature comparisons, records, and streaks
- keep question-answering logic separate from HTTP handlers and the ACIS client
- use sample ACIS responses for tests so they work without an internet connection

## plain language queries eventually

let someone type:

> was last July hotter than usual in Raleigh?

use AI to turn the question into choices our app already supports. show the station, dates, measurement, and comparison before running it. ask when something is unclear and let the user adjust the choices.

our Go code will validate the query, fetch the data, and calculate the answer through the same backend used by the form.

## ideas for later

- compare stations, years, and seasons
- shareable and saved searches
- cache repeated requests if needed
- add a database when we have something worth storing
