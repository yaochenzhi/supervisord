go installation:

    download go zip file

    unzip file 

    mv go/bin/* /usr/local/bin/ 

    mv go /usr/local/

go prog:
    
    fastcgi.go 

start:

    supervisorctl update
    supervisorctl start fastcgi_test:*

stop:

    supervisorctl start fastcgi_test:fastcgi_test_03
    supervisorctl start fastcgi_test:*

test:
    
    for i in `seq 1 10`; do curl 'http://127.0.0.1:8080/app?helloworld' & done

