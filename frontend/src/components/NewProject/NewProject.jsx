import React, { useState } from 'react';
import axios from "axios";
import "./newproject.css"
import { axiosInstance } from '../../axios/axios';

export const NewProject = (props) => {
    const [name, setName] = useState("");
    const [description, setDescription] = useState("");
    const [startDate, setStartDate] = useState("");
    const [endDate, setEndDate] = useState("");
    const {setIsProject} = props

    function formatDate(dateString) {
        const date = new Date(dateString);
    
        const day = date.getDate();
        const month = date.getMonth() + 1; 
        const year = date.getFullYear();
      
        const formattedDay = day < 10 ? `0${day}` : day;
        const formattedMonth = month < 10 ? `0${month}` : month;
      
        return `${formattedDay}.${formattedMonth}.${year}`;
    }

    async function createProject() {
        try {
            const createProject = await axiosInstance.post("/projects", JSON.stringify(
                { 
                    "title": name, 
                    "description": description, 
                    "start": formatDate(startDate), 
                    "end": formatDate(endDate) 
                }
            ))
            console.log(localStorage.getItem("user"))
            setIsProject(false)
        } catch(e) {
            console.log(e)
        }
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
