
import React from 'react'
import FullCalendar from '@fullcalendar/react'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin from '@fullcalendar/interaction'
import esLocale from '@fullcalendar/core/locales/es';
import axios from 'axios';


export default class DemoApp extends React.Component {

  state = {
    calendarEvents: [],
    info: []
  }
  state2 = {
    Events: []
  }

  componentDidMount() {
    this.sendGet();
  }


  fixEvent(E_id, local, visit, s_local, s_visit, fechaI, fechaF) {
    let Etitle = local + ' - ' + visit;
    let Edesc = s_local + ' - ' + s_visit;
    let Estart = fechaI
    let Efinal = fechaF
    this.state2.Events.push({
      id: E_id,
      title: Etitle,
      local: local,
      visit: visit,
      description: Edesc,
      start: Estart,
      end: Efinal
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
          calendarEvents: this.state2.Events
        })
      });
  }


  render() {
    return (
      <div className='demo-app'>
        {this.renderSidebar()}
        <div className='demo-app-main'>
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
      <div className='demo-app-sidebar'>
       

        <div className='demo-app-sidebar-section'>
          <h2>Instructions</h2>
          <ul>
            <li>Select dates and you will be prompted to create a new event</li>
            <li>Drag, drop, and resize events</li>
            <li>Click an event to delete it</li>
          </ul>
        </div>

        <div className='demo-app-sidebar-section'>
          <div className="card text-center">
            <div className="card-header" >
            </div>
            <div className="card-body">
              <h5 className="card-title">Local:
                {this.state.info.local}</h5>
                <h5 className="card-title">Visitante:
                {this.state.info.visit}</h5>
                Local-Visitante
              <p className="card-text">{this.state.info.results}</p>
              <div className="form-group">
                    <label>Local</label>
                    <input type="text" name="usuario" value={this.state.user} onChange={this.OnInputChange} className="form-control" placeholder="Ingrese su nombre de usuario" />
                </div>

                <div className="form-group">
                    <label>Visitante</label>
                    <input type="password" name="clave" value={this.state.pass} onChange={this.OnInputChange} className="form-control" placeholder="Ingrese su contraseña" />
                </div>

                <form onSubmit={this.onSubmit}>

                    <button type="submit" className="btn btn-primary btn-block">Ingresar</button>

                </form>
            </div>
            <div className="card-footer text-muted">
              {this.state.info.end}
            </div>
          </div>
        </div>
        {/* <div className='demo-app-sidebar-section'>
          <h2>All Events ({this.state.calendarEvents.length})</h2>
          <ul>
            {this.state.calendarEvents.map(renderSidebarEvent)}
          </ul>
        </div> */}
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
          end: Intl.DateTimeFormat('es-ES', this.options).format(clickInfo.event.start),
          title: clickInfo.event.title,
          id: clickInfo.event.id,
          results: clickInfo.event._def.extendedProps.description,
          visit: clickInfo.event._def.extendedProps.visit,
          local: clickInfo.event._def.extendedProps.local
        }
      }))
  }
}

function renderEventContent(eventInfo) {
  return (
    <>
      {/* <b>{eventInfo.timeText}</b> */}
      <i>{eventInfo.event.title}</i>
    </>
  )
}

// function renderSidebarEvent(event) {
//   return (
//     <li key={event.id}>
//       <b>{formatDate(event.start, { year: 'numeric', month: 'short', day: 'numeric' })}</b>
//       <i>{event.title}</i>
//     </li>
//   )
// }
function renderCardEvent(event) {
  console.log(event.event.start)
  return (

    <i>{event.event.title}</i>

  )
}