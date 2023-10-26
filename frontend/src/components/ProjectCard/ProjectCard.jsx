import "./getprojects.css"
import {Link} from "react-router-dom";

export const ProjectCard = (props) => {
    const {id, title, description, start, end} = props.pr

    return (
        <Link to={`/projects/${id}`} className="card">
            <div className="info">
                <div className="title">
                    <div>
                        <h1 className="big">{title}</h1>
                    </div>
                </div>
                <div className="description">
                    <p className="text">{description}</p>
                </div>
                <div className="div-dates">
                    <div className="dates">
                        <h5>{start} - {end}</h5>
                    </div>
                </div>
            </div>
        </Link>

    )
}