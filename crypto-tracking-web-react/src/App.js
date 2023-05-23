import logo from './logo.svg';
import './App.css';
import ReactTable from './ReactTable'

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        {/*<a*/}
        {/*  className="App-link"*/}
        {/*  href="https://reactjs.org"*/}
        {/*  target="_blank"*/}
        {/*  rel="noopener noreferrer"*/}
        {/*>*/}
        {/*  Learn React*/}
        {/*</a>*/}
          <div class="btn-group">
              <button>All</button>
              <button>1m-5m-10m</button>
              <button>15m-30m-60m</button>
          </div>
      </header>
        <div className="grid-container">
            <ReactTable endPoint={'https://0469-123-24-56-232.ngrok-free.app/1minup'}/>
        </div>
        <footer> Contact me</footer>
    </div>
  );
}

export default App;
