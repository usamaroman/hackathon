import { Provider } from 'react-redux';
import './App.css';
import { Login } from './components/Login/Login';
import { Registration } from './components/Registration/Registration';
import { NewProject } from './components/NewProject/NewProject';
import { Header } from './components/Header/Header';

function App() {
  return (
    <div>
      <Header/>
    </div>
  );
}

export default App;
