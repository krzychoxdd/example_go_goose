# example_go_goose
Example of usage golang goose migration library for real project

Step 1
``
git clone https://github.com/krzychoxdd/example_go_goose.git
``

Step 2

``make dbinstall``

Step 3

``make dbstart``

Step 4

``make dbcreate``

Step 5

``make migrate``

After all you shoud see something like this

```
cd cmd && go run migrator.go root:1234567@tcp\(0.0.0.0:6603\)/exmpldb?parseTime=true up
2023/01/26 15:16:37 OK   20230125232556_create_schema.sql (424.66ms)
2023/01/26 15:16:40 OK   20230126001631_seed_schema.go (17ms)
2023/01/26 15:16:40 goose: no migrations to run. current version: 20230126001631
```
