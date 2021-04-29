import React,{Component} from 'react';
import axios from 'axios';

class HomeContainer extends Component{
    componentDidMount(){
        axios.get('https://pokeapi.co/api/v2/pokemon/ditto/')
        .then(result=>{
            console.log(result)
        }).catch(console.log)
    }
    render(){
        return(<div></div>);
    }
}

export default HomeContainer;