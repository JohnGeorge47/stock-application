import axios from "axios";
import "../"

const formUrlEncoded = x =>
Object.keys(x).reduce((p, c) => p + `&${c}=${encodeURIComponent(x[c])}`, '')

const LoginApi=(email_id,password)=>{
        let data={
                email_id:email_id,
                password:password
        }
        let encodedData=formUrlEncoded(data)
        console.log(data)
        return axios({
                method:"post",
                url:" http://ae641705e92e.ngrok.io/login",
                data:encodedData,
                config: {
                        headers: {
                            'Content-Type': 'application/x-www-form-urlencoded',
                        },
                    },
        })
}
export default LoginApi