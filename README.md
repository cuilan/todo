# Todo

To-do cli application with golang.

Help you record to-do items and work more efficiently.

# Table of Contents

- [Overview](#overview)
- [Installing](#installing)
- [Configuration](#configuration)
- [Getting Started](#getting-started)
  * [Version](#version)
  * [Todo items list](#todo-items-list)
  * [Add todo item](#add-todo-item)
  * [Done item](#done-item)
- [License](#license)

# Overview

Use Cobra library.

```shell
todo version
```

# Installing

Using todo is easy. First, use `go get` to install latest version of the library.
This command will install the `todo` generator executable along with the library and its dependencies:

```go
go get -u github.com/cuilan/todo
```

Next, build it:

```go
go build -o full/path/todo
```

# Configuration

First, Create a `config.yml` configuration file in the `$HOME` directory.

Next, Only need to add a configuration for configuring the data file storage location:

```yml
datafile: C:\Users\xxxxx\.todo.json
```

Create your json data file.

# Getting Started

## Version

Show current todo version:

```shell
todo version
```

## Todo items list

Show undone todo items list:

```shell
todo list
```

Show all todo itesm list:

```shell
todo list -a
```

Show done todo itesm list:

```shell
todo list -d
```

## Add todo item

Add item to list, priority:(1, 2, 3), default: 3:

```shell
todo add -p 1 xxxx
```

## Done item

Done item with index:

```shell
todo done 1
```

# License

Todo is released ubder the Apache 2.0 license.See [LICENSE](https://github.com/cuilan/todo/blob/master/LICENSE)
