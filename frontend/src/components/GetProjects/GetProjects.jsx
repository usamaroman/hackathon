import react, { useEffect, useState } from "react"
import { axiosInstance } from "../../axios/axios";
import { ProjectCard } from "../ProjectCCArd/ProjectCard";

export const GetProjects = () =>{
    const [projects, setProjects] = useState([])

    useEffect(() => {
        fetchProjects()
    }, [])

    const fetchProjects = async () => {
        const res = await axiosInstance.get("/projects")
        setProjects(res.data)
        console.log(projects)
    }
   
    return(
        <div style={{display: "flex"}}>
            {projects !== undefined && projects.map(pr => <ProjectCard pr={pr} />)}
        </div>
        
    )
}