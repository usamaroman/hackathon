import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import {axiosInstance} from "../../axios/axios";
import "./index.css"

export const TaskPage = () => {
    const params = useParams()
    const [tasks, setTasks] = useState(null)
    const [project, setProject] = useState(null)

    useEffect(() => {
        // fetchTasks()
        fetchInfo()
    }, [])

    // const fetchTasks = async () => {
    //     const res = await axiosInstance.get(`/tasks/${params.id}/tasks`)
    //     setTasks(res.data)
    //     console.log(tasks)
    // }

    const fetchInfo = async () => {
        const res = await axiosInstance.get(`/tasks/${params.id}`)
        setProject(res.data)
        console.log(project)
    }

    return (
        <div className={"project-page"}>
            {project !== null && <div>
                <h1>{project.title}</h1>
                <p>{project.description}</p>
                <div>
                    <div>{project.start}</div>
                    <div>{project.end}</div>
                </div>
            </div>}
            {tasks !== null && tasks.map(t => <div>
                {t.title} {t.start}-{t.end}
            </div>)}
        </div>

    )
}