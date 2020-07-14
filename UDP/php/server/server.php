<?php

$host = '127.0.0.1';
$port = '8082';
$callback = 'var_dump';

function receive_udp_message($host, $port, $callback)
{
    $socket = socket_create(AF_INET, SOCK_DGRAM, SOL_UDP);

    $r1=@socket_bind($socket, $host, $port);
    $r2=@set_time_limit(0);
var_dump([$r1,$r2]);

    while (true)
    {
        usleep(100);

        $ret = @socket_recvfrom($socket, $request, 16384, 0, $remote_host, $remote_port);
        var_dump($ret);
        if ($ret)
        {
            $callback([$remote_host, $remote_port, $request]);
        }

        // 不需要返回给客户端任何消息, 继续循环
    }
}

// 客户端来的任何请求都会打印到屏幕上
receive_udp_message($host, $port, $callback);
// 如果程序没有出现异常，该进程会一直存在
