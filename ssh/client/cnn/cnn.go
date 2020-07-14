package cnn
func SSHConnect( user, password, host string, port int ) ( *ssh.Session, error ) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	clientConfig = &ssh.ClientConfig{
		User:               user,
		Auth:               auth,
		// Timeout:             30 * time.Second,
		HostKeyCallback:    hostKeyCallbk,
	}

	// connet to ssh
	addr = fmt.Sprintf( "%s:%d", host, port )

	if client, err = ssh.Dial( "tcp", addr, clientConfig ); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

func runSsh(){

	var stdOut, stdErr bytes.Buffer

	session, err := SSHConnect( "username", "passworld", "192.168.1.4", 22 )
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.Stdout = &stdOut
	session.Stderr = &stdErr

	session.Run("if [ -d liujx/project ]; then echo 0; else echo 1; fi")
	ret, err := strconv.Atoi( str.Replace( stdOut.String(), "\n", "", -1 )  )
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d, %s\n", ret, stdErr.String() )

}


//远程登录服务器
//id：任务id，name：任务名称，command：命令详情，server：服务器信息
func RemoteCommandJobByPassword(id int, name string, command string, servers *models.TaskServer) *Job {
	var client *ssh.Client
	var err error
	var session *ssh.Session
	job := &Job{
		id: id,
		name: name,
	}
	job.runFunc = func(timeout time.Duration) (string, string, error, bool) {
		auth := make([]ssh.AuthMethod, 0)
		auth = append(auth, ssh.Password(servers.Password))
		clientConfig := &ssh.ClientConfig{
			User: servers.ServerAccount,
			Auth: auth,
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		}
		//servers.ServerIp：内网ip, servers.Port：服务器端口号
		addr := fmt.Sprintf("%s:%d", servers.ServerIp, servers.Port)
		//通过ssh连接服务器
		if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
			return "", "", err, false
		}
		defer client.Close()
		//建立会话
		if session, err = client.NewSession(); err != nil {
			return "", "", err, false
		}
		bufOut := new(bytes.Buffer)
		bufErr := new(bytes.Buffer)
		session.Stdout = bufOut
		session.Stderr = bufErr
		//执行命令
		if err = session.Run(command); err != nil {
			return "", "", err, false
		}
		isTimeout := false
		return bufOut.String(), bufErr.String(), err, isTimeout
	}
	return job
}