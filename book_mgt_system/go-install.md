# Install Go (Linux)

## Download and extract

```bash
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
```

## add Go to PATH

```bash
export PATH=$PATH:/usr/local/go/bin
```

### check version
```bash
go version
```
