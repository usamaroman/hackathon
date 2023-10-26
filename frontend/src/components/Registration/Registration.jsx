import React, {useState} from 'react';
import "./registration.css"
import {Link, useNavigate} from "react-router-dom";
import { axiosInstance } from '../../axios/axios';
import image from "./image.jpg"


export const Registration = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();


    const registration = async () => {
        try {
            const res = await axiosInstance.post("/auth/registration", JSON.stringify({
                    "email": email,
                    "password": password,
                }
            ))

            console.log(res)

            navigate("/login")
        } catch (e) {
            console.log(e)
        }

    }

    return (
        <div className="login">
         <img src={image} alt="image" className="login__bg"/>

         <div className="login__form">
            <h1 className="login__title">Регистрация</h1>

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

            <button className="login__button" onClick={registration}>Зарегестрироваться</button>

            <div className="login__register">
               Есть аккаунт? <Link to="/login">Войти</Link>
            </div>
         </div>
      </div>
    )
};