# Basics command lines for Go

### Importing Go package 
`
go get -u [url]
`
<br/><br/>
### Set offline Go Docs
`
godoc --http=:6060
`
<br/><br/>
### Checking your Go Env and Path
`
go env || go path
`
<br/><br/>
### Checking not used code
`
go vet gofile.go
`
<br/><br/>
### Building a Go Program
`
go build gofile.go
`
<br/><br/>
### Executing a Go File (Accepts * for recursive)
`
go run gofile.go
`