#

## create folder
```
mkdir go-sample
```

## create go.mod
```
go mod init go-sample
```

## run main.go
```
go run main.go
```

## build & run

◆Local起動用
```
go build -o ./release/bin/app main.go
```
-o : BinaryFileName

◆Linux起動用
```
GOOS=linux GOARCH=amd64 go build -o ./release/bin/app main.go
```

◆Run
```
./app
```

## external package install
```
go mod tidy
```


## Official Document

go-cs-opensource
```
https://cs.opensource.google/go/go
```

go-staticcheck(linter)
```
error list
https://staticcheck.dev/docs/checks/
```

go dot env 
```
github
https://github.com/joho/godotenv/blob/main/godotenv.go
```