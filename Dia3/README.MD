### Create a mod
```bash
> go mod init main.go
```

### Add refer internal
```go
import (
  _ "github.com/mattn/go-sqlite3"
)
```

### Access sqlite3
```shell
> sqlite3 [file-name].db
```

### Access sqlite3 in MEMORY
```shell
> sqlite3
```

### List tables
```shell
sqlite3> .table
```

### Create sample table
```sql
CREATE TABLE USERS (
  Id INTEGER PRIMARY KEY,
  Name TEXT,
  Email TEXT
);
```

### List users
```shell
sqlite3> select * from user;
```
