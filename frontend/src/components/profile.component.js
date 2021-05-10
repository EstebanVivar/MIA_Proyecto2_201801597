import React, { Component } from "react";
import axios from 'axios';

export default class Profile extends Component {
    user = JSON.parse(localStorage.getItem('user'));
    state = {
        usuario: this.user.user,
        clave: this.user.pass,
        nombre: this.user.name,
        apellido: this.user.last,
        nacimiento: this.user.birth,
        correo: this.user.email,
        foto: this.user.photo,
        id: this.user.id
    }
    OnInputChange = e => {
        this.setState({
            [e.target.name]: e.target.value,
        });
        console.log(this.state)
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
            photo: this.state.foto,
            id:this.state.id
        }
        const res = await axios
            .post("http://localhost:4000/actualizar/", user)

        console.log(res);
    }
    render() {
        return (

            <div className="container">

                    <div className="row  align-items-center">
                        <div className="col-md-12 text-center ">
                            <div className="avatar avatar-xl mb-3" >
                                <img src="https://bootdey.com/img/Content/avatar/avatar6.png" alt="..." className="avatar-img rounded-circle" />
                                <p className="small mb-3"><span className="badge badge-dark">Invitado</span></p>
                            </div>
                        </div>
                    </div>
                    <input type="text" name="usuario" value={this.state.usuario} onChange={this.OnInputChange} className="form-control "
                        style={{ textAlign: 'center', position: 'relative', left: '42%', maxWidth: 16 + '%' }} />



                    <hr className="my-4 " />
                    <div className="form-row ">
                        <div className="form-group col-md-6 ">
                            <label >Nombre</label>
                            <input type="text" name="nombre" value={this.state.nombre} onChange={this.OnInputChange} className="form-control" />
                        </div>
                        <div className="form-group col-md-6">
                            <label >Apellido</label>
                            <input type="text" name="apellido" value={this.state.apellido} onChange={this.OnInputChange} className="form-control" />
                        </div>
                    </div>
                    <div className="form-group">
                        <label>Correo</label>
                        <input type="text" name="correo" value={this.state.correo} onChange={this.OnInputChange} className="form-control" />

                    </div>
                    <div className="form-group">
                        <label>Fecha de nacimiento</label>
                        <input type="date" name="nacimiento" value={this.state.nacimiento} onChange={this.OnInputChange} className="form-control" />
                    </div>


                    
                            <div className="form-group">
                                <label >Cambiar contraseña Actual</label>
                                <input type="password" name="clave" value={this.state.clave} onChange={this.OnInputChange} className="form-control" />
                            </div>
                            <div className="row  align-items-center">
                                <div className="col-md-12 text-center ">
                                    <form onSubmit={this.onSubmit}>
                                        <button type="submit" className="btn btn-primary">Guardar Cambios</button>
                                    </form>
                                </div>
                            </div>
                        </div>


        
        )
    }
}

