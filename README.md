This is a simple stock application written in go <br/> and
uses gorilla websocket

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
GET /validate_user?email_id=email&session_token=token