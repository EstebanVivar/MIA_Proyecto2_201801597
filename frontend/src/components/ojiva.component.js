import React, { Component } from "react";
import axios from "axios";
import { Line } from "react-chartjs-2";

export default class Ojiva extends Component {
    state = {
        labels: [],
        data: [],
    };

    async componentDidMount() {
        await axios.get("http://localhost:4000/ojiva/")
            .then(response => {
                console.log(response)
                this.setState({
                    labels: response.data.labels,
                    data: response.data.data,
                })
            });
    }

    render() {
        const data = {
            labels: this.state.labels,
            datasets: [
                {
                    label: "Frecuencia de ganancias",
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
                <h3 >Ojiva de Ganancias</h3>
                <div className="dark" style={{ backgroundColor: "#181C30" }}>
                    <Line data={data} options={options} />
                </div>
            </div>
        );
    }

}
