# api-differ 
CLI for validating delta comparisons of API responses

## Usage Manual

```shell
NAME:
   CLI for validating delta comparisons of API responses-o -o  

USAGE:
   main [global options] command [command options] [arguments...]

DESCRIPTION:
   Compare API response differences

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -c value    config file path
   -o value    dist report path
   --help, -h  show help (default: false)
```

### -c (config)
The path of the JSON file where you wrote the test scenario.  
Please refer to the following sample for JSON

```shell
{
  "actual": {
    "url": "http://sample-users.old.xxx.xx",
    "header": {
      "User-Agent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"
    },
    "cookie": "hoge=fuga; fuga=hoge"
  },
  "expect": {
    "url": "http://sample-users.new.xxx.xx",
    "header": {
      "User-Agent": "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"
    },
    "cookie": "hoge=fuga; fuga=hoge"
  },
  "scenario": {
    "method": "GET",
    "type": "query",
    "params": [
      "id=1",
      "id=2"
    ]
  }
}
```

### -o (out)

The path to output the result.  
Please specify with xxx.md because it is fixed to .md.

## Develop
test run

```shell
go run main.go -c sample/config.json ./dist.md
```

build 

```shell
api-diff-tester % go build -o api-differ main.go 
```
