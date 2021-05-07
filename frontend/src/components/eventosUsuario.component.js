
import React,{Component} from 'react'
import FullCalendar from '@fullcalendar/react'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin from '@fullcalendar/interaction'
import esLocale from '@fullcalendar/core/locales/es';
import axios from 'axios';


export default class EventosUsuario extends Component {
  data = JSON.parse(localStorage.getItem('user'));
  state = {
    weekendsVisible: true,
    calendarEvents: [],
    info: [],
    Events: []
  }


  componentDidMount() {
    this.sendGet();
  }


  fixEvent(E_id, local, visit, s_local, s_visit, fechaI, fechaF, p_local, p_visit) {

    this.state.Events.push({
      id: E_id,
      title: ' ' + local + ' - ' + visit,
      local: local,
      visit: visit,
      s_local: s_local,
      s_visit: s_visit,
      p_local: p_local,
      p_visit: p_visit,
      start: fechaI,
      end: fechaF
    })
  }


  sendGet = async () => {
    console.log(this.data.id)
    const info = {
      user: this.data.id
    }
    await axios
      .post("http://localhost:4000/eventosUsuario/", info)
      .then(response => {
        if (response) {
          console.log(response.data)
          response.data.forEach(element => {
            this.fixEvent(element.id, element.local, element.visita, element.m_local,
              element.m_visita, element.fecha_inicio, element.fecha_final, element.p_local, element.p_visita);
          });
          this.setState({
            calendarEvents: this.state.Events
          })
        }
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
            editable={false}
            selectable={true}
            eventDisplay={'block'}
            selectMirror={true}
            dayMaxEvents={true}
            weekends={this.state.weekendsVisible}
            select={this.handleDateSelect}
            eventContent={renderEventContent}
            eventClick={this.handleEventClick}
            eventsSet={this.handleEvents}
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
              {this.state.info.id}
            </div>
            <div className="card-body">
              <h6 className="card-title">Local: {this.state.info.local}</h6>
              <h6 className="card-title">Visitante: {this.state.info.visit} </h6>
              <p className="card-text"> Local - Visitante <br />{this.state.info.s_local} - {this.state.info.s_visit} </p>
              <div className="form-group">
                <h4>Prediccion</h4>
                <label>Local</label>
                <input type="text" name="p_local" value={this.state.info.p_local} onChange={this.OnInputChange} className="form-control" placeholder="Ingrese su nombre de usuario" />
              </div>

              <div className="form-group">
                <label>Visitante</label>
                <input type="text" name="p_visit" value={this.state.info.p_visit} onChange={this.OnInputChange} className="form-control" placeholder="Ingrese su contraseña" />
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
      </div>
    )
  }


  handleWeekendsToggle = () => {
    this.setState({
      weekendsVisible: !this.state.weekendsVisible
    })
  }

  handleDateSelect = (selectInfo) => {
    let title = prompt('Please enter a new title for your event')
    let calendarApi = selectInfo.view.calendar

    calendarApi.unselect() // clear date selection

    if (title) {
      calendarApi.addEvent({
        title,
        start: selectInfo.startStr,
        end: selectInfo.endStr,
        allDay: selectInfo.allDay
      })
    }
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
          p_local: clickInfo.event._def.extendedProps.p_local,
          p_visit: clickInfo.event._def.extendedProps.p_visit
        }
      }))
  }
}

function renderEventContent(eventInfo) {
  return (
    <>
      <b>{eventInfo.timeText}</b>
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
  console.log(event)
  return (

    <i>{event.event.title}</i>

  )
}