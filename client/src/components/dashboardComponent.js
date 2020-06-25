import React,{Component} from "react"
import Cookies from 'universal-cookie';
 
const cookies = new Cookies();


export default class DashBoardComponent extends Component{
    constructor(props){
        super(props);
        this.state = { messages : "" }
      }
    componentDidMount(){
        let session_token=cookies.get('session_token')
        let email_id=cookies.get('email_id')
        this.connection=new WebSocket(`ws://ae641705e92e.ngrok.io/ws?email_id=${email_id}&session_token=${session_token}`)
        this.connection.onmessage = evt => { 
              this.setState({
              messages : evt.data
            })
          };
    }
    render(){
        let messageJson={}
        let listItems=<li></li>
        if (this.state.messages!==""){
           messageJson=JSON.parse(this.state.messages)
        }
        if (messageJson.message!==undefined){
             listItems = messageJson.message.map((d) => <li key={d.stock_name}>{d.stock_name}:{d.value}</li>);
        }
        return(
            <div>
                <h1 style={{textAlign: 'center'}}>Your Stocks</h1>
                <div style={{textAlign: 'center'}}>
                   {listItems}
                </div>
            </div>
        )
    }
}