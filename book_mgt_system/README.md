# Book Management System

## Teckstack

**Go** and **MySQL**

### Setup

```bash
go mod init github.com/kobiam/devprojects
```

### Download packages

```bash
go get "github.com/jinzhu/gorm"
go get "github.com/jinzhu/gorm/dialects/mysql"
go get "github.com/gorilla/mux"
go get "github.com/kobiam/devprojects/book_mgt_system/pkg/routes"
go get "github.com/kobiam/devprojects/book_mgt_system/pkg/controllers"
```

### Build

```bash
cd /cmd/main
go build
```
