import React, { useState, useEffect } from "react";
import Moment from 'react-moment';


function TopCoinTable({endPoint, caption}){// Define state variables for JSON data and loading status
    const [data, setData] = useState([]);
    const [loading, setLoading] = useState(0);
    let url = endPoint + '/topcoin'

// Use useEffect hook to call fetchData when component mounts
    useEffect( () => {
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
                    setLoading((c)=> c+1);
                } catch (error) {
                    // Handle any errors
                    console.error(error);
                }
                setTimeout(()=>{fetchData(endPoint)}, 5000)
            }
            fetchData(url);
        }
    , [url]);

// Return JSX elements for rendering the table
    return (
        <div className={TopCoinTable}>
            {loading === 0 ? (
// Show a loading message while data is being fetched
                <p>Loading...</p>
            ) : (
// Show the table when data is ready
                <table>
                    <caption>{caption}</caption>
                    <thead>
                    <tr>
                        <th>Time</th>
                    </tr>
                    </thead>
                    <tbody>
                    {data.map((item) => (
// Render each item as a table row with table cells
                        <tr>
                            <td><Moment format={"HH:mm:ss"}>{item.Time}</Moment></td>
                            {item.TopCoin.map((it) => (
                                <td>{it.Symbol.replace("BUSD", "")}</td>
                            ))
                            }
                        </tr>
                    ))}
                    </tbody>
                </table>
            )}
        </div>
    );
}

export default TopCoinTable;