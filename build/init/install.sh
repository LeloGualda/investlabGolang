echo "install golang: get websocket"
go get github.com/gorilla/websocket
echo "install golang: install websocket"
go install github.com/gorilla/websocket
echo "install golang: install mysql"
go install github.com/go-sql-driver/mysql
echo "install golang: crypy password"
go get -u golang.org/x/crypto/bcrypt