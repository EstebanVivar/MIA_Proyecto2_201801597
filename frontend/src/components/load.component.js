import React from "react";
import yaml from "js-yaml";
import { FilePicker } from "react-file-picker";
import axios from 'axios';


export default class CargaMasiva extends React.Component {
  load={};
  handleFileChange = file => {
    const reader = new FileReader();
    reader.readAsText(file);
    reader.onload = e => {
      try {
        const doc = yaml.load(e.target.result);
        this.load = JSON.stringify(doc, null, 2);
        console.log(this.load);
        this.sendPost();
      } catch (e) {
        console.log(e);
      }
    };

    this.setState({ title: file.name });
  };
  sendPost = async () => {
    
    const res = await axios
        .post("http://localhost:4000/test/", this.load)
    console.log(res)
};
  render() {
    return (
      <div className="container">


        <FilePicker
          extensions={["yaml"]} 
          onChange={this.handleFileChange}
          onError={errMsg => console.log(errMsg)}
        >
          <button className='btn btn-primary'>
            Seleccionar Archivo
          </button>
        </FilePicker>
      </div>
    )
  }

}