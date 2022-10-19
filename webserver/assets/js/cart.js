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




function AddCartItem(uuid) {
    var data = {
        uuid: uuid,
    };

    var json = JSON.stringify(data);

    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/cart/additem");
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(json);
}



