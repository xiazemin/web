#!/usr/local/opt/php@7.0/bin/php
<?php
print "Content-Type: text/html\r\n";
print "X-CGI-Pid: $$\r\n";
print "X-Test-Header: X-Test-Value\r\n";
print "\r\n";
phpinfo();