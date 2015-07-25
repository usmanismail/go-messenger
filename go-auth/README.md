# go-auth

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
