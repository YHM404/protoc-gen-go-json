# Test

``` shell
protoc --go_out=. --go_opt=paths=source_relative \
       --go-json_out=. --go-json_opt=paths=source_relative,generate_binary=true \
       proto2_test.proto

protoc --go_out=. --go_opt=paths=source_relative \
       --go-json_out=. --go-json_opt=paths=source_relative,generate_marshal=true \
       proto2_test.proto

protoc --go_out=. --go_opt=paths=source_relative \
       --go-json_out=. --go-json_opt=paths=source_relative,generate_unmarshal=true \
       proto2_test.proto
```