# Anginx

## Analyze nginx log
## Install
```
go get github.com/terryshi96/anginx
```
or download binary [release](https://github.com/terryshi96/anginx/releases)
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
