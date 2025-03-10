# gopm – A Simple gRPC-Based Process Manager in Go

**gopm** provides a command-line interface (CLI) to manage local processes via a Go-based gRPC server. You can run the server in the foreground or background, then issue commands like `start`, `stop`, `list`, `log`, and `remove` to control processes remotely.

---

## Installation

1. **Clone** this repository:
   ```bash
   git clone https://github.com/<your-username>/gopm.git
   cd gopm

2. Build the CLI:
   ```bash
   go build -o gopm ./cmd
   sudo mv gopm /usr/local/bin


## Usage

gopm <command> [flags] [arguments...]

Available Commands:

**init**  
Runs the gRPC server in the foreground (blocking the terminal). Example: 

`gopm init`

**init-bg**  
Spawns the gRPC server in a background process, returning control to the shell immediately. Example:  
`gopm init-bg`

**start <name> <command> [args...]**  
Starts a named process using the specified command and optional arguments. Example:  
`gopm start myapp python3 myscript.py`

**stop <name>**  
Stops a running process by name. Optional flag: --force (for immediate kill). Example:  
`gopm stop myapp`

**list**  
Lists all tracked processes. Optional flag: --verbose (for more info). Example:  
`gopm list`

**log <name>**  
Streams log output of a process. Optional flag: --follow (for real-time logs). Example:  
`gopm log myapp`

**remove <name>**  
Removes a process record from the manager. Depending on your setup, you may need to stop it first. Example:  
`gopm remove myapp`

Examples:

1) Start the server in the foreground, then start and stop a process:
   - Terminal 1: `gopm init`
   - Terminal 2: `gopm start myapp python3 myscript.py, gopm list, gopm stop myapp, gopm remove myapp`

2) Start the server in the background, manage a process, and tail logs:
   - `gopm init-bg`
   - `gopm start worker python3 worker.py`
   - `gopm log worker --follow`
   - `gopm stop worker`
