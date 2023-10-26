import React, { useState, useEffect } from "react";
import Modal from "react-modal";
import { Login } from "../Login/Login";
import { Registration } from "../Registration/Registration";
import { NewProject } from "../NewProject/NewProject";
import "./header.css";
import { NewTask } from "../NewTask/NewTask";
import {useDispatch, useSelector} from "react-redux";
import { userActions } from "../../userState/loginUserSlice";
import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";

Modal.setAppElement("#root");

export const Header = () => {
  const [isLoginModalOpen, setLoginModalOpen] = useState(false);
  const [isRegistrationModalOpen, setRegistrationModalOpen] = useState(false);
  const [isNewProjectModalOpen, setNewProjectModalOpen] = useState(false);
  const [isNewTaskOpen, setNewTaskOpen] = useState(false);

  const isAuth = useSelector((state) => state.user.isAuth)
  const role = useSelector((state) => state.user.role)
  const dispatch = useDispatch()

  useEffect(() => {
    const token = JSON.parse(localStorage.getItem('access_token'))
    if (!token) {
        dispatch(userActions.logout)
        } else {
        const user = JSON.parse(localStorage.getItem('user'))
        if (user.is_verified) {
            dispatch(userActions.setIsVerified())
        }
    }
    }, [dispatch]);

  return (
    <>
    </>
  );
};
