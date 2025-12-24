# Command Notes

## Ch1 cmd notes

### Running Go Programs
```bash
# Run a Go program
go run ./ch1/helloworld/main.go

# Run a server in background
go run ./ch1/server1/main.go &

# Run with output redirection
go run ./ch1/lissajous/main.go > out.gif
```

### Building Go Programs
```bash
# Build a program (saves binary in current directory)
go build ./ch1/helloworld

# Build with custom output location
go build -o ./ch1/helloworld/helloworld ./ch1/helloworld/main.go

# Install a program (saves to $GOPATH/bin or $GOBIN)
go install ./ch1/helloworld
```

### Process Management
```bash
# Find process on port 8000
lsof -i :8000

# Kill process on port 8000 (with verbose output)
lsof -ti :8000 | xargs -t kill -9

# List background jobs
jobs

# Kill background job by number
kill %1

# Kill by PID
kill <PID>
```

### HTTP Server Testing
```bash
# Fetch from local server
go run ./ch1/fetch/main.go http://localhost:8000

# Fetch with curl
curl http://localhost:8000

# Open in browser (macOS)
open http://localhost:8000
```

### Package Management
```bash
# Check module info
go mod tidy

# View module dependencies
go list -m all
```

