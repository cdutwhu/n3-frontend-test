// function readable(s) {
//     return s.replace(/\\u003c/gi, "<").replace(/\\u003e/gi, ">").replace(/\\r\\n/g, "\n").replace(/\\n/g, "\n").replace(/\\"/g, "\"");
// }

function init(name) {

    var finput, form, fname;
    var info = document.getElementById('info');

    switch (name) {
        case "privacy":
            finput = document.getElementById('selfile_pri');
            form = document.getElementById("form_pri");
            break;

        case "sif2json":
            finput = document.getElementById('selfile_s2j');
            form = document.getElementById("form_s2j");
            break;

        case "csv2json":
            finput = document.getElementById('selfile_c2j');
            form = document.getElementById("form_c2j");
            break;

        default:
            alert("need correct service name [privacy, sif2json csv2json]");
            return;
    }

    finput.addEventListener('change', function (evt) {
        var files = evt.target.files; // files is a FileList of File objects. List some properties.
        var output = [];
        for (var i = 0, f; f = files[i]; i++) {
            output.push('<li><strong>', escape(f.name), '</strong> (', f.type || 'n/a', ') - ',
                f.size, ' bytes, last modified: ',
                f.lastModifiedDate ? f.lastModifiedDate.toLocaleDateString() : 'n/a',
                '</li>');
        }
        info.innerHTML = ('<ul>' + output.join('') + '</ul>');
        fname = files[0].name;
    });

    form.addEventListener('submit', function (e) {
        e.preventDefault();  // avoid to execute the actual submit of the form

        var ip = location.host;
        var url = '';

        switch (name) {
            case "privacy":
                ip = '192.168.31.168:1323';
                url = 'http://' + ip + '/policy-service/0.1.0/enforce'
                break;

            case "sif2json":
                ip = '192.168.31.168:1324';
                if (fname.endsWith('.xml')) {
                    url = 'http://' + ip + '/sif2json/0.1.0'
                } else if (fname.endsWith('.json')) {
                    url = 'http://' + ip + '/json2sif/0.1.0'
                }
                break;

            case "csv2json":
                ip = '192.168.31.168:1325';
                if (fname.endsWith('.csv')) {
                    url = 'http://' + ip + '/csv2json/0.1.0'
                } else if (fname.endsWith('.json')) {
                    url = 'http://' + ip + '/json2csv/0.1.0'
                }
                break;

            default:
                alert("need correct service name [privacy, sif2json csv2json]");
                return;
        }

        // -------------------------------------------------

        fetch(
            url, {
            method: 'POST',
            body: finput.files[0],
        }
        ).then((response) => {
            console.log(response.status);
            return response.json();
        }).then((result) => {
            console.log('data:', result.data);
            console.log('info:', result.info);
            console.log('empty:', result.empty);
            console.log('error:', result.error);
        }).catch((error) => {
            console.error('Error:', error);
        });

        // -------------------------------------------------

        // var xhr = new XMLHttpRequest;
        // xhr.open('POST', url, true);
        // // xhr.setRequestHeader("Authorization", "Basic " + btoa("user" + ":" + "user"));
        // xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest');
        // xhr.setRequestHeader('X-File-Name', file.name);
        // xhr.setRequestHeader('Content-Type', file.type || 'application/octet-stream');
        // xhr.onreadystatechange = function () {
        //     if (xhr.readyState === 4 && xhr.status === 200) { // 0:UNSENT  1:OPENED  2:HEADERS_RECEIVED  3:LOADING  4:DONE
        //         var response = xhr.responseText;
        //         response = readable(response);
        //         console.log(response);
        //     }
        // };
        // xhr.send(file);

        // -------------------------------------------------
        
        // !!! with "*webkitformboundary*" in the sending body !!! //

        // var formdata = new FormData();
        // formdata.append("file", file);

        // $.ajax({ // make an AJAX request
        //     type: 'POST',
        //     method: 'POST',
        //     url: url,
        //     data: formdata, // $("#form").serialize(), // serializes the form's elements
        //     cache: false,
        //     contentType: false,
        //     processData: false,
        //     crossDomain: true,
        //     beforeSend: function (xhr) {
        //         console.log("beforeSend");
        //         // console.log(xhr);
        //         // xhr.setRequestHeader("Authorization", "Basic " + btoa("user" + ":" + "user"));
        //         $('#selectfile').prop('disabled', true);
        //         $('#pub').prop('disabled', true);
        //         $('#waiting').show();
        //     },
        //     success: function (data) {
        //         console.log("success");
        //         console.log(data);
        //         $('#selectfile').val('');
        //     },
        //     error: function (jqXHR, textStatus, errorThrown) {
        //         console.log("error");
        //         console.log(jqXHR.responseText);
        //         $('#pub').prop('disabled', false);
        //     },
        //     complete: function () {
        //         console.log("complete");
        //         $('#selectfile').prop('disabled', false);
        //         $('#waiting').hide();
        //     }
        // });

        // -------------------------------------------------
    });
}

window.onload = function () {
    init("privacy");
    init("sif2json");
    init("csv2json");
}
