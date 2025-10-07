# Setup Guide: Go with Visual Studio Code

This guide explains how to install **Go**, set up **Visual Studio Code (VS Code)**, and configure the **Go extension** on **Windows, macOS, and Linux**.

---

## 1. Install Go

### Windows
1. Download the Windows installer from [https://go.dev/dl/](https://go.dev/dl/).
2. Run the `.msi` installer and follow the setup wizard.
3. By default, Go will be installed in `C:\Go` and the installer will update your `PATH` automatically.

### macOS
1. Download the `.pkg` installer for macOS from [https://go.dev/dl/](https://go.dev/dl/).
2. Run the installer and follow the instructions.
3. Alternatively, you can install via **Homebrew**:
   ```bash
   brew install go
   ```

### Linux
1. Download the appropriate `.tar.gz` archive from [https://go.dev/dl/](https://go.dev/dl/).
2. Extract it to `/usr/local` (requires sudo):
   ```bash
   sudo tar -C /usr/local -xzf goX.X.X.linux-amd64.tar.gz
   ```
3. Add Go to your PATH by editing your shell config (`~/.bashrc`, `~/.zshrc`, etc.):
   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```
4. Reload your shell:
   ```bash
   source ~/.bashrc
   ```

---

## 2. Verify Go Installation

Run the following command in your terminal (Command Prompt/PowerShell on Windows, Terminal on macOS/Linux):

```bash
go version
```

Expected output (example):
```
go version go1.23.0 windows/amd64
```

Check your environment setup:

```bash
go env
```

This shows details like `GOPATH`, `GOROOT`, and module configuration.

---

## 3. Install Visual Studio Code

### Windows / macOS / Linux
1. Download VS Code from [https://code.visualstudio.com/](https://code.visualstudio.com/).
2. Install it by running the installer (Windows/macOS) or by extracting/installing the package (Linux).

#### Linux package managers:
- Ubuntu/Debian:
  ```bash
  sudo apt update
  sudo apt install code
  ```
- Fedora/RHEL:
  ```bash
  sudo dnf install code
  ```
- Arch Linux:
  ```bash
  sudo pacman -S code
  ```

---

## 4. Install the Go Extension in VS Code

1. Open **Visual Studio Code**.
2. Go to **Extensions** (`Ctrl+Shift+X` on Windows/Linux, `Cmd+Shift+X` on macOS).
3. Search for **Go** and install the extension published by the **Go Team at Google**.
4. After installation, VS Code may prompt you to install extra tools like:
   - `gopls` (language server)
   - `dlv` (debugger)
   - formatting/linting tools  
   Confirm and install them for full functionality.

---

## 5. Verify the Setup in VS Code

1. In your existing project folder, use the provided [hello_world.go](./hello_world/hello_world.go) file or create it if not already present:

   ```go
   package main

   import "fmt"

   func main() {
       fmt.Println("hello world")
   }
   ```

2. Run it in the integrated terminal:

   ```bash
   go run hello_world.go
   ```

   Output should be:
   ```
   Hello, World!
   ```

---

✅ Your Go development environment is now fully set up on **Windows, macOS, or Linux**!
