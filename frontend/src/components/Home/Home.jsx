import {useState} from "react";
import {NewProject} from "../NewProject/NewProject";
import {NewTask} from "../NewTask/NewTask";

export const Home = () => {
    const [isTask, setIsTask] = useState(false)
    const [isProject, setIsProject] = useState(false)

    return (
        <div className={"home-inner"}>
            {isProject ? <div className={"new-form"}>
                <NewProject />
                <button onClick={() => setIsProject(false)}>закрыть</button>
            </div> : <div className={"label"} onClick={() => setIsProject(true)}>создать проект</div>}
            {isTask ? <div className={"new-form"}>
                <NewTask />
                <button onClick={() => setIsTask(false)}>закрыть</button>
            </div> : <div className={"label"} onClick={() => setIsTask(true)}>создать таску</div>}
        </div>
    )
};
