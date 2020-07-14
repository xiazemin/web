Navigated to http://127.0.0.1:8000/
(index):1 Access to XMLHttpRequest at 'http://127.0.0.1:8001/ajax' from origin 'http://127.0.0.1:8000' has been blocked by CORS policy: No 'Access-Control-Allow-Origin' header is present on the requested resource.
(index):23 Cross-Origin Read Blocking (CORB) blocked cross-origin response http://127.0.0.1:8001/ajax with MIME type application/json. See https://www.chromestatus.com/feature/5629709824032768 for more details.
loadXMLDoc @ (index):23
onclick @ (index):30
Navigated to http://127.0.0.1:8000/
(index):1 Access to XMLHttpRequest at 'http://127.0.0.1:8001/ajax' from origin 'http://127.0.0.1:8000' has been blocked by CORS policy: The 'Access-Control-Allow-Origin' header has a value 'http://localhost:8000' that is not equal to the supplied origin.
(index):23 Cross-Origin Read Blocking (CORB) blocked cross-origin response http://127.0.0.1:8001/ajax with MIME type application/json. See https://www.chromestatus.com/feature/5629709824032768 for more details.
loadXMLDoc @ (index):23
onclick @ (index):30
Navigated to http://127.0.0.1:8000/


$ curl -i http://127.0.0.1:8001/ajax
HTTP/1.1 200 OK
Access-Control-Allow-Origin: http://127.0.0.1:8000
Content-Type: application/json; charset=utf-8
Date: Sun, 22 Sep 2019 14:21:49 GMT
Content-Length: 38

{"name":"benben_2015","msg":"success"}