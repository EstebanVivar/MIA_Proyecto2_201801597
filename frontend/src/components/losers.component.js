import React, { Component } from "react";
import axios from "axios";
import { Pie } from "react-chartjs-2";

export default class Losers extends Component {
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

        await axios.post("http://localhost:4000/perdedores/", season)
            .then(response => {
                console.log(response.data)
                this.setState({
                    data: response.data
                })

            });
        console.log(this.state.data)
    }

    render() {
        const data = {
            labels: ["Bronze", "Silver", "Gold"],
            datasets: [
                {
                    label: "Perdedores",
                    backgroundColor: this.state.data.backgroundColor,
                    data: this.state.data.data,
                },
            ],
        };

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

                <h3 >Ganadores por membresia</h3>
                <div style={{marginLeft:20+'%',width:60+'%'}}>
                <Pie  data={data}          options={{responsive: true, maintainAspectRatio: true}}  />
                </div>
            </div>
        );
    }

}
