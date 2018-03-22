One host

```
connection-checker -hosts=ya.ru:80
connection-checker -hosts=87.250.250.242:80
```

Multiple hosts

```
connection-checker -hosts=87.250.250.242:80,google.com:80
```

TLS

```
connection-checker -hosts=ya.ru:443,google.com:443 -tls
```
