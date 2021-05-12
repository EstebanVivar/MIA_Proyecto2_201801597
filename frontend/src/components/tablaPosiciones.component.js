import React, { Component } from "react";
import axios from "axios";
import { Column } from 'primereact/column';
import { Toast } from 'primereact/toast';

import { DataTable } from 'primereact/datatable';

export default class Tabla extends Component {
    state = {
        data: [],
        list: [],
        selected: ""
    };

    onSubmit = async (e) => {
        e.preventDefault();
        await this.sendPost();
    };
    async componentDidMount() {

        this.sendGet();
    }
    OnInputChange = e => {
        this.setState({
            [e.target.name]: e.target.value,
        })
    }
    sendGet = async () => {
        await axios
            .get("http://localhost:4000/temporadas/")
            .then(response => {
                this.setState({ list: response.data })
            });
    }


    sendPost = async () => {
        const season = {
            nombre: this.state.selected
        }

        await axios.post("http://localhost:4000/tablaPosiciones/", season)
            .then(response => {
                console.log(response.data)
                this.setState({
                    data: response.data
                })

            });
        console.log(this.state.data)
    }

    render() {        

        return (


            <div className="container" >

                <h3>Selecciona una temporada</h3>

                <div className="form-group">
                    <select
                        className="dropdown-toggle"
                        name="selected"
                        onChange={this.OnInputChange}
                        value={this.state.selected}
                    >
                        {this.state.list.map((valor, index) => (
                            <option key={index} value={valor.nombre}>
                                {valor.nombre}
                            </option>
                        ))}
                    </select>
                </div>




                <form onSubmit={this.onSubmit}>
                    <button type="submit" className="btn btn-primary btn-block">Aplicar</button>
                </form>

                <div className="datatable-selection-demo">
                    <Toast ref={(el) => this.toast = el} />

                    <div className="card">
                        <h5>tabla de posiciones</h5>

                        <DataTable value={this.state.data} selectionMode="single" dataKey="id">
                        
                        <Column field="user" header="Usuario"></Column>
                            <Column field="name" header="Nombre"></Column>
                            <Column field="last" header="Apellido"></Column>
                            <Column field="tier" header="Membresia"></Column>
                            <Column field="season" header="Temporada"></Column>
                            <Column field="p10" header="P10"></Column>
                            <Column field="p5" header="P5"></Column>
                            <Column field="p3" header="P3"></Column>
                            <Column field="p0" header="P0"></Column>
                            <Column field="total" header="Total"></Column>
                        </DataTable>
                    </div>
                </div>
            </div>
        );
    }

}
