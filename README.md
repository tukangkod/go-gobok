# Go-Gobok

Gobok is a malay word shortformed of Gerobok which translate as Cupboard/Shelve/Closet.

## what is this project about?
A test project to use following:
- gRPC on go (https://github.com/grpc/grpc-go)
- Go Postgres driver (https://github.com/go-pg/pg)
- configuration (https://github.com/spf13/viper)
- loging (https://github.com/uber-go/zap)

We have two gRPC service:
- Put - save client ip, server ip, tags and messages
- Search - search by client ip, server ip, tags and return all matched records

## how to run?
1. configure config.yaml
    ```cmd
    $ cp config.yaml.example config.yaml
    $ vi config.yaml
    ```
    Default setting will be:
    - user: postgres
    - database: postgres

2. run the server.
    ```cmd
    $ go run server/*.go
    ```

3. run put client to save data
    ```cmd
    $ go run client/put.go "127.0.0.1" "" "tag1=value" "testing go-gobok"
    ```

4. run search client to search data
    ```cmd
    $ go run client/search.go "" "" "tag1=value"
    ```

## note:
- test and benchmark only for utils at the moment
