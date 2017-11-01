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

 - inputfile: log file you want to parse
 - startdate: must be like /23\/Oct\/2017/
 - enddate:
 - overtime: request over this time will be recorded(default 1) 
 - topip: record top x ip addresses(default 10)
 - toprequest: record top x requests(default 10)
 - logformat: '$remote_addr |$remote_port |$request_time |$remote_user |$time_local |$request |$status |$body_bytes_sent |$http_referer |$http_user_agent'
 > log must be splited by ' |' 
 - truncatedatabase: if you want to load data(default false)
 - emailconfig: email support config

## Based on sqlite3
