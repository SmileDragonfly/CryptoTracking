import logo from './logo.svg';
import './App.css';
import ReactTable from './ReactTable'

function App() {
    let url = "https://6d16-118-70-124-128.ngrok-free.app"
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
            <ReactTable endPoint={url+'/1minup'}/>
            <ReactTable endPoint={url+'/5minup'}/>
            <ReactTable endPoint={url+'/10minup'}/>
            <ReactTable endPoint={url+'/15minup'}/>
            <ReactTable endPoint={url+'/30minup'}/>
            <ReactTable endPoint={url+'/60minup'}/>
        </div>
        <footer> Contact me</footer>
    </div>
  );
}

export default App;
