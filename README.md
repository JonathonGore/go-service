# go-service

go-service is a base repository to quickly develop micro-services using Go. 
It provides functionality and structure common to Go micro-services. 

## Getting Started

1) Clone the repository to a new folder within your go path

```
$ pwd
/Users/jack/go/src/github.com/JonathonGore
$ git clone --depth 1 git@github.com:JonathonGore/go-service.git new-service
```

2) Update git to point to the GitHub repository you want to use
```
git remote set-url origin https://github.com/USERNAME/REPOSITORY.git
```

3) Update import paths
```
./scripts/setup.sh 'github.com/USERNAME/REPOSITORY'
```
4) Get dependencies
```
dep ensure
```
5) Run the project
```
 go run main.go
```
