This is a simple web application which runs in an ops environment. 

1. Compile the executable for the ops environment using:

GOOS=linux go build count_visitor.go

2. Run the application within ops environment using config file

ops run -c config.json count_visitor
