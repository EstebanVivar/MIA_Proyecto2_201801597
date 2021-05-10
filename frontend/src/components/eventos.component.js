
import React, { Component } from 'react'
import FullCalendar from '@fullcalendar/react'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin from '@fullcalendar/interaction'
import esLocale from '@fullcalendar/core/locales/es';

import axios from 'axios';


export default class Eventos extends Component {

  state = {
    calendarEvents: [],
    info: [],
    Events: [],
    E_local: "",
    E_visita: "",
    E_deporte: 1,
    E_fecha: "",
    E_jornada: 137,
    R_visita:"",
    R_local:""
  }


  componentDidMount() {
    this.sendGet();
  }

  fixEvent(E_id, local, visit, s_local, s_visit, fechaI, fechaF) {
    this.state.Events.push({
      id: E_id,
      title: ' ' + local + ' - ' + visit,
      local: local,
      visit: visit,
      s_local: s_local,
      s_visit: s_visit,
      start: fechaI,
      end: fechaF
    });
  }
   /////////////////////////////////
  OnInputChange = e => {
    this.setState({
      [e.target.name]: e.target.value,
    });
    console.log(this.state)
  }
  /////////////////////////////////
  onSubmitResultado = async (e) => {
    e.preventDefault();
    await this.sendPostResultado();

  };

  sendPostResultado = async () => {
    const results = {
      id: this.state.info.id,
      r_local: this.state.R_local,
      r_visita: this.state.R_visita
    }
    await axios.post("http://localhost:4000/u_results/", results)
      .then(response => {
        console.log(response.data)
      });
      this.state.Events=[];
      this.sendGet();
      
  }
  /////////////////////////////////
  onSubmitEvento = async (e) => {

    e.preventDefault();
    await this.sendPostEvento();
  };

  sendPostEvento = async () => {
    const event = {
      local: this.state.E_local,
      visita: this.state.E_visita,
      deporte: this.state.E_deporte,
      fecha_inicio: this.state.E_fecha,
      jornada: this.state.E_jornada
    }
    await axios.post("http://localhost:4000/crear_evento/", event)
      .then(response => {
        console.log(response.data)
      });
      this.state.Events=[];
      this.sendGet();
  }
  /////////////////////////////////
  sendGet = async () => {
    await axios
      .get("http://localhost:4000/eventos/")
      .then(response => {
        response.data.forEach(element => {
          this.fixEvent(element.id, element.local, element.visita, element.m_local, element.m_visita, element.fecha_inicio, element.fecha_final);
        });
        this.setState({
          calendarEvents: this.state.Events
        })
        console.log(response)
      });
  }
  /////////////////////////////////


  render() {
    return (
      <div className='eventos-app'>
        {this.renderSidebar()}
        <div className='eventos-app-main'>
          <FullCalendar
            plugins={[dayGridPlugin, timeGridPlugin, interactionPlugin]}
            headerToolbar={{
              left: 'prev,next today',
              center: 'title',
              right: 'dayGridMonth,timeGridWeek'
            }}
            locale={esLocale}
            initialView='dayGridMonth'
            contentHeight={537 + 'px'}
            eventClick={this.handleEventClick}
            events={this.state.calendarEvents}
          /> 
        </div>
      </div>
    )
  }
  renderSidebar() {
    return (
      <div className='Container'>
        <div className='eventos-app-sidebar-section'>
          <div className="card text-center">
            <div className="card-body">
              <label className="card-title">Local: {this.state.info.local}<br />Visitante: {this.state.info.visit} </label>
              <div className="form-group">
                <b>Resultado</b>
                <br />
                <label>Local: &nbsp;</label>
                <input type="text" name="R_local" value={this.state.R_local} onChange={this.OnInputChange} placeholder="Resultado local" />
              </div>
              <div className="form-group">
                <label>Visita: &nbsp;</label>
                <input type="text" name="R_visita" value={this.state.R_visita} onChange={this.OnInputChange} placeholder="Resultado visitante" />
              </div>
              <form onSubmit={this.onSubmitResultado}>
                <button type="submit" className="btn btn-primary">Ingresar</button>
              </form>
            </div>
          </div>
          <br />


          <div className="card text-center">
            <div className="card-body">
              <h5>Crear Evento</h5>
              <div className="form-group">
                <label>&nbsp;&nbsp;&nbsp;&nbsp;Local:&nbsp;</label>
                <input type="text" name="E_local" value={this.state.E_local} onChange={this.OnInputChange} placeholder="Equipo local" />
              </div>

              <div className="form-group">
                <label>&nbsp;&nbsp;&nbsp;&nbsp;Visita:&nbsp;</label>
                <input type="text" name="E_visita" value={this.state.E_visita} onChange={this.OnInputChange} placeholder="Equipo visitante" />
              </div>

              <div className="form-group">
                <label>Deporte:&nbsp;</label>
                <input type="text" name="E_deporte" value={this.state.E_deporte} onChange={this.OnInputChange} placeholder="Deporte" />
              </div>

              <div className="form-group">
                <label>Fecha I.:&nbsp;</label>
                <input type="text" name="E_fecha" value={this.state.E_fecha} onChange={this.OnInputChange} placeholder="DD-MM-YYYY HH24:MI" />
              </div>

              <form onSubmit={this.onSubmitEvento}>
                <button type="submit" className="btn btn-primary">Ingresar</button>
              </form>
            </div>
          </div>
        </div>
      </div>

    )
  }

  options = {
    day: "numeric",
    year: "numeric",
    month: "long",
    weekday: "long",
    hour: "numeric",
    minute: "numeric",
  }

  handleEventClick = (clickInfo) => {
    this.renderSidebar()
    return (
      this.setState({
        info: {
          end: Intl.DateTimeFormat('es-ES', this.options).format(clickInfo.event.end),
          title: clickInfo.event.title,
          id: clickInfo.event.id,
          local: clickInfo.event._def.extendedProps.local,
          visit: clickInfo.event._def.extendedProps.visit,
          s_local: clickInfo.event._def.extendedProps.s_local,
          s_visit: clickInfo.event._def.extendedProps.s_visit,
        },
        R_local:clickInfo.event._def.extendedProps.s_local,
        R_visita:clickInfo.event._def.extendedProps.s_visit,
      }))
  }
}

