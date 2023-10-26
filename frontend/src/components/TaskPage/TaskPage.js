import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import {axiosInstance} from "../../axios/axios";
import "./index.css"

export const TaskPage = () => {
    const params = useParams()
    const [comments, setComments] = useState(null)
    const [task, setTask] = useState(null)

    useEffect(() => {
        fetchComments()
        fetchInfo()
    }, [])

    const fetchComments = async () => {
        const res = await axiosInstance.get(`/tasks/${params.id}/comments`)
        setComments(res.data)
        console.log(comments)
    }

    const fetchInfo = async () => {
        const res = await axiosInstance.get(`/tasks/${params.id}`)
        setTask(res.data)
        console.log(task)
    }

    return (
        <div className={"project-page"}>
            {task !== null && <div className={"project-inner"}>
                <div >
                    <h1>{task.title}</h1>
                    <p>{task.description}</p>
                    <div>
                        <div>{task.start}</div>
                        <div>{task.end}</div>
                    </div>
                </div>
                <div className={"dop-info"}>
                    <div>Приоритетность: {task.priority}</div>
                    <div>Сложность: {task.difficulty}</div>
                    <div>Статус: {task.status}</div>
                </div>
            </div>}
            <hr/>
            <h2>Комментарии</h2>
            {comments !== null && comments.map(c =>
                <div>{c.text}</div>
            )}
        </div>

    )
}