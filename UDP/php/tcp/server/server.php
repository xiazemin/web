<?php
$host = '127.0.0.1';
$port = '81';
$callback = 'echo';

function receive_tcp_message($host, $port, $callback)
{
    $socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);

    // socket_bind() 的参数 $host 必传, 由于是监听本机, 此处可以固定写本机地址
    // 注意: 监听本机地址和内网地址效果不一样
    @socket_bind($socket, $host, $port);
    @set_time_limit(0);

    // 绑定端口之后调用监听函数, 实现端口监听
    @socket_listen($socket, 5);

    // 接下来只需要一直读取, 检查是否有来源连接即可, 如果有, 则会得到一个新的 socket 资源
    while ($child = @socket_accept($socket))
    {
        // 休息 1 ms, 也可以不用休息
        usleep(1000);

        if (false === socket_getpeername($child, $remote_host, $remote_port))
        {
            @socket_close($child);
            continue;
        }

        // 读取请求数据
        // 例如是 http 报文, 则解析 http 报文
        $request = '';
        do
        {
            $buffer = @socket_read($child, 1024, PHP_BINARY_READ);
            if (false === $buffer)
            {
                @socket_close($child);
                continue 2;
            }
            $request .= $buffer;
        } while (strlen($buffer) == 1024);

        // 此处省略如何调用 $callback
        $response = $callback($remote_host, $remote_port, $request);

        if (!strlen($response))
        {
            // 至少返回含有一个空格的字符串
            $response = ' ';
        }

        // 因为是 TCP 链接, 需要返回给客户端处理数据
        $num = 0;
        $length = strlen($response);
        do
        {
            $buffer = substr($response, $num);
            $ret = @socket_write($child, $buffer);
            $num += $ret;
        } while ($num < $length);

        // 关闭 socket 资源, 继续循环
        @socket_close($child);
    }
}

// 客户端来的任何请求都会打印到屏幕上
receive_tcp_message($host, $port, $callback);
// 如果程序没有出现异常，该进程会一直存在
