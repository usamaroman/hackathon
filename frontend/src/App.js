import { Provider, useDispatch, useSelector } from 'react-redux';
import './App.css';
import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import { Header } from './components/Header/Header';
import react, { useEffect } from 'react';
import { userActions } from './userState/loginUserSlice';
import { Registration } from './components/Registration/Registration';
import { Login } from './components/Login/Login';



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
    }
    }, [dispatch]);

  return (
    <BrowserRouter>
      <Routes>
        {role !== "admin" && <>
          <Route path="/" element={isAuth ? <Navigate to={"/"}/> : <Navigate to={"/login"} />}/>
          <Route path={"/login"} element={isAuth ? <Navigate to={"/"} /> : <Login />} />
          <Route path={"/registration"} element={isAuth ? <Navigate to={"/"} /> : <Registration />} />
        </>
        }
        {role == "admin" && <>
          <Route path="/" element={isAuth ? <Navigate to={"/"}/> : <Navigate to={"/login"} />}/>
          <Route path={"/login"} element={isAuth ? <Navigate to={"/"} /> : <Login />} />
          <Route path={"/registration"} element={isAuth ? <Navigate to={"/"} /> : <Registration />} />
        </>
        }
      </Routes>
    </BrowserRouter>
  );
}

export default App;
