import {Link, useNavigate} from "react-router-dom";
import "./home.css"
import {useSelector} from "react-redux";

export const Navbar = () => {
    const navigate = useNavigate();
    const isAuth = useSelector((state) => state.user.isAuth)

    function logout(){
        console.log("here")
        localStorage.clear()
        navigate("/")
    }

    function login() {
        navigate("/login")
    }

    return (
        <div className="navbar">
            <Link to={"/home"}>
                <h2 className="icon logo">
                    Hackathon
                </h2>
            </Link>
            <div className="menu">
                <div>
                    <Link to="/home">Главная</Link>
                </div>
                <div>
                    <Link to="/projects">Проекты</Link>
                </div>
                {/* <div>
                    <Link to="/tasks">Задачи</Link>
                </div> */}
            </div>
            <div className={"logout"}>
                {isAuth ? <button onClick={logout}>Выйти</button> : <button onClick={login}>логин</button>}
            </div>
        </div>
    )
}