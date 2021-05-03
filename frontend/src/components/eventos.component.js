
import React from 'react'
import FullCalendar from '@fullcalendar/react'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin from '@fullcalendar/interaction'
import esLocale from '@fullcalendar/core/locales/es';
import axios from 'axios';

export default class DemoApp extends React.Component {

  state = {
    weekendsVisible: true,
    calendarEvents: [
      { title: "Event Now", start: new Date() }
    ]
    
  }
  componentDidMount() {
    this.sendGet();
  }
  

 
  fixEvent(E_id,local, visit, s_local,s_visit,fechaI) {    
    let Etitle = local + ' - ' + visit+'\n'+s_local + ' - ' + s_visit;
    let Estart = fechaI;
    this.setState({
      calendarEvents: this.state.calendarEvents.concat({
        id:E_id,
        title: Etitle,
        start: Estart
      })
    })
  }

  
  sendGet = async () => {

    const res = await axios
      .get("http://localhost:4000/eventos/")
      .then(response => {
        //localStorage.setItem('events', JSON.stringify(response.data));
        response.data.forEach(element => {
          this.fixEvent(element.id,element.local, element.visita,element.m_local,element.m_visita, element.fecha_inicio);
        });
        console.log(this.state.calendarEvents);
      });
  }


  render() {
    return (
      <div className='demo-app'>
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
            editable={true}
            selectable={true}
            eventDisplay={'block'}
            selectMirror={true}
            dayMaxEvents={true}
            eventSources={this.state.eventos}
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





  handleEventClick = (clickInfo) => {
    clickInfo.event.remove()
    console.log(this.state.currentEvents)
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

