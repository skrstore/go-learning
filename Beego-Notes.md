# Learn Beego

## Important Commands

```sh
# To get beego and bee
go get -u github.com/beego/beego/v2
go get -u github.com/beego/bee/v2

# to create a new project using bee
bee new NAME
bee new . # to create in current directory

# to run the project
cd NAME
bee run

# to run the project and create docs with it
bee run -downdoc=true -gendoc=true
```

## References

- https://beego.me/quickstart
- https://beego.me/docs/intro/
- [Golang Beego Tutorial: Quick Start with Beego & CLI (Tutorial 1) - CoderVlogger - YouTube](https://youtu.be/-WqrOptT1Tw)

## Content

- [x] Simple Server
- [x] Static file server
- [x] Template parsing
- [x] Getting data from form(GET, POST), Query params, Parameterized URL and single, multiple files
- [ ] How to save data to DB using concept of ORM
- [ ] How to store, read data in DB with relations

## Simple Server

- [Example Code](./projects/simple_server/main.go)
- Important Commands

```sh
mkdir simple_server
go mod init simple_server
touch main.go
go mod vendor
go run main.go
```

## Serve Static Files

- [Example Code](projects/serve_static_files/main.go)
- Reference - https://beego.me/docs/mvc/view/static.md

## Template parsing

- [Example Code](projects/template_parsing/main.go)
- Reference - https://beego.me/docs/mvc/view/view.md

### Content

- [x] Reading variable in Template file
- [x] Rendering Forms
- [ ] Using Template Layout
- [ ] Using Layout Selection

## Getting Data

- [Example Code](projects/getting_data/main.go)
- Reference - https://beego.me/docs/mvc/controller/params.md

### Content

- [x] From Query parameter
- [x] From Parameterized URL
- [x] From HTML Form using GET
- [x] From HTML Form using POST
- [ ] Parse to struct(with form Parsing)
- [ ] Automatic Parameter Routing
- [ ] Retrieving data from request body
- [ ] Uploading single file
- [ ] Uploading multiple files
