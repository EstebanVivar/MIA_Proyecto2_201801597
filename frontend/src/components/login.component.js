import React, { Component } from "react";
import axios from 'axios';
export default class Login extends Component {

    state = {
        usuario: "",
        clave: ""
    }
    isUser() {
        var data = JSON.parse(localStorage.getItem('user'))
        console.log(data.id)
        if (data.id !== "") {
            return (
                this.props.history.push("/perfil")
            )

        }


    }
    componentDidMount() {
        // this.sendInfo();
    }

    OnInputChange = e => {
        this.setState({
            [e.target.name]: e.target.value,
        });
    }


    onSubmit = async (e) => {
        e.preventDefault();
        await this.sendPost();

    };

    sendPost = async () => {
        const user = {
            user: this.state.usuario,
            pass: this.state.clave
        }
        const res = await axios
            .post("http://localhost:4000/login/", user)
            .then(response => {
                localStorage.setItem('user', JSON.stringify(response.data));
                this.isUser()

            })
    };

    render() {
        return (

            <div className="container">

                <h3>Iniciar sesión</h3>

                <div className="form-group">
                    <label>Usuario</label>
                    <input type="text" name="usuario" value={this.state.user} onChange={this.OnInputChange} className="form-control" placeholder="Ingrese su nombre de usuario" />
                </div>

                <div className="form-group">
                    <label>Contraseña</label>
                    <input type="password" name="clave" value={this.state.pass} onChange={this.OnInputChange} className="form-control" placeholder="Ingrese su contraseña" />
                </div>

                <form onSubmit={this.onSubmit}>

                    <button type="submit" className="btn btn-primary btn-block">Ingresar</button>

                </form>
            </div>
        );
    }
}