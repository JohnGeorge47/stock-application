import React from 'react';
import Login from "./components/loginComponent"
import Signup from "./components/signupComponent"
import DashBoardComponent from "./components/dashboardComponent"
import {Switch,Route} from "react-router-dom"
import PrivateRoute from "./components/privateRoute"
import './App.css';

function App() {
  return (
    <Switch>
      <PrivateRoute
              path="/dashboard"
              component={DashBoardComponent}
      ></PrivateRoute>
      <Route path="/login" component={Login}>
      </Route>
      <Route path="/">
        <Signup/>
      </Route>
    </Switch>
  );
}

export default App;
