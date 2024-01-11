import logo from './logo.svg';
import './App.css';
import UserDetail from './components/UserDetails';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <UserDetail>
          
        </UserDetail>
      </header>
    </div>
  );
}

export default App;
