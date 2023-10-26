import react from "react";
import { Header } from "../Header/Header";
import "./home.css";
import { Link, useNavigate } from "react-router-dom";
import { GetProjects } from "../GetProjects/GetProjects";

export const Home = () => {
  const navigate = useNavigate();

  function logout(){
    console.log("here")
    localStorage.clear()
    navigate("/")
  }

  return (
    <div className="main">
      <div className="navbar">
        <div className="icon">
          <h2 className="logo">Hackatone</h2>
        </div>

        <div className="menu">
          <ul>
            <li>
              <Link to="/projects">Проекты</Link>
            </li>
            <li>
              <Link to="/registration">Задачи</Link>
            </li>
            <li>
              <a onClick={logout}>Выйти</a>
            </li>
          </ul>
        </div>
        {/* <GetProjects /> */}
      </div>
    </div>
  );
};
