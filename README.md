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