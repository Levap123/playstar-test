# playstar-test

## Run locally
```bash
git clone https://github.com/Levap123/playstar-test
cd playstar-test
make build && make run
```

## Example 

![alt text](https://github.com/Levap123/playstar-test/blob/main/img/request-example.jpg)


## get logs from db

```bash
docker exec -ti  playstar-test_db_1 psql -d logs -U root
```

```bash
select * from logs;
```