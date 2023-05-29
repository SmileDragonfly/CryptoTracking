import ReactTable from "./ReactTable";

function UpTables({endPoint}){
    return(
        <div className="grid-container">
            <ReactTable endPoint={endPoint+'/1minup'} caption={'1 Min Up'}/>
            <ReactTable endPoint={endPoint+'/5minup'} caption={'5 Min Up'}/>
            <ReactTable endPoint={endPoint+'/10minup'} caption={'10 Min Up'}/>
            <ReactTable endPoint={endPoint+'/15minup'} caption={'15 Min Up'}/>
            <ReactTable endPoint={endPoint+'/30minup'} caption={'30 Min Up'}/>
            <ReactTable endPoint={endPoint+'/60minup'} caption={'60 Min Up'}/>
        </div>
    )
}

export default UpTables;