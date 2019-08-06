# Tree
A UNIX-like ```tree``` command. That produces a depth-indented listing of files. 
## Usage
1) ```$ go build -o ./bin/tree ./cmd/tree```
2) ```$ cd bin```
3) ```./tree <flags>  <directory name(current by default)>```
### Flags
```-f```: output files(optional)
## Run tests
```go test ./...```
