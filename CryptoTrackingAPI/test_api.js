async function Get1MinUp(){
    const response = await fetch('https://ed1e-118-70-124-128.ngrok-free.app/1minup');
    // const response = await fetch('https://jsonplaceholder.typicode.com/users');
    const jsonData = await response.json();
    return jsonData
}

console.log(Get1MinUp().then(function (result) {console.log(result)}))