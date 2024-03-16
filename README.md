# CDN E-Masjid.My

1.
```
docker build -t cdn-emasjidmy .
```

2.
```
docker run -d -p 80:80 -v /Users/rohaizan/Codes/microk8s/emasjid/cdn.e-masjid.my/repository/:/usr/share/nginx/html/images/ cdn-emasjidmy
```