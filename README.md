# Unit Converter Terminal Client

A terminal-based TUI (Terminal User Interface) client for the [Unit Converter Server](https://github.com/DevSatyamCollab/unit-converter-server). Built with the [Charm](https://charm.sh/) ecosystem in Go, it lets you perform unit conversions entirely from your terminal — no browser required.

## Overview

This client connects to the Unit Converter backend server and provides an interactive, keyboard-driven interface to convert between units of:

- **Length** — millimeter, centimeter, meter, kilometer, inch, foot, yard, mile
- **Weight** — milligram, gram, kilogram, ounce, pound
- **Temperature** — Celsius, Fahrenheit, Kelvin

## Project Structure

```
.
├── LICENSE
├── go.mod
├── go.sum
├── main.go
└── internal/
    ├── api/
    │   └── converter.go
    └── tui/
        ├── commands.go
        ├── form.go
        ├── message.go
        ├── model.go
        ├── style.go
        ├── update.go
        └── view.go
```

## Tech Stack

| Library                                                 | Purpose                                        |
| ------------------------------------------------------- | ---------------------------------------------- |
| [Bubbletea](https://github.com/charmbracelet/bubbletea) | TUI framework (Elm-inspired model/update/view) |
| [Huh](https://github.com/charmbracelet/huh)             | Interactive terminal forms                     |
| [Lipgloss](https://github.com/charmbracelet/lipgloss)   | Terminal styling and layout                    |

## Prerequisites

- **Go 1.25+**
- The [Unit Converter Server](https://github.com/DevSatyamCollab/unit-converter-server) must be running before launching this client, as all conversions are processed through the server's API.

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/DevSatyamCollab/unit-converter-terminal-client.git
cd unit-converter-terminal-client
```

2. Start the backend server first (in a separate terminal window):

```bash
# In the unit-converter-server directory
go run .
```

3. Run the TUI client:

```bash
go run .
```

> **Note:** The terminal client relies entirely on the server's API to process conversions. Always ensure the Go server is running in a separate terminal window before launching the client.

## Controls

| Key           | Action                                 |
| ------------- | -------------------------------------- |
| `Esc`         | Reset / go back to the main menu       |
| `Ctrl + L`    | Navigate right                         |
| `Ctrl + H`    | Navigate left                          |
| `Tab`         | Move forward between fields in a form  |
| `Shift + Tab` | Move backward between fields in a form |

## How It Works

1. Launch the client — a menu lets you choose a conversion category (Length, Weight, or Temperature).
2. Fill in the form — enter your value and select the source and target units using `Tab` / `Shift+Tab` to move between fields.
3. Submit — the client sends the request to the running server and displays the converted result inline.
4. Press `Esc` at any point to reset and return to the main menu.

## License

This project is licensed under the [MIT License](LICENSE).
