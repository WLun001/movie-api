# movie-api
A simple API that retrieve trending movies data from [The movie database](https://www.themoviedb.org/), and store into MongoDB


## How to run

1. Download [Realize](https://github.com/oxequa/realize)
2. Make sure to add `go/bin` to env
3. Run `realize start`

## API endpoints
`base url = api/v1`

| Route  | Method | Description |
| ------------- | ------------- | ----------- |
| /save | GET | Save trending movies from [The movie database](https://www.themoviedb.org/) to database |
| /trending  | GET | Get trending movies| Get trending movies from database |
| /trending?title=keyword | GET | Full text search on movie title |
| /trending | DELETE | Delete all the trending movies data from database | 
