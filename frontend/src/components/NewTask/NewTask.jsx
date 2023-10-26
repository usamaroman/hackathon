import React, { useState } from "react";
import axios from "axios";
import "./newtask.css";
import { axiosInstance } from "../../axios/axios";

export const NewTask = () => {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [difficulty, setDifficulty] = useState("");
  const [priority, setPriority] = useState("");
  const [startDate, setStartDate] = useState("");
  const [endDate, setEndDate] = useState("");

  function formatDate(dateString) {
    const date = new Date(dateString);

    const day = date.getDate();
    const month = date.getMonth() + 1;
    const year = date.getFullYear();

    const formattedDay = day < 10 ? `0${day}` : day;
    const formattedMonth = month < 10 ? `0${month}` : month;

    return `${formattedDay}.${formattedMonth}.${year}`;
  }

  const handlePriorityChange = (event) => {
    let value = event.target.value;

    if (value > 100) {
      value = 100;
    }
    setPriority(value);
  };

  async function createTask() {
    console.log({
        title: name,
        description: description,
        difficulty: difficulty,
        priority: priority,
        start: formatDate(startDate),
        end: formatDate(endDate),
      })
    const createProject = await axiosInstance.post("/tasks/create", {
      "title": name,
      "description": description,
      "difficulty": difficulty,
      "priority": priority,
      "start": formatDate(startDate),
      "end": formatDate(endDate),
    });
    
  }

  return (
    <div className="new-task-container">
      <input
        className="new-task-input"
        type="text"
        placeholder={"Название задания"}
        value={name}
        onChange={(event) => setName(event.target.value)}
      />
      <input
        className="new-task-input"
        type="text"
        placeholder={"Описание задания"}
        value={description}
        onChange={(event) => setDescription(event.target.value)}
      />
      <input
        className="new-task-input"
        type="number"
        placeholder={"Приоритет задания"}
        min={1}
        max={100}
        value={priority}
        onChange={handlePriorityChange}
      />
      <select
        className="new-task-input"
        id="difficulty"
        onChange={event => setDifficulty(event.target.value)}
      >
        <option value="" disabled selected hidden>Сложность</option>
        <option value="Низкая">Низкая</option>
        <option value="Средняя">Средняя</option>
        <option value="Высокая">Высокая</option>
      </select>
      <input
        className="new-task-input"
        type="date"
        placeholder={"Дата начала"}
        value={startDate}
        onChange={(event) => setStartDate(event.target.value)}
      />
      <input
        className="new-task-input"
        type="date"
        placeholder={"Дата окончания"}
        value={endDate}
        onChange={(event) => setEndDate(event.target.value)}
      />
      <button className="new-task-button" onClick={createTask}>
        Создать задание
      </button>
    </div>
  );
};
