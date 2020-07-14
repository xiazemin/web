<?php
$host = '127.0.0.1';
$port = '81';
$message = 'Hello TCP Server';

function send_tcp_message($host, $port, $message)
{
    $socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
    @socket_connect($socket, $host, $port);

    $num = 0;
    $length = strlen($message);
    do
    {
        $buffer = substr($message, $num);
        $ret = @socket_write($socket, $buffer);
        $num += $ret;
    } while ($num < $length);

    $ret = '';
    do
    {
        $buffer = @socket_read($socket, 1024, PHP_BINARY_READ);
        $ret .= $buffer;
    } while (strlen($buffer) == 1024);

    socket_close($socket);

    return $ret;
}

$ret = send_tcp_message($host, $port, $message);
