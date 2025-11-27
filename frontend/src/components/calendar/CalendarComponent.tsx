import bootstrap5Plugin from '@fullcalendar/bootstrap5';
import allLocales from '@fullcalendar/core/locales-all';
import dayGridPlugin from '@fullcalendar/daygrid';
import interaction from '@fullcalendar/interaction';
import FullCalendar from '@fullcalendar/react';
import timeGridPlugin from '@fullcalendar/timegrid';
import 'bootstrap-icons/font/bootstrap-icons.css';
import 'bootstrap/dist/css/bootstrap.css';

export function CalendarComponent() {
  return (
    <FullCalendar
      // https://fullcalendar.io/docs
      plugins={[bootstrap5Plugin, timeGridPlugin, dayGridPlugin, interaction]}
      // theme
      themeSystem='bootstrap5'
      // Toolbar
      initialView='timeGridWeek'
      headerToolbar={{
        start: 'prev,next today',
        center: 'title',
        end: 'timeGridDay,timeGridWeek,dayGridMonth,dayGridYear',
      }}
      titleFormat={{ year: 'numeric', month: '2-digit', day: '2-digit' }}
      // Sizing
      height={'100%'}
      // Date & Time Display
      // TODO: Add a toggle for this in the settings
      weekends={true}
      dayHeaderFormat={{
        weekday: 'long',
        month: '2-digit',
        day: '2-digit',
        omitCommas: true,
      }}
      // Date Nav Links
      navLinks={true}
      // Week Numbers
      weekNumbers={true}
      // Locale
      locales={allLocales}
      // TODO: Add a dropdown for this option in the settings
      locale={'de'}
      // Date Clicking & Selecting
      selectable={true}
      selectMirror={true}
      // TODO: implement functionality to create events
      select={(dateSelectArg) => { dateSelectArg }}
      // Now Indicator
      nowIndicator={true}
      // Business Hours
      businessHours={{
        daysOfWeek: [1, 2, 3, 4, 5],
        startTime: '07:00',
        endTime: '17:00',
      }}
      events={[]}
    />
  )
}
