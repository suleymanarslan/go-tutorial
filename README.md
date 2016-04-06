# hoditgo
 Install Go 

$ sudo add-apt-repository ppa:ubuntu-lxc/lxd-stable
$ sudo apt-get update
$ sudo apt-get install golang

Install Redis with the steps mentioned in the link below:

https://www.digitalocean.com/community/tutorials/how-to-install-and-use-redis

Install Mysql 

Declare following environment variables:

go path.
export PATH=$PATH:/usr/local/go/bin 

go project workspace path.
export GOPATH=$HOME/work

compiled binary path.
export GOBIN=$HOME/work/bin

project path for settings.
export HODITGO=$HOME/work/src/hoditgo

go to your project directory $HODITGO directory
cd $HODITGO

get project from github:

git pull https://github.com/suleymanarslan/hoditgo.git

After getting the whole project, run
go get 

Configuration file for development is under

$HODITGO/settings/pre.json

Add or edit your mysql connection and redis parameters in the pre.json file.

after getting all dependencies, run:
go run server.go

Try to create new user with the following command:

curl -H "Content-Type: application/json" -X POST -d '{"Username":"suleyman","Password":"dummy","Email": "me@example.com"}' http://localhost:5000/create-user

Try to login with the following command:
curl -H "Content-Type: application/json" -X POST -d '{"Email":"me@suleymanarslan.com","Password":"354216"}' http://localhost:5000/token-auth

When you successfully logged on to the system you will get JWT like the following one.

{"token":"eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjAyMzEzNzYsImlhdCI6MTQ1OTk3MjE3Niwic3ViIjoiIn0.xpbzlKTfLtUzqXwMFiYdmW6hdfg0sMLeHq2m1yz3cGZATgADZWAK8V7qzsXTbgWJ9X0jsplWM5RJlPsHC0AJKd1P6XkH7NdcAppo6ILOeE7RwukiPDc3TDwiIjFb539YxwZFYBnW3D5UdF2_jGMfZUTMWv9viHBJ43K2S_rHvixMlQdXeF0TJoI_JfxeXdUDWRthrc22na2k2rK5_Ethe3pCWOa25iYpnWwbcNRJEw5ZWnTifEcyliPc5hecPnfcw4izylVxGKyY0xGDGOeQz6IHKt3D6AmM_LHpKW8rIfD68UbFItpeTixvOuN2hbAyEey5zj66GwHZTE-8ouQtQw"}


Use this token to get response from hello service.

curl -H "Authorization: Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjAyMzE4NTgsImlhdCI6MTQ1OTk3MjY1OCwic3ViIjoiIn0.nK0FMNmVFv__N2BkVi3nibBi8csp-LYjSWr6irMyW_XMKEbm5IvRDe51XvK4WTAf96rlkUyQGQGctQtTUb9rQYWt8Xd5HIjycHD_CMgs07MZuW6RDHho5VB90bYxk1L1jA5guMEF6oDR_WjmmQ5LMWxDMZDuztWJI-4YGMH11eiFTjt4IOkQxni0tnS5dOSusBnH7PmOuH247fxq4WrQCgjbhc429MlE_XiOc7nAyy9uL-IZHvWTOTRnOpv9Zm5OX_UdV8ySFrQePs1FXlF__cDvQOR9MjyraeVovHfQdpOWttARlXg9pgnEyOEm6N47RYLQUvPFjN6qbtIAGXnnEw" http://localhost:5000/test/hello


In order to logout from the service:

curl -H "Authorization: Bearer eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0NjAyMzE4NTgsImlhdCI6MTQ1OTk3MjY1OCwic3ViIjoiIn0.nK0FMNmVFv__N2BkVi3nibBi8csp-LYjSWr6irMyW_XMKEbm5IvRDe51XvK4WTAf96rlkUyQGQGctQtTUb9rQYWt8Xd5HIjycHD_CMgs07MZuW6RDHho5VB90bYxk1L1jA5guMEF6oDR_WjmmQ5LMWxDMZDuztWJI-4YGMH11eiFTjt4IOkQxni0tnS5dOSusBnH7PmOuH247fxq4WrQCgjbhc429MlE_XiOc7nAyy9uL-IZHvWTOTRnOpv9Zm5OX_UdV8ySFrQePs1FXlF__cDvQOR9MjyraeVovHfQdpOWttARlXg9pgnEyOEm6N47RYLQUvPFjN6qbtIAGXnnEw" http://localhost:5000/logout
