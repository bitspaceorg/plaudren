# Yet Another Amateur Router Implementation (plaudren)

## Dont ask what's the name supposed to mean

Because the world definitely needed one more HTTP router implementation in Go! 🎉

## Features

- Simple and intuitive API (so simple that even you can probably use it)
- Support for nested routers
- HTTP method handlers (GET, POST, etc.)!
- Structured API implementation support (fancy words for "organizing your code")
- Custom request/response handling with `ApiData` and `ApiError` types (because error handling should be an adventure)
- Path-based routing (I didn't get too creative here)

## Why This Router?

Because sometimes you just want to write your own router instead of using the perfectly good ones that already exist.

## Installation

```bash
go get github.com/suryaaprakassh/plaudren # Warning: Amateur code ahead
```

## Quick Start

### Basic Router (The "Hello World" of Routing)

Create a simple router with a single endpoint (because we all have to start somewhere):

```go
server := New(":8000")  // Yes, I hardcoded the port. Fight me.
router := NewRouter("/")

// Add a route handler (the fun part)
router.Get("/", func(w http.ResponseWriter, r *http.Request) (*ApiData, *ApiError) {
    return NewApiData("Hello,World!"), nil  // Look ma, no errors!
})

// Register the router (it's official now)
server.Register(router)
```

### Nested Routers (Inspired by Inception)

You can create nested routers, because I heard you like routers in your routers:

```go
server := New(":8000")

// Create parent router
parentRouter := NewRouter("/api")

// Create child router
childRouter := NewRouter("/")
childRouter.Get("/", func(w http.ResponseWriter, r *http.Request) (*ApiData, *ApiError) {
    return NewApiData("Child says hi!"), nil
})

// Mount child router to parent (thats what she said!)
parentRouter.Handle("/v1", childRouter)

// Register parent router (make it official)
server.Register(parentRouter)
```

### Structured API Implementation (For the Organized Code)

For those who like their code as organized as their sock drawer:

```go
type UserAPI struct {
    *Router
}

func (a *UserAPI) Register() {
    a.Router.Get("/", a.GetUsers)    // List all users (if any exist)
    a.Router.Post("/", a.CreateUser) // Create a user (good luck!)
}

func (a *UserAPI) GetUsers(w http.ResponseWriter, r *http.Request) (*ApiData, *ApiError) {
    // Handle get users (or pretend to)
    return nil, nil
}

func (a *UserAPI) CreateUser(w http.ResponseWriter, r *http.Request) (*ApiData, *ApiError) {
    // Handle create user (what could go wrong?)
    return nil, NewError("What could go wrong?") //a error would'nt hurt though
}

// Usage (it's easier than it looks)
server := New(":8000")
api := &UserAPI{
    Router: NewRouter("/users"),
}
server.Register(api)
```

## Route Handler Signature (The magic)

Route handlers use this signature (I tried to make it look professional):

```go
func(w http.ResponseWriter, r *http.Request) (*ApiData, *ApiError)
```

## Available Methods

It supports all your favorite HTTP methods (well, most of them):

- `Get(path string, handler HandlerFunc)` - For when you want to get stuff
- `Post(path string, handler HandlerFunc)` - For when you want to create stuff
- Too lazy to list the rest.

## Error Handling (Because Things Will Go Wrong)

This router uses custom error types, because regular errors weren't complicated enough:

```go
router.Get("/", func(w http.ResponseWriter, r *http.Request) (*ApiData, *ApiError) {
    if err := someOperation(); err != nil {
        return nil, &ApiError{
            Code:    http.StatusInternalServerError,
            Message: "Oops! Something went wrong (as usual)",
        }
    }
    return &ApiData{
        // your response data (assuming you have any)
    }, nil
})
```

### Middleware (Everyone has it,and so do I)

This router supports middleware for both individual routes and entire routers. Here's how to add some trust issues to your routes:

```go
// Define a middleware function
func AuthMiddleware(w http.ResponseWriter, r *http.Request) *ApiError {
    // Check something important (or not)
    if unauthorized := checkAuth(r); unauthorized {
        return NewError("Nice try!").SetCode(http.StatusUnauthorized)
    }
    return nil  // All good, proceed!
}

// Apply middleware to a single route
router.Post("/secure", func(w http.ResponseWriter, r *http.Request) (*ApiData, *ApiError) {
    return NewApiData("secure data!"), nil
}).Use(AuthMiddleware)

// Or apply middleware to an entire router (trust no one!)
router := NewRouter("/api").Use(AuthMiddleware)
```

Middleware Chaining (Because One Layer of Security Isn't Enough):

```go
router.Post("/fort-knox", handler).
    Use(AuthMiddleware,RateLimiter).
```

### Error Handling in Middleware

Middleware can return ApiError for clean error handling:

```go
func MockMiddleware(w http.ResponseWriter, r *http.Request) *ApiError {
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

## Testing (Yes, I Actually Tested My code)

It may pass sometimes, if not try running again.

```sh
go test -v ./...
```

## Todo(Anything else create a issue)
- [ ] File Server Handler for static files
- [ ] Templates handler(htmx baby...)
- [ ] CORS and cookie middleware(coz i use those a lot...)

## Contributing

Found a bug? Want to add a feature? Know how to actually write good code? I am all ears! Please feel free to submit a Pull Request. I promise to read it eventually.

## License

This software is released under the "It Works On My Machine" license. Use at your own risk!
