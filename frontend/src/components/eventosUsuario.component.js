
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
    Events: [],
    R_visita:"/",
    R_local:"/",
    P_visita:"",
    P_local:""
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

  onSubmit = async (e) => {
    e.preventDefault();
    await this.sendPost();

  };
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
  sendPost = async () => {
    const Prediccion = {
      user: this.data.id,
      event:this.state.info.id,
      p_local: this.state.P_local,
      p_visita: this.state.P_visita
    }
    console.log(Prediccion)
    await axios.post("http://localhost:4000/NewPred/", Prediccion)
      .then(response => {
        console.log(response.data)
      });
      this.state.Events=[];
      this.sendGet();
      
  }
  OnInputChange = e => {
    this.setState({
      [e.target.name]: e.target.value,
    });
    console.log(this.state)
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
        <div className='eventos-app-sidebar-section'>
          <div className="card text-center">
            
            <div className="card-body">
            <label className="card-title">Local: {this.state.info.local}<br />Visitante: {this.state.info.visit} </label>
              
             
              <div className="form-group">
                <b>Resultado</b>
                <br/>
                <label>Local: &nbsp;</label>
                <input type="text" name="R_local" value={this.state.R_local}  disabled />
              </div>

              <div className="form-group">
                <label>Visita:  &nbsp;</label>
                <input type="text" name="R_visita" value={this.state.R_visita} disabled />
              </div>


             
              <div className="form-group">
                <h5>Prediccion</h5>
                <label>Local</label>
                <input type="text" name="P_local" value={this.state.P_local} onChange={this.OnInputChange} className="form-control" placeholder="Prediccion de visitante" />
              </div>

              <div className="form-group">
                <label>Visitante</label>
                <input type="text" name="P_visita" value={this.state.P_visita} onChange={this.OnInputChange} className="form-control" placeholder="Prediccion de local" />
              </div>

              <form onSubmit={this.onSubmit}>

                <button type="submit" className="btn btn-primary btn-block">Ingresar</button>

              </form>
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
          p_local: clickInfo.event._def.extendedProps.p_local,
          p_visit: clickInfo.event._def.extendedProps.p_visit
        },
        R_local:clickInfo.event._def.extendedProps.s_local,
        R_visita:clickInfo.event._def.extendedProps.s_visit,
        P_local:clickInfo.event._def.extendedProps.p_local,
        P_visita:clickInfo.event._def.extendedProps.p_visit,
      }),console.log(this.state.info))
  }
}


