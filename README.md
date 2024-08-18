# toho 

Stands for `To` `H`eader `O`nly

A cli tool to convert a multi-file c library to a header only one

> [!WARNING]
> This tool currently works for the specific way
> I write C libraries. If you are using other conventions
> Some things might not work correctly. 
> Use at your own risk

## Building 

```bash
go build ./cmd/toho-cli
```

## Usage

```bash
$ toho <project-path> <filename> <library-define>
```

## LICENSE

[MIT](./LICENSE)
