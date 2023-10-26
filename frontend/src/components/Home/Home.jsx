import react from "react";
import { Header } from "../Header/Header";
import "./home.css";
import { Link, useNavigate } from "react-router-dom";

export const Home = () => {
  const navigate = useNavigate();

  return (
    <div className="main">
      <div className="navbar">
        <div className="icon">
          <h2 className="logo">Hackatone</h2>
        </div>

        <div className="menu">
          <ul>
            <li>
              <Link to="/registration">Проекты</Link>
            </li>
            <li>
              <Link to="/registration">Задачи</Link>
            </li>
            <li>
              <a href="#">Выйти</a>
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
};
