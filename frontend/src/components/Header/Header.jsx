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
import { GetProjects } from "../GetProjects/GetProjects";
import { GetTasks } from "../GetTasks/GetTasks";

Modal.setAppElement("#root");

export const Header = () => {
  const [isLoginModalOpen, setLoginModalOpen] = useState(false);
  const [isRegistrationModalOpen, setRegistrationModalOpen] = useState(false);
  const [isNewProjectModalOpen, setNewProjectModalOpen] = useState(false);
  const [isNewTaskOpen, setNewTaskOpen] = useState(false);
  const [isGetProjectsOpen, setGetProjectsOpen] = useState(false);
  const [isGetTasksOpen, setGetTasksOpen] = useState(false);

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
    <div className="header">
      <div>
        <button
          className="header-button"
          onClick={() => setLoginModalOpen(true)}
        >
          Логин
        </button>
        <button
          className="header-button"
          onClick={() => setRegistrationModalOpen(true)}
        >
          Регистрация
        </button>
        <button
          className="header-button"
          onClick={() => setNewProjectModalOpen(true)}
        >
          Создать проект
        </button>

        <button
          className="header-button"
          onClick={() => setNewTaskOpen(true)}
        >
          Создать задание
        </button>
        <button
          className="header-button"
          onClick={() => setGetProjectsOpen(true)}
        >
          Список проектов
        </button>

        <button
          className="header-button"
          onClick={() => setGetTasksOpen(true)}
        >
          Список заданий
        </button>
      </div>

      <Modal
        isOpen={isLoginModalOpen}
        onRequestClose={() => setLoginModalOpen(false)}
        className="modal-container"
      >
        <Login />

      </Modal>

      <Modal
        isOpen={isRegistrationModalOpen}
        onRequestClose={() => setRegistrationModalOpen(false)}
        className="modal-container"
      >
        <Registration />
      </Modal>

      <Modal
        isOpen={isNewProjectModalOpen}
        onRequestClose={() => setNewProjectModalOpen(false)}
        className="modal-container"
      >
        <NewProject />
      </Modal>

      <Modal
        isOpen={isNewTaskOpen}
        onRequestClose={() => setNewTaskOpen(false)}
        className="modal-container"
      >
        <NewTask />

      </Modal>

      <Modal
        isOpen={isGetProjectsOpen}
        onRequestClose={() => setGetProjectsOpen(false)}
        className="modal-container"
      >
        <GetProjects />

      </Modal>

      <Modal
        isOpen={isGetTasksOpen}
        onRequestClose={() => setGetTasksOpen(false)}
        className="modal-container"
      >
        <GetTasks />

      </Modal>
    </div>
  );
};
