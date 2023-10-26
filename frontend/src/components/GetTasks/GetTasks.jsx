import react, { useEffect, useState } from "react"
import { axiosInstance } from "../../axios/axios";

export const GetTasks = () =>{
    const [tasks, setTasks] = useState([])

    const fetchTasks = async () => {
        const getTasks = await axiosInstance.get("/tasks")
        setTasks(getTasks.data)
    }

    useEffect(() => {
        fetchTasks()
    }, [])

    return(
        <div>
            {tasks.map(pr => <div>{pr.title}</div>)}
        </div>
    )
}