
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
    Events: []
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
    })
  }


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
      });
  }


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
            eventClick={this.handleEventClick}
            events={this.state.calendarEvents}
          />
        </div>
      </div>
    )
  }
  renderSidebar() {
    return (
      <div className='eventos-app-sidebar'>
        <div className='eventos-app-sidebar-section'>
          <div className="card text-center">
            <div className="card-header" >
            </div>
            <div className="card-body">
              <h6 className="card-title">Local: {this.state.info.local}</h6>
              <h6 className="card-title">Visitante: {this.state.info.visit} </h6>
              <p className="card-text"> Local - Visitante <br />{this.state.info.s_local} - {this.state.info.s_visit} </p>
            </div>
            <div className="card-footer text-muted">
              {this.state.info.end}
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
    renderCardEvent(clickInfo)
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
        }
      }))
  }
}

function renderCardEvent(event) {
  console.log(event.event.start)
  return (

    <i>{event.event.title}</i>

  )
}