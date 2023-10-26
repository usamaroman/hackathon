import react, { useEffect, useState } from "react"
import { axiosInstance } from "../../axios/axios";
import "./gettasks.css"
import {TaskCard} from "../TaskCard/TaskCard";

export const GetTasks = () =>{
    const [tasks, setTasks] = useState([])

    const fetchTasks = async () => {
        const res = await axiosInstance.get("/tasks")
        setTasks(res.data)
    }

    useEffect(() => {
        fetchTasks()
    }, [])

    return(
        <div className={"tasks"}>
            {tasks.map(pr => <TaskCard pr={pr} />)}
        </div>
    )
}