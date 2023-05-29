import ReactTable from "./ReactTable";

function DownTables({endPoint}){
    return(
        <div className="grid-container">
            <ReactTable endPoint={endPoint+'/1mindown'} caption={'1 Min Down'}/>
            <ReactTable endPoint={endPoint+'/5mindown'} caption={'5 Min Down'}/>
            <ReactTable endPoint={endPoint+'/10mindown'} caption={'10 Min Down'}/>
            <ReactTable endPoint={endPoint+'/15mindown'} caption={'15 Min Down'}/>
            <ReactTable endPoint={endPoint+'/30mindown'} caption={'30 Min Down'}/>
            <ReactTable endPoint={endPoint+'/60mindown'} caption={'60 Min Down'}/>
        </div>
    )
}

export default DownTables;