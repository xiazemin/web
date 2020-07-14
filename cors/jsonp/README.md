$ curl -i http://127.0.0.1:8002/jsonp?callback=myCallback
HTTP/1.1 200 OK
Content-Type: application/javascript
Date: Sun, 22 Sep 2019 14:48:16 GMT
Content-Length: 60

myCallback({"Accept":["*/*"],"User-Agent":["curl/7.49.0"]});