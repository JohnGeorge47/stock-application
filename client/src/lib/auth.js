import axios from "axios"
import Cookies from 'universal-cookie';
 import config from "../config.json"
const cookies = new Cookies();

const getAuth=()=>{
   let session_token=cookies.get('session_token')
    let email_id=cookies.get('email_id')
        return axios.get(` http://ae641705e92e.ngrok.io?email_id=${email_id}&session_token=${session_token}`)
}

export default getAuth