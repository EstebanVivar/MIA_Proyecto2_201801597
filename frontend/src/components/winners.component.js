import React, { Component } from "react";
import axios from "axios";
import { Bar } from "react-chartjs-2";

export default class Winners extends Component {
    state = {
        labels: [],
        data: [],
    };

    async componentDidMount() {
        await axios.get("http://localhost:4000/ganadores/")
            .then(response => {
                console.log(response)
                this.setState({
                    labels: response.data.labels,
                    data: response.data,
                })
               
            });
            console.log(this.state.data)
    }

    render() {
        const data = {
            labels: ['Primero','Segundo','Tercero'],
            datasets: this.state.data,
           
           
        };

        const options = {
            scales: {
                xAxes: [{
                    stacked: true
                }],
                yAxes: [{
                    stacked: false
                }]
            }
        };
        return (

            <div className="container" >
                <h3 >Ganadores por membresia</h3>
                    <Bar data={data} options={options} />
            </div>
        );
    }

}
