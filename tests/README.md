# Functional Tests

In this test, latest `gcli` binary is executed and check the generated codes. The followings are tested,

- The codes are `go build`-able
- The codes pass `go test` & `go vet`
- Create binary from the generated codes and check outputs

## Usage

To run test, 

```bash
$ cd $GOPATH/src/github.com/tcnksm/gcli
$ make  test-functional
```
