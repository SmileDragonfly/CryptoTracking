import logo from './logo.svg';
import './App.css';
import UpTables from './components/UpTables'
import DownTables from "./components/DownTables";
import {useState} from "react";

function App() {
    let endPoint = "https://f666-118-70-124-128.ngrok-free.app"
    // let url = "http://localhost:8888"
    // create a state variable to store the name of the component to display
    const [component, setComponent] = useState("Up");
    // create a handler function to update the state variable
    const handleClick = (name) => {
        setComponent(name);
    };
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <div class="btn-group">
          <button onClick={() => handleClick("Up")}>Up</button>
          <button onClick={() => handleClick("Down")}>Down</button>
          <button onClick={() => handleClick("History")}>History</button>
        </div>
        <p>Current: {component}</p>
      </header>
        <DisplayComponent name={component} endpoint={endPoint}/>
    </div>
  );
}

// create a functional component that takes the name of the component as a prop
function DisplayComponent({ name, endpoint }) {
// create an object that maps the name to the component
    const components = {
        Up: <UpTables endPoint={endpoint}/>,
        Down: <DownTables endPoint={endpoint}/>,
        History: <UpTables endPoint={endpoint}/>,
    };

// return the component that matches the name or null if none
    return components[name] || null;
}

export default App;
