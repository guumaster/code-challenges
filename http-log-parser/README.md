# log-monitor - companyx take-home assignment

This folder contains the source code and a binary version for take-home assignment project.


## Usage

You can invoke the binary from a console with the following commands:

- With file as parameters
```
$> ./bin/companyx-log-monitor --from some_path/some_log_file.csv
```

- With file piped to the CLI program:
```
$> cat <any_path>/some_log_file.csv | ./bin/companyx-log-monitor 
```
Alternatively with this command: 
```
$> ./bin/companyx-log-monitor  < some_path/some_log_file.csv 
```

## Development

----------------------------

### Run 
You can run the CLI program with the following command:

```
src/$> go run cmd/log-monitor/*.go --from test/data/sample_log.csv
```


### Tests

You can run tests with the following command: 

```
src/$> go test -cover ./...
```

### Lint

You can check the code with `golint` tool:

```
src/$> golint ./...
```



## Developer Notes

See [DEV Notes file](./DEV_NOTES.md)
