// function readable(s) {
//     return s.replace(/\\u003c/gi, "<").replace(/\\u003e/gi, ">").replace(/\\r\\n/g, "\n").replace(/\\n/g, "\n").replace(/\\"/g, "\"");
// }

function init(name) {

    var finput, info, btnSend, uploadform, fname;

    switch (name) {
        case "privacy":
            finput = document.getElementById('selectfile0');
            info = document.getElementById('info0');
            btnSend = document.getElementById('pub0');
            uploadform = document.getElementById("uploadform0");
            break;

        case "sif2json":
            finput = document.getElementById('selectfile1');
            info = document.getElementById('info1');
            btnSend = document.getElementById('pub1');
            uploadform = document.getElementById("uploadform1");
            break;

        case "csv2json":
            finput = document.getElementById('selectfile2');
            info = document.getElementById('info2');
            btnSend = document.getElementById('pub2');
            uploadform = document.getElementById("uploadform2");
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
        var disabled = !fname;
        switch (name) {
            case 'privacy':
                disabled |= (!fname.endsWith('.json'))
                break;
            case 'sif2json':
                disabled |= (!fname.endsWith('.xml') && !fname.endsWith('.json'))
                break;
            case 'csv2json':
                disabled |= (!fname.endsWith('.csv'))
                break;
        }
        btnSend.disabled = disabled;
    });

    uploadform.addEventListener('submit', function (e) {
        e.preventDefault();  // avoid to execute the actual submit of the form

        var ip = location.host;
        var url = '';
        var finput;

        switch (name) {
            case "privacy":
                ip = '192.168.31.233:1323';
                finput = document.getElementById('selectfile0')
                break;

            case "sif2json":
                ip = '192.168.31.233:1324';
                finput = document.getElementById('selectfile1')
                if (fname.endsWith('.xml')) {
                    url = 'http://' + ip + '/sif2json/0.1.0'
                } else if (fname.endsWith('.json')) {
                    url = 'http://' + ip + '/json2sif/0.1.0'
                }
                break;

            case "csv2json":
                ip = '192.168.31.233:1325';
                finput = document.getElementById('selectfile2')
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
        //     data: formdata, // $("#uploadform").serialize(), // serializes the form's elements
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
