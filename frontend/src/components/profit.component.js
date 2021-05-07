import React, { Component } from "react";
import axios from "axios";
import { Line } from "react-chartjs-2";

export default class RepoOjiva extends Component {
    state = {
        labels: [],
        data: [],
        year: ""
    };
    OnInputChange = e => {
        this.setState({
            [e.target.name]: e.target.value,
        });
    }
    onSubmit = async (e) => {
        e.preventDefault();
        await this.sendPost();
    };

    async componentDidMount() {
        await axios.get("http://localhost:4000/ganancias/")
            .then(response => {
                console.log(response)
                this.setState({
                    labels: response.data.labels,
                    data: response.data.data,
                })
            });
    }
    async sendPost() {
        const anyo = {
            year: this.state.year
        }
        const res = await axios.post("http://localhost:4000/gananciasY/", anyo)
            .then(response => {
                console.log(response)
                this.setState({
                    labels: response.data.labels,
                    data: response.data.data,
                })
            });
        console.log(anyo)
    }
    year = {
        year: (new Date()).getFullYear()
    };
    years = {
        array: Array.from(new Array(20), (val, index) => index + this.year.year - 5)
    }
    render() {



        const data = {
            labels: this.state.labels,
            datasets: [
                {
                    label: "Ganancias",
                    data: this.state.data,
                    fill: false,
                    borderWidth: 5,

                    backgroundColor: "rgb(0,204,204)",
                    borderColor: "rgba(0,204,204, 1)",
                },
            ],
        };

        const options = {
            scales: {

                yAxes: [
                    {

                        ticks: {
                            beginAtZero: true,
                        },
                    },
                ],

            },
        };
        return (


            <div className="container" >
                <div className="form-group">
                    <label>Ingrese un año</label>
                
                    <div className="form-group">
                <select
                    onChange={this.OnInputChange}
                    value={this.state.year}
                    className="dropdown-toggle"
                    name="year">
                        
                    {
                        this.years.array.map((year, index) => {
                            return <option key={`year${index}`} value={year}>{year}</option>
                        })
                    }
                </select>
                </div></div>

                <form onSubmit={this.onSubmit}>
                    <button type="submit" style={{width:160+'px'}} className="btn btn-primary btn-block">Obtener reporte</button>
                </form>
                <br /><br />
                <div className="form-group">
                    <h3 >Grafica de Ganancias</h3>
                    <div style={{ backgroundColor: "#181C30" }}>
                        <Line data={data} options={options} />
                    </div>
                </div>
            </div>

        );
    }

}
