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


    xhr.onload = function() {
        console.log(xhr)
        if (xhr.status != 200) { // анализируем HTTP-статус ответа, если статус не 200, то произошла ошибка
            alert(`Ошибка ${xhr.status}: ${xhr.statusText}`); // Например, 404: Not Found
        } else { // если всё прошло гладко, выводим результат
            alert(`Готово, получили ${xhr.response.length} байт`); // response -- это ответ сервера
        }
    };

}