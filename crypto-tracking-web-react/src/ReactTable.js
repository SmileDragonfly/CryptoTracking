import React, { useState, useEffect } from "react";

function ReactTable({endPoint}){// Define state variables for JSON data and loading status
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(true);
    const [reload, setReload] = useState(0);

// Define an async function to fetch data from API
    async function fetchData(endPoint) {
        try {
            let headers = new Headers()
            headers.append('ngrok-skip-browser-warning', true)
            // headers.append('Access-Control-Allow-Origin', '*')
// Make a GET request and get the response
            const response = await fetch(endPoint, {headers: headers});
// Convert the response to JSON format
            const json = await response.json();
// Update the data state with the JSON data
            setData(json);
// Set the loading state to false
            setLoading(false);
        } catch (error) {
// Handle any errors
            console.error(error);
        }
    }

// Use useEffect hook to call fetchData when component mounts
    useEffect(() => {
        fetchData(endPoint);
    }, [reload, endPoint]);

// Return JSX elements for rendering the table
    return (
        <div>
            <p>Loading {String(reload)}</p>
            <button onClick={() =>setReload((c)=>c+1)}>Reload</button>
            {loading ? (
// Show a loading message while data is being fetched
                <p>Loading...</p>
            ) : (
// Show the table when data is ready
                <table>
                    <thead>
                    <tr>
                        <th>Time</th>
                        <th>Symbol</th>
                        <th>Price</th>
                        <th>PrevPrice</th>
                        <th>Percent</th>
                    </tr>
                    </thead>
                    <tbody>
                    {data.map((item) => (
// Render each item as a table row with table cells
                        <tr key={item.ID}>
                            <td>{item.Time}</td>
                            <td>{item.Symbol}</td>
                            <td>{item.Price}</td>
                            <td>{item.PrevPrice}</td>
                            <td>{item.Percent}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>
            )}
        </div>
    );
}

export default ReactTable;