# go-auth
## Dependencies

* [Docker](http://docker.io)

## Usage

	go-auth - A RESTful Authentication Service with a Database backend

	USAGE:
	   go-auth [global options] command [command options] [arguments...]

	VERSION:
	   0.0.0

	AUTHOR:
	  Usman Ismail - <usman@techtraits.com>

	COMMANDS:
	   run		Run the authentication service
	   help, h	Shows a list of commands or help for one command

	GLOBAL OPTIONS:
	   --log-type -t "syslog, console"
	   --log-level, -l "Info"	The log level to use
	   --help, -h			show help
	   --version, -v		print the version

	# For Example
	go-auth -l debug run --db-host 192.168.59.103 -p 8080


To run a containerized mysql database for your application use the following command:

    docker run -d --name mysql -e MYSQL_ROOT_PASSWORD=rootpass \
        -e MYSQL_DATABASE=messenger \
        -e MYSQL_USER=messenger \
        -e MYSQL_PASSWORD=messenger \
        -p 3306:3306 mysql
        
##### To Add a new User

    curl -i -X PUT -d userid=USERNAME -d password=PASSWORD localhost:8080/user
    
##### To Delete a User

    curl -i -X DELETE 'localhost:8080/user?userid=USERNAME&password=PASSWORD'
    
##### To Get an Auth Token

    curl 'http://localhost:8080/token?userid=USERNAME&password=PASSWORD'

##### To Verify an Auth Token

    curl -i -X POST 'localhost:8080/token/USERNAME' --data "IHuzUHUuqCk5b5FVesX5LWBsqm8K...."

## Building

    ./build.sh
