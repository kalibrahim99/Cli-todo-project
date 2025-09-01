# CLI Todo Project

A simple command-line todo list manager written in Go.

## Features

- Add tasks to your todo list
- Tasks are persisted in a local `tasks.json` file

## Getting Started

### Prerequisites

- Go 1.18+ installed
- Git (optional, for cloning)

### Installation

Clone the repository:

```sh
git clone <your-repo-url>
cd Cli-todo-project
```

Build the project:

```sh
go build -o todo
```

### Usage
if you write a task use "-" after word
Add a new task:

```sh
./todo add "Your-task-description-here"
```

Example:

```sh
./todo add "Buy-groceries"
```

## Project Structure

```
.
├── main.go
├── tasks/
│   └── todolist.go
├── tasks.json
└── README.md
```

## Contributing

Contributions are welcome! Please open issues or submit pull requests for improvements and bug fixes.

## License

This project is licensed under the MIT License.
