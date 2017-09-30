# Anginx

## Analyze nginx log
## Install
```
go get github.com/terryshi96/anginx
```
## Usage
```
anginx -c config_file_path
```

## Config file
```
log must use '|' to split
for example '$remote_addr|$remote_port|$request_time|$remote_user|$time_local|$request|$status|$body_bytes_sent|$http_referer|$http_user_agent'
```

## Based on sqlite3
