import React, {useState} from 'react';
import "./registration.css"
import {useNavigate} from "react-router-dom";

// "city": "string",
// "email": "string",
// "full_name": "string",
// "password": "string",
// "telephone_number": "string"

export const Registration = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    // const navigate = useNavigate();


    const registration = async () => {
        try {
            // await axiosInstance.post("/auth/registration", JSON.stringify(
            const res = {
                    "email": email,
                    "password": password,
                }
            

            console.log(res)

            // navigate("/login")
        } catch (e) {
            console.log(e)
        }
    }

    return (
        <div className={"registration"}>
            <h1 style={{textAlign:"center"}}    >Регистрация</h1>
            <div className={"registration_form"}>
                <input type="text" placeholder={"email"} value={email} onChange={event => setEmail(event.target.value)} />
                <input type="text" placeholder={"password"} value={password} onChange={event => setPassword(event.target.value)} />
                <button onClick={registration}>ok</button>
            </div>
        </div>
    )
};