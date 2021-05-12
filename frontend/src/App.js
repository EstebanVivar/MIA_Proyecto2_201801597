import React, { Component } from 'react';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import Login from "./components/login.component";
import SignUp from "./components/signup.component";
import Eventos from "./components/eventos.component";
import EventosUsuario from "./components/eventosUsuario.component";
import Carga from "./components/load.component";
import Profile from "./components/profile.component";
import Membresia from "./components/tier.component";
import Ojiva from "./components/ojiva.component";
import Ganancias from "./components/profit.component";
import Ganadores from "./components/winners.component";
import Perdedores from "./components/losers.component";
import Tabla from "./components/tablaPosiciones.component";


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
              <Link className="navbar-brand" to={"/ingresar"}>QUINIELAS</Link>
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
                <Route path="/membresia" component={Membresia} />
                <Route path="/test" component={Carga} />
                <Route path="/perfil" component={Profile} />
                <Route path="/ojiva" component={Ojiva} />
                <Route path="/ganancias" component={Ganancias} />

                <Route path="/ganadores" component={Ganadores} />
                <Route path="/perdedores" component={Perdedores} />
               
                <>
                <div className="event-inner">
                  <Route path="/posiciones" component={Tabla} />
                  <Route path="/eventos" component={Eventos} />
                  <Route path="/eventosUsuario" component={EventosUsuario} />
                </div>
                </>
              </Switch>
            </div>

          </div>
        </div>

      </Router>
    )
  }
}

export default App;