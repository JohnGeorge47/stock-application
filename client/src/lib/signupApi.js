import axios from "axios";
import "../"

const formUrlEncoded = x =>
Object.keys(x).reduce((p, c) => p + `&${c}=${encodeURIComponent(x[c])}`, '')

const SignUpApi=(email_id,password,username)=>{
        let data={
                user_email:email_id,
                user_name:username,
                password:password
        }
        let encodedData=formUrlEncoded(data)
        console.log(data)
        return axios({
                method:"post",
                url:" http://ae641705e92e.ngrok.io/create_user",
                data:encodedData,
                config: {
                        headers: {
                            'Content-Type': 'application/x-www-form-urlencoded',
                        },
                    },
        })
}
export default SignUpApi