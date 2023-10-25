import React, { useState } from 'react';
import axios from "axios";
import "./newproject.css"

export const NewProject = () => {
    const [name, setName] = useState("");
    const [description, setDescription] = useState("");
    const [startDate, setStartDate] = useState("");
    const [endDate, setEndDate] = useState("");

    async function createProject() {
        // const createProject = await axios.post(url, { name: name, description: description, startDate: startDate, endDate: endDate });
        console.log({ name, description, startDate, endDate });
    }

    return (
        <div className="new-project-container">
            <input
                className="new-project-input"
                type="text"
                placeholder={"Название проекта"}
                value={name}
                onChange={event => setName(event.target.value)}
            />
            <input
                className="new-project-input"
                type="text"
                placeholder={"Описание проекта"}
                value={description}
                onChange={event => setDescription(event.target.value)}
            />
            <input
                className="new-project-input"
                type="date"
                placeholder={"Дата начала"}
                value={startDate}
                onChange={event => setStartDate(event.target.value)}
            />
            <input
                className="new-project-input"
                type="date"
                placeholder={"Дата окончания"}
                value={endDate}
                onChange={event => setEndDate(event.target.value)}
            />
            <button
                className="new-project-button"
                onClick={createProject}
            >
                Создать проект
            </button>
        </div>
    )
}
