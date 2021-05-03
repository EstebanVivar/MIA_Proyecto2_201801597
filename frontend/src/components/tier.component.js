import React, { Component } from "react";
import axios from 'axios';
export default class Login extends Component {

    state = {
        selected: "",
        precio: "",
        list: []
    }
    data = localStorage.getItem('user')

    componentDidMount() {
        this.sendGet();
    }


    sendGet = async () => {
        const res = await axios
            .get("http://localhost:4000/membresias/")
            .then(response => {
                this.setState({ list: response.data })
                console.log(this.state.list)
            });
    }

    OnInputChange = e => {
        this.setState({
            [e.target.name]: e.target.value,
            precio: this.state.list.find(element => element.id === e.target.value).precio
        });
    }


    onSubmit = async (e) => {
        e.preventDefault();
        await this.sendPost();

    };

    sendPost = async () => {
        const detail = {
            subscription: 'Y',
            user: this.data.id,
            tier: this.state.selected
        }
        await axios
            .post("http://localhost:4000/tier/", detail)

    };

    render() {
        return (

            <div className="container">

                <h3>Seleccionar membresia</h3>
                <label>Membresia</label>

                <div className="form-group">
                    <select
                        className="custom-select"
                        name="selected"
                        onChange={this.OnInputChange}
                        value={this.state.selected}
                    >
                        {this.state.list.map((valor, index) => (
                            <option key={index} value={valor.id}>
                                {valor.descripcion}
                            </option>
                        ))}
                    </select>

                    <label>Precio: Q  {this.state.precio}

                    </label>
                </div>


                <form onSubmit={this.onSubmit}>
                    <button type="submit" className="btn btn-primary btn-block">Aplicar</button>
                </form>
            </div>
        );
    }
}