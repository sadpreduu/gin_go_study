import React, { useState } from 'react';
import axios from 'axios';

const Reminder = () => {
  const [reminder, setReminder] = useState({
    title: '',
    time: '',
  });

  const handleChange = (event) => {
    setReminder({
      ...reminder,
      [event.target.name]: event.target.value,
    });
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    axios.post('/api/reminders', reminder)
      .then(res => console.log(res.data))
      .catch(err => console.log(err));
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          name="title"
          placeholder="Title"
          value={reminder.title}
          onChange={handleChange}
        />
        <input
          type="time"
          name="time"
          value={reminder.time}
          onChange={handleChange}
        />
        <button type="submit">Add Reminder</button>
      </form>
    </div>
  );
};

export default Reminder;
