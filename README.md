# Bare Go Server 

This repo is inspired by this [post](https://medium.com/code-zen/why-i-don-t-use-go-web-frameworks-1087e1facfa4). Basically I'm going to create a full-fledged go web server without the use of any existing frameworks (echo, gin, revel, iris, etc.). I will be using the [gorilla/mux](http://www.gorillatoolkit.org/pkg/mux) package, because it is not a framework, just a collection of packages that helps extend the net/http package; additionally, I'm using [negroni](https://github.com/urfave/negroni) for middleware and [bolt](https://github.com/boltdb/bolt) for storage.

This repo is also a part of a docker related tutorial that I'm currently writing. (I will post the actual article when it's done)


## What this server is going to do

- This server should be able to create a user
- This server should be able to fetch a list of users

## Todos

**Priority**

- [x] Create a User Model
- [x] Create a User route handler
- [x] Setup Bolt DB to store users
- [ ] Write tests for each of the routes
