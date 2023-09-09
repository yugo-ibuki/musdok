# Overview

This is a just simple TODO apps used through CLI.

## Usage

### Install
```bash
go install github.com/yugo-ibuki/musdok
```

### Create DB
```bash
musdok init
```

### Check ALL TODOS
```bash
musdok all
```

### Add TODO
```bash
musdok create --title "title" --desc "description"
```

### Update TODO
```bash
musdok update --id 1 --title "title" --desc "description"
```

### Delete TODO
```bash
musdok delete --id 1
```

## Options

### Root Action
```bash
$ musdok -h
musdok is a CLI to add and search for the todo list

Usage:
  musdok [flags]
  musdok [command]

Available Commands:
  all         shows all todos
  completion  Generate the autocompletion script for the specified shell
  create      Create a new todo
  delete      Delete a todo
  help        Help about any command
  init        Initialize the database
  update      Update a new todo

Flags:
  -h, --help   help for musdok
```

### init
```bash
musdok init -h                                                                                                                    土  9/ 9 16:48:13 2023
Initialize the database

Usage:
  musdok init [flags]

Flags:
  -h, --help   help for init
```

### all
```bash
musdok all -h                                                                                                                     土  9/ 9 16:47:59 2023
shows all todos

Usage:
  musdok all [flags]

Flags:
  -h, --help   help for all
```

### create
```bash
$ musdok create -h                                                                                                          437ms  土  9/ 9 16:46:10 2023
Create a new todo

Usage:
  musdok create [flags]

Flags:
      --desc string    TODO description
  -h, --help           help for create
      --title string   TODO title
```

### update
```bash
musdok update -h                                                                                                                  土  9/ 9 16:47:23 2023
Update a new todo

Usage:
  musdok update [flags]

Flags:
      --desc string    TODO description
  -h, --help           help for update
      --id string      todo id
      --title string   TODO title
```

### delete
```bash
 musdok delete -h                                                                                                                  土  9/ 9 16:46:19 2023
Delete a todo

Usage:
  musdok delete [flags]

Flags:
  -h, --help        help for delete
      --id string   TODO id
```
