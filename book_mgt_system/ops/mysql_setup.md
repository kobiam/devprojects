# MySQL Setup

## create admin user

```sql
CREATE USER 'admin'@'%' IDENTIFIED BY '123';
```

## create database

```sql
CREATE DATABASE testdb;
```

## set permissions to admin user

```sql
GRANT ALL PRIVILEGES ON testdb TO 'admin'@'%';
```

## show privileges

```sql
SHOW GRANTS FOR 'admin'@'%';
```