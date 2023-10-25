import React, { useState } from 'react';
import Modal from 'react-modal';
import { Login } from '../Login/Login';
import { Registration } from '../Registration/Registration';
import { NewProject } from '../NewProject/NewProject';
import './header.css';

Modal.setAppElement('#root'); 

export const Header = () => {
    const [isLoginModalOpen, setLoginModalOpen] = useState(false);
    const [isRegistrationModalOpen, setRegistrationModalOpen] = useState(false);
    const [isNewProjectModalOpen, setNewProjectModalOpen] = useState(false);

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
            </div>

            <Modal
                isOpen={isLoginModalOpen}
                onRequestClose={() => setLoginModalOpen(false)}
                className="modal-container"
            >
                <Login />
                <button
                    className="modal-button"
                    onClick={() => setLoginModalOpen(false)}
                >
                    Закрыть
                </button>
            </Modal>

            <Modal
                isOpen={isRegistrationModalOpen}
                onRequestClose={() => setRegistrationModalOpen(false)}
                className="modal-container"
            >
                <Registration />
                <button
                    className="modal-button"
                    onClick={() => setRegistrationModalOpen(false)}
                >
                    Закрыть
                </button>
            </Modal>

            <Modal
                isOpen={isNewProjectModalOpen}
                onRequestClose={() => setNewProjectModalOpen(false)}
                className="modal-container"
            >
                <NewProject />
                <button
                    className="modal-button"
                    onClick={() => setNewProjectModalOpen(false)}
                >
                    Закрыть
                </button>
            </Modal>
        </div>
    );
}
