
function n3ms_help_local(name) {
    fetch(
        'http://' + location.host + '/help/?service=' + name, {
        method: 'GET',
        body: null,
    }
    ).then((response) => {
        console.log(response.status);
        return response.text();
    }).then((result) => {
        console.log(result);
    }).catch((error) => {
        console.error('Error:', error);
    });    
}