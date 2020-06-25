import React, { Component } from "react";
import { Form, Input, Button,Row, Col } from 'antd';
import LoginApi from "../lib/loginApi"
import Cookies from 'universal-cookie';

const cookies=new Cookies()



export default class Login extends Component {
    constructor(props){
        super()
    }
    onFinish = values => {
        LoginApi(values.email,values.password).then(res=>{
          console.log(res)
          console.log("here")
          cookies.set('email_id',values.email,{path:'/'})
          cookies.set('session_token',res.data.request_token,{path:'/'})
          
          window.sessionStorage.setItem('email_id',values.email)
          window.sessionStorage.setItem('session_token',res.data.request_token)
          this.props.history.push("/dashboard")
        }).catch(err=>{
          console.log(err)
        })
      };
    

    render() {
        return (
          <Row align="middle" type="flex">
            <Col span={6} offset={8}>
             <h1 style={{textAlign: 'center'}}>Log in</h1>   
            <Form
            name="basic"
            initialValues={{
              remember: true,
            }}
            onFinish={this.onFinish}
          >
            <Form.Item
              label="Email-id"
              name="email"
              rules={[
                {
                  required: true,
                  message: 'Please inpt your email-id',
                },
              ]}
            >
              <Input/>
            </Form.Item>
            <Form.Item
              label="Password"
              name="password"
              rules={[
                {
                  required: true,
                  message: 'Please input your password!',
                },
              ]}
            >
              <Input.Password />
            </Form.Item>
            <Form.Item>
                <div  style={{textAlign: 'center'}}>
              <Button type="primary" htmlType="submit">
                Log in
              </Button>
              </div>
            </Form.Item>
          </Form>
          </Col>
          </Row>
        );
    }
}