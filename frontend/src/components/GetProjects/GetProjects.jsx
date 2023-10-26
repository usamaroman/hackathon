import react, { useEffect, useState } from "react"
import { axiosInstance } from "../../axios/axios";
import { ProjectCard } from "../ProjectCard/ProjectCard";
import "./index.css"

export const GetProjects = () =>{
    const [projects, setProjects] = useState([])

    useEffect(() => {
        fetchProjects()
    }, [])

    const fetchProjects = async () => {
        const res = await axiosInstance.get("/projects")
        setProjects(res.data)
    }
   
    return(
        <div className={"projects"} >
            {projects !== undefined && projects.map(pr => <ProjectCard key={pr.id} pr={pr} />)}
        </div>
        
    )
}