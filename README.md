# Tree
A UNIX-like ```tree``` command. Displays a directory hierarchy, with optional file names and sizes. The existing repository includes a tree representation, command-line rendering, fixtures, and tests. 

## Planned work
I plan to add directory statistics, JSON snapshots, snapshot comparison, filtering, and clearer handling of incomplete filesystem scans.
See [ROADMAP.md](ROADMAP.md). The main use case is explaining growth in a build or artifact directory from file metadata.

## Usage
1) ```$ go build -o ./bin/tree ./cmd/tree```
2) ```$ cd bin```
3) ```$ ./tree <flags>  <directory name(current by default)>```
s
### Flags
```-f```: output files(optional)

## Run tests
```$ go test ./...```


## Ownership and license
I created this personal project and own it. Its existing history will be preserved. License: MIT; see [LICENSE](LICENSE).
