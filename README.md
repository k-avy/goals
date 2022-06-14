# gols CLI Tool

A CLI tool in go language that works as ll command.

## Prerequisite

Please [install and set-up Golang](https://go.dev/doc/install) on your system in advance.

## How to run this project?

1. Clone this Project and Navigate to the folder.

```bash
git clone https://github.com/k-avy/goals.git
cd goals
```

2. Build the project using following command.

```bash
go build ./cmd/gol
```

3. Run the executable in your vscode terminal.

```bash
./gol
```

4. You can also directly install this in your system.

```bash
 go install ./cmd/gols
```

## Features

1. To get path of current working directory we can use following command:

```bash
gols current
```

2. To get all the files and folders of given path we can use following command.

```bash
gols all -p "path"
```

3. To get all the files of given path we can use following command.

```bash
gols file -p "path"
```

4. To get all the folders of given path we can use following command:

 ```bash
gols folder -p "path"
```

5. You can also take help from the following command:

 ```bash
gols help
```

