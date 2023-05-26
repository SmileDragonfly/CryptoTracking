import logo from './logo.svg';
import './App.css';
import ReactTable from './ReactTable'

function App() {
    let url = "https://5c05-14-177-7-113.ngrok-free.app"
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
            <ReactTable endPoint={url+'/1minup'} caption={'1 Min Up'}/>
            <ReactTable endPoint={url+'/5minup'} caption={'5 Min Up'}/>
            <ReactTable endPoint={url+'/10minup'} caption={'10 Min Up'}/>
            <ReactTable endPoint={url+'/15minup'} caption={'15 Min Up'}/>
            <ReactTable endPoint={url+'/30minup'} caption={'30 Min Up'}/>
            <ReactTable endPoint={url+'/60minup'} caption={'60 Min Up'}/>
        </div>
        <footer> Contact me</footer>
    </div>
  );
}

export default App;
