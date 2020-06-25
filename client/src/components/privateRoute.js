import React,{ Component } from "react";
import { Redirect } from 'react-router-dom';
import getAuth from "../lib/auth"
import { createHashHistory } from 'history'
export const history = createHashHistory()


class PrivateRoute extends Component{

    state={
        haveAccess:false,
        isLoading:true,

    }

    componentDidMount(){
        this._isMounted = true;
        getAuth().then(res=>{
                if(res.data.valid===true){
                    this.setState({
                        haveAcces:true,
                        isLoading:false
                    })
                    console.log("here")
                    console.log(this.state.haveAccess)
                }
          }).catch(err=>{
            this.setState({isLoading: false});
              console.log(err)
          })
    }
 
    checkAccess=()=>{
    }
    render() {
        const { component: Component} = this.props;
        const {haveAccess,isLoading} = this.state;
        console.log(haveAccess)
        if(isLoading) {
            return <div>Loading...</div>
        }
        if(!haveAccess&&isLoading) {
            return <Redirect to="/login" />
        }
        return <Component {...this.props} /> 
      }
}

export default PrivateRoute