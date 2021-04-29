import React, { Component } from "react";
import axios from 'axios';
export default class SignUp extends Component {

    state = {
        usuario: "",
        clave: "",
        nombre: "",
        apellido: "",
        nacimiento: "",
        correo: "",
        foto: ""
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
            pass: this.state.clave,
            name: this.state.nombre,
            last: this.state.apellido,
            birth: this.state.nacimiento,
            email: this.state.correo,
            photo: this.state.foto
        }
        const res = await axios
            .post("http://localhost:4000/registrar/", user)

        console.log(res);
    }

    render() {
        return (
            <div className="container">
                <h3>Registrar usuario</h3>

                <div className="form-group">
                    <label>Nombre de usuario</label>
                    <input type="text" name="usuario"  value={this.state.usuario} onChange={this.OnInputChange} className="form-control" placeholder="Ingresar nombre de usuario" />
                </div>

                <div className="form-group">
                    <label>Contraseña</label>
                    <input type="password" name="clave"  value={this.state.clave} onChange={this.OnInputChange} className="form-control" placeholder="Ingresar contraseña" />
                </div>

                <div className="form-group">
                    <label>Nombre</label>
                    <input type="text" name="nombre"  value={this.state.nombre} onChange={this.OnInputChange} className="form-control" placeholder="Ingresar nombre" />
                </div>

                <div className="form-group">
                    <label>Apellido</label>
                    <input type="text" name="apellido"  value={this.state.apellido} onChange={this.OnInputChange} className="form-control" placeholder="Ingresar apellido" />
                </div>

                <div className="form-group">
                    <label>Fecha de nacimiento</label>
                    <input type="date" name="nacimiento"  value={this.state.nacimiento} onChange={this.OnInputChange} className="form-control" />
                </div>

                <div className="form-group">
                    <label>Correo electrónico</label>
                    <input type="email" name="correo"  value={this.state.correo} onChange={this.OnInputChange} className="form-control" placeholder="Ingresar correo electrónico" />
                </div>

                <div className="form-group">
                    <label>Foto de perfil</label>
                    <input type="file" name="foto"  value={this.state.foto} onChange={this.OnInputChange} className="form-control" placeholder="Seleccionar foto" />
                </div>

                <form onSubmit={this.onSubmit}>
                    <button type="submit" className="btn btn-primary btn-block">Registrar</button>
                </form>

            </div>
        );
    }
}