$ telnet 127.0.0.1 8877
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
^C
Connection closed by foreign host.




$ curl http://127.0.0.1:8877/hijack -H 'Connection:Keep-alive;Keep-Alive:timeout=20'
Now we're speaking raw TCP. Say hi: localhost:hijack


after sleep

$ curl http://127.0.0.1:8877/hijack -H 'Connection:Keep-alive;Keep-Alive:timeout=20'
curl: (52) Empty reply from server
