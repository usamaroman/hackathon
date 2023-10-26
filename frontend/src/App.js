import { Provider, useDispatch, useSelector } from 'react-redux';
import './App.css';
import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import { Header } from './components/Header/Header';
import react, { useEffect } from 'react';
import { userActions } from './userState/loginUserSlice';
import { Registration } from './components/Registration/Registration';
import { Login } from './components/Login/Login';
import { Home } from './components/Home/Home';
import { GetProjects } from './components/GetProjects/GetProjects';



function App() {
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
        dispatch(userActions.setIsAuth)
    }
    }, [dispatch]);

  return (
    <BrowserRouter>
      <Routes>
        {role !== "admin" && <>
          <Route path="/" element={isAuth ? <Navigate to={"/home"}/> : <Navigate to={"/login"} />}/>
          <Route path={"/login"} element={isAuth ? <Navigate to={"/home"} /> : <Login />} />
          <Route path={"/registration"} element={isAuth ? <Navigate to={"/home"} /> : <Registration />} />
          <Route path={"/home"} element={isAuth ? <Home/> : <Navigate to={"/login"}/>} />
          <Route path={"/projects"} element={isAuth ? <GetProjects/> : <Navigate to={"/login"}/>} />
        </>
        }
        {role == "admin" && <>
          <Route path="/" element={isAuth ? <Navigate to={"/home"}/> : <Navigate to={"/login"} />}/>
          <Route path={"/login"} element={isAuth ? <Navigate to={"/home"} /> : <Login />} />
          <Route path={"/registration"} element={isAuth ? <Navigate to={"/home"} /> : <Registration />} />
          <Route path={"/home"} element={isAuth ? <Home/> : <Navigate to={"/login"}/>} />
          <Route path={"/projects]"} element={isAuth ? <GetProjects/> : <Navigate to={"/login"}/>} />
        </>
        }
      </Routes>
    </BrowserRouter>
  );
}

export default App;
