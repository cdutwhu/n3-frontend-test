// https://www.html5rocks.com/en/tutorials/file/dndfiles/

// Check for the various File API support.
// if (window.File && window.FileReader && window.FileList && window.Blob) {
//     alert('Great success! All the File APIs are supported.')
// } else {
//     alert('The File APIs are not fully supported in this browser.');
// }

function n3ms_help(name) {

    var ip = location.host;
    switch (name) {
        case "privacy":
            ip = '192.168.31.233:1323';
            break;
        case "sif2json":
            ip = '192.168.31.233:1324';
            break;
        case "csv2json":
            ip = '192.168.31.233:1325';
            break;
        default:
            alert("need correct service name [privacy, sif2json csv2json]");
            return;
    }

    fetch(
        'http://' + ip, {
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

    // -------------------------------------------------

    // var xhr = new XMLHttpRequest;
    // xhr.open('GET', url, true);
    // xhr.onreadystatechange = function () {
    //     if (xhr.readyState === 4 && xhr.status === 200) { // 0:UNSENT  1:OPENED  2:HEADERS_RECEIVED  3:LOADING  4:DONE
    //         var response = xhr.responseText;
    //         console.log(response);
    //     }
    // };
    // xhr.send(null);

    // -------------------------------------------------

    // $.ajax({
    //     url: url,
    //     type: 'GET',
    //     contentType: 'application/json',
    //     data: '',
    //     dataType: 'json',
    //     crossDomain: true,
    //     beforeSend: function (xhr) {
    //         console.log("beforeSend");
    //         console.log(xhr);
    //         // xhr.setRequestHeader("Authorization", "Basic " + btoa("user" + ":" + "user"));
    //     },
    //     success: function (data) {
    //         console.log("success");
    //         console.log(data);
    //     },
    //     error: function (jqXHR, textStatus, errorThrown) {
    //         console.log("error: " + textStatus + " " + errorThrown);
    //         console.log(jqXHR.responseText);
    //         console.log(jqXHR.status);
    //     },
    //     complete: function () {
    //         console.log("complete");
    //     }
    // });
}