# nginx-reverse-porxy

## Nginx usage

```shell
nginx # start the nginx server

nginx -s quit # stop the nginx server

nginx -s stop # stop immediately

nginx -s reopen # reopen the log file

```

## Config nginx for reverse proxy (for load balance)

```shell
nginx -t # find the path of config file, then edit it
```

```nginx.conf
http {
  // define the upstream server in http block
  upstream app_server {
    ip_hash; // the same ip will visit same server
    server 127.0.0.1:8000 weight=3; // the weight of server
    server 127.0.0.1:8001;
    server 127.0.0.1:8002;
  }

  server {
    location /app { // the location of the reverse proxy
      proxy_pass http:// app_server; // same name as the upstream server
    }
  }
}
```

```shell
nginx -s reload # reload the config file after edit
```
