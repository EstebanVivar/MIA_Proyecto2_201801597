import React, { Component } from 'react';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import Login from "./components/login.component";
import SignUp from "./components/signup.component";
import Eventos from "./components/eventos.component";
import A from "./components/load.component";
// eslint-disable-next-line
import axios from 'axios'

class App extends Component {
  render() {
    return (this.apx())
  }
  apx() {
    return (
      <Router>
        <div className="App">
          <nav className="navbar navbar-expand-lg navbar-light fixed-top">
            <div className="container">
              <Link className="navbar-brand" to={"/registrar"}>QUINIELAS</Link>
              <div className="collapse navbar-collapse" id="navbarTogglerDemo02">
                <ul className="navbar-nav ml-auto">
                  <li className="nav-item">
                    <Link className="nav-link" to={"/eventos"}>Eventos</Link>
                  </li>
                  <li className="nav-item">
                    <Link className="nav-link" to={"/ingresar"}>Iniciar sesion</Link>
                  </li>
                  <li className="nav-item">
                    <Link className="nav-link" to={"/registrar"}>Registrarme</Link>
                  </li>
                </ul>
              </div>
            </div>
          </nav>

          <div className="auth-wrapper">
            <div className="auth-inner">
              <Switch>
                <Route exact path='/' component={Login} />
                <Route path="/ingresar" component={Login} />
                <Route path="/registrar" component={SignUp} />
                <Route path="/test" component={A} />
                <div className="event-inner">
                  <Route path="/eventos" component={Eventos} />
                </div>
              </Switch>
            </div>

          </div>
        </div>

      </Router>
    )
  }
}

export default App;