import { useDispatch, useSelector } from 'react-redux';
import './App.css';
import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import react, { useEffect } from 'react';
import { userActions } from './userState/loginUserSlice';
import { Registration } from './components/Registration/Registration';
import { Login } from './components/Login/Login';
import { Home } from './components/Home/Home';
import { GetProjects } from './components/GetProjects/GetProjects';
import {Navbar} from "./components/Navbar/Navbar";
import {GetTasks} from "./components/GetTasks/GetTasks";
import {ProjectPage} from "./components/ProjectPage/ProjectPage";
import {TaskPage} from "./components/TaskPage/TaskPage";

export const App = () => {
      const isAuth = useSelector((state) => state.user.isAuth)
      const role = useSelector((state) => state.user.role)
      const dispatch = useDispatch()

      useEffect(() => {
            const token = JSON.parse(localStorage.getItem('access_token'))
            if (!token) {
                dispatch(userActions.logout)
            } else {
                const user = JSON.parse(localStorage.getItem('user'))
                dispatch(userActions.setIsAuth())
                dispatch(userActions.setUser(user))
            }
      }, [dispatch]);

  return (
      <div className={"app"}>
    <BrowserRouter>
      <Navbar />
      <Routes>
        {role !== "admin" && <>
          <Route path="*" element={isAuth ? <Navigate to={"/home"}/> : <Navigate to={"/login"} />}/>
          <Route path="/" element={isAuth ? <Navigate to={"/home"}/> : <Navigate to={"/login"} />}/>
          <Route path={"/login"} element={isAuth ? <Navigate to={"/home"} /> : <Login />} />
          <Route path={"/registration"} element={isAuth ? <Navigate to={"/home"} /> : <Registration />} />
          <Route path={"/home"} element={isAuth ? <Home/> : <Navigate to={"/login"}/>} />
          <Route path={"/projects"} element={isAuth ? <GetProjects/> : <Navigate to={"/login"}/>} />
          <Route path={"/projects/:id"} element={isAuth ? <ProjectPage /> : <Navigate to={"/login"} />}/>
          <Route path={"/tasks"} element={isAuth ? <GetTasks /> : <Navigate to={"/login"}/>} />
          <Route path={"/tasks/:id"} element={isAuth ? <TaskPage /> : <Navigate to={"/login"} />}/>
        </>
        }
        {role === "admin" && <>
          <Route path="/" element={isAuth ? <Navigate to={"/home"}/> : <Navigate to={"/login"} />}/>
          <Route path={"/login"} element={isAuth ? <Navigate to={"/home"} /> : <Login />} />
          <Route path={"/registration"} element={isAuth ? <Navigate to={"/home"} /> : <Registration />} />
          <Route path={"/home"} element={isAuth ? <Home/> : <Navigate to={"/login"}/>} />
          <Route path={"/projects]"} element={isAuth ? <GetProjects/> : <Navigate to={"/login"}/>} />
          <Route path={"/projects/:id"} element={<ProjectPage />}/>
          <Route path={"/tasks/:id"} element={<TaskPage />}/>
            {/*<Route path={"/projects"} element={isAuth ? <GetProjects/> : <Navigate to={"/login"}/>} >*/}
            {/*    <Route path={":id"} element={<ProjectPage />}/>*/}
            {/*</Route>*/}
          {/*<Route path={"/tasks"} element={isAuth ? <GetTasks /> : <Navigate to={"/login"}/>} />*/}
        </>
        }
      </Routes>
    </BrowserRouter>
      </div>
  );
}

export default App;
