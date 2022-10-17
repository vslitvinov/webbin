async function Request(method, url, data) {
    return fetch(url, {
        method: method,
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data),
    })
}




function AddCartItem(id) {
    var data = {
        name: "helloworld",
        age: 123
    };

    var json = JSON.stringify(data);

    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/cart");
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(json);
}