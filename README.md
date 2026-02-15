
# GoAdmin Instruction

GoAdmin is a golang framework help gopher quickly build a data visualization platform. 

- [github](https://github.com/GoAdminGroup/go-admin)
- [forum](http://discuss.go-admin.com)
- [document](https://book.go-admin.cn)

## Directories Introduction

```
.
├── Makefile            Makefile
├── adm.ini             adm config
├── build               binary build target folder
├── config.yml          config file
├── go.mod              go.mod
├── go.sum              go.sum
├── html                frontend html files
├── logs                logs
├── main.go             program entrance file
├── main_test.go        test file
├── pages               page controllers
├── tables              table models
└── uploads             upload directory
```

## Generate Table Model

### online tool

visit: http://127.0.0.1:8081/admin/info/generate/new

### use adm

```
adm generate
```

### tools

Generate project template success~~🍺🍺

1. Import and initialize database:

- sqlite: https://github.com/GoAdminGroup/go-admin/raw/master/data/admin.db
- mssql: https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.mssql
- postgresql: https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.pgsql
- mysql: https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.sql

1. Execute the following command to run:

> make init module=app
> make install
> make serve

1. Visit and login:

- Login: http://127.0.0.1:8081/admin/login
account: admin  password: admin

- Generate CRUD models: http://127.0.0.1:8081/admin/info/generate/new

1. See more in README.md

see the docs: https://book.go-admin.com
visit forum: http://discuss.go-admin.com
