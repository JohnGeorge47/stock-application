This is a simple stock application written in go <br/> and
uses gorilla websocket and MySQL.The db schema is stored in 
the db/db.sql and the entry point of the application is cmd/http/server.go
You can do a go run of server.go or build it and use the binary and it will run
on port 3003

The entire backend is build using the standard go library except for gorilla websocket<br/>

There is an http server which accepts ws connections<br/>

There is a seperate goroutine which runs which sends a message to all online sessions of a users sending
the user info about the stocks he/she is subscribed to<br/>

There  is another goroutine which continously updates all stocks to mimic
changes in the stock value

The front end is very minimal and uses react,react-router and ant design there
are three pages /login /signup dashboard . The application is running on
port 3000


The api's available are <b/>

POST /create_user <br/>
{<br/>
    user_email:emailid,<br/>
    user_name:username,<br/>
    password:password<br/>
}<br/>
POST /login <br/>
{<br/>
    email_id:email_id,<br/>
    password:password<br/>
}<br/>

Both the login as well as signup api responds
with a session token which needs to be used to make further requests.

GET /validate_user?email_id=email&session_token=token