# Plaudren


Because the world definitely needed one more HTTP router implementation in Go! 🎉

## Features

- Simple and intuitive API (so simple that even you can probably use it)
- Support for nested routers
- HTTP method handlers (GET, POST, etc.)!
- Structured API implementation support 
- Custom request/response handling with `Data` and `Error` types 
- Path-based routing 

## Why This Router?

Because sometimes you just want to write your own router instead of using the perfectly good ones that already exist.

## Installation

```bash
go get github.com/bitspaceorg/plaudren 
```

## Quick Start

### Basic Router 

Create a simple router with a single endpoint:

```go
server := New(":8000")
router := NewRouter("/")

// Add a route handler (the fun part)
router.Get("/", func(w http.ResponseWriter, r *http.Request) (*Data, *Error) {
    return NewData("Hello,World!"), nil 
})

server.Register(router)
```

### Nested Routers 

You can create nested routers
```go
server := New(":8000")

parentRouter := NewRouter("/api")

childRouter := NewRouter("/")
childRouter.Get("/", func(w http.ResponseWriter, r *http.Request) (*Data, *Error) {
    return NewData("Child says hi!"), nil
})

parentRouter.Handle("/v1", childRouter)

server.Register(parentRouter)
```

### Structured API Implementation (For the Organized Code):

```go
type UserAPI struct {
    *Router
}

func (a *UserAPI) Register() {
    a.Router.Get("/", a.GetUsers)    
    a.Router.Post("/", a.CreateUser)
}

func (a *UserAPI) GetUsers(w http.ResponseWriter, r *http.Request) (*Data, *Error) {
    return nil, nil
}

func (a *UserAPI) CreateUser(w http.ResponseWriter, r *http.Request) (*Data, *Error) {
    return nil, NewError("What could go wrong?") 
}

server := New(":8000")
api := &UserAPI{
    Router: NewRouter("/users"),
}
server.Register(api)
```

## Route Handler Signature

Route handlers use this signature:

```go
func(w http.ResponseWriter, r *http.Request) (*Data, *Error)
```

## Available Methods

It supports all your favorite HTTP methods (well, most of them):

- `Get(path string, handler HandlerFunc)` 
- `Post(path string, handler HandlerFunc)`
- Too lazy to list the rest.

## Error Handling

This router uses custom error types:

```go
router.Get("/", func(w http.ResponseWriter, r *http.Request) (*Data, *Error) {
    if err := someOperation(); err != nil {
        return nil, &Error{
            Code:    http.StatusInternalServerError,
            Message: "Oops! Something went wrong (as usual)",
        }
    }
    return &Data{
        // your response data (assuming you have any)
    }, nil
})
```

### Middleware

This router supports middleware for both individual routes and entire routers:

```go
// Define a middleware function
func AuthMiddleware(w http.ResponseWriter, r *http.Request) *Error {
    if unauthorized := checkAuth(r); unauthorized {
        return NewError("Nice try!").SetCode(http.StatusUnauthorized)
    }
    return nil
}

// Apply middleware to a single route
router.Post("/secure", func(w http.ResponseWriter, r *http.Request) (*Data, *Error) {
    return NewData("secure data!"), nil
}).Use(AuthMiddleware)

// Or apply middleware to an entire router
router := NewRouter("/api").Use(AuthMiddleware)
```

Middleware Chaining :

```go
router.Post("/fort-knox", handler).
    Use(AuthMiddleware,RateLimiter).
```

### Error Handling in Middleware

Middleware can return Error for clean error handling:

```go
func MockMiddleware(_ http.ResponseWriter, r *http.Request) *Error {
    body := struct {
        Type int `json:"type"`
    }{}

    if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        return NewError("Invalid JSON").SetCode(http.StatusBadRequest)
    }

    if body.Type == 0 {
        return NewError("Invalid type").SetCode(http.StatusInternalServerError)
    }

    return nil  // All good!
}
```

### Serving Static Files

FileServer serves the entire directory.

```go
    router.ServeDir("/test",http.Dir("./test_dir"))
```


## Testing 

It may pass sometimes, if not try running again.

```sh
go test -v ./...
```

## Todo(Anything else create a issue)
- [x] File Server Handler for static files
- [ ] Templates handler(htmx baby...)
- [ ] CORS and cookie middleware(

## Contributing

Found a bug? Want to add a feature? Know how to actually write good code? I am all ears! Please feel free to submit a Pull Request. I promise to read it eventually. Please refer [CONTRIBUTING.md](./CONTRIBUTING.md) for more information.

## License

This project is licensed under the [GLWTSPL](https://github.com/suryaaprakassh/plaudren/blob/main/LICENSE).
