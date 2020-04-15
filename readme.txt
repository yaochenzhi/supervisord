# http://supervisord.org/configuration.html

Ⅰ   config file:

    $CWD/supervisord.conf
    $CWD/etc/supervisord.conf
    /etc/supervisord.conf
    /etc/supervisor/supervisord.conf (since Supervisor 3.3.0)
    ../etc/supervisord.conf (Relative to the executable)
    ../supervisord.conf (Relative to the executable)

Ⅱ   config file format:
    
    Windows-INI-style (Python ConfigParser) file

Ⅲ   config file env var:
    
    %(ENV_X)s       # env | grep LOGLEVEL

    [program:example]
    command=/usr/bin/example --loglevel=%(ENV_LOGLEVEL)s

Ⅳ   config file section sample:

    [program:cat]

        command=/bin/cat
        process_name=%(program_name)s
        numprocs=1
        directory=/tmp
        umask=022
        priority=999
        autostart=true
        autorestart=unexpected
        startsecs=10
        startretries=3
        exitcodes=0
        stopsignal=TERM
        stopwaitsecs=10
        stopasgroup=false
        killasgroup=false
        user=chrism
        redirect_stderr=false
        stdout_logfile=/a/path
        stdout_logfile_maxbytes=1MB
        stdout_logfile_backups=10
        stdout_capture_maxbytes=1MB
        stdout_events_enabled=false
        stderr_logfile=/a/path
        stderr_logfile_maxbytes=1MB
        stderr_logfile_backups=10
        stderr_capture_maxbytes=1MB
        stderr_events_enabled=false
        environment=A="1",B="2"
        serverurl=AUTO

        # notice:
            program name is used for process name by default
            process name use the program name by default
            
        # take a look at fields above
        -   
            command
            
            /usr/local/bin/program
            
            program
            
            program -p "foo bar"
            
            program --port=80%(process_num)02d      # group_name, host_node_name, process_num, program_name, here

        -
            process_name

            A Python string expression that is used to compose the supervisor process name for this process. You usually don’t need to worry about setting this unless you change numprocs. The string expression is evaluated against a dictionary that includes group_name, host_node_name, process_num, program_name, and here (the directory of the supervisord config file).

            Default: %(program_name)s

        -
            numprocs

            Supervisor will start as many instances of this program as named by numprocs. Note that if numprocs > 1, the process_name expression must include %(process_num)s (or any other valid Python string expression that includes process_num) within it.

            Default: 1

            Required: No.

            Introduced: 3.0
        -
            priority

            The relative priority of the program in the start and shutdown ordering. Lower priorities indicate programs that start first and shut down last at startup and when aggregate commands are used in various clients (e.g. “start all”/”stop all”). Higher priorities indicate programs that start last and shut down first.

            Default: 999

            Required: No.

            Introduced: 3.0

    [group:foo]
        programs=bar,baz
        priority=999

    [eventlistener:x]
    [rpcinterface:x]
    