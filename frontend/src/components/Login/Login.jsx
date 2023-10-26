import React, {useState} from 'react';
import "./login.css"
import {Link, useNavigate} from "react-router-dom";
import {userActions} from "../../userState/loginUserSlice"
import {useDispatch} from "react-redux";
import { axiosInstance } from '../../axios/axios';
import image from "./login-bg.png"

export const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const dispatch = useDispatch();
    const navigate = useNavigate();

    const login = async () => {
        console.log(email, password);
        try {
            const res = await axiosInstance.post("/auth/login", JSON.stringify(
                {
                    "email": email,
                    "password": password,
                }
            ))

            localStorage.setItem('access_token', JSON.stringify(res.data.access_token))
            localStorage.setItem('refresh_token', JSON.stringify(res.data.refresh_token))
            localStorage.setItem('user', JSON.stringify(res.data.user))

            dispatch(userActions.setUser(res.data.user))
            dispatch(userActions.setIsAuth())

            navigate("/")
        } catch (e) {
            console.log(e)
        }
    }

    return (
        <div className="login">
         <img src={image} alt="image" className="login__bg"/>

         <div className="login__form">
            <h1 className="login__title">Войти</h1>

            <div className="login__inputs">
               <div className="login__box">
                  <input type="text" placeholder="Электронная почта" className="login__input" onChange={event => setEmail(event.target.value)} value={email}/>
                  <i className="ri-mail-fill"></i>
               </div>

               <div className="login__box">
                  <input type="password" placeholder="Пароль" className="login__input" onChange={event => setPassword(event.target.value)} value={password}/>
                  <i className="ri-lock-2-fill"></i>
               </div>
            </div>

            <button className="login__button" onClick={login}>Войти</button>

            <div className="login__register">
               Нет аккаунта? <Link to="/registration">Зарегестрироваться</Link>
            </div>
         </div>
      </div>

    )
};