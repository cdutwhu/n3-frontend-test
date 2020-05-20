var filename;

function txt2xml(s) {
    return s.replace(/\\u003c/gi, "<").replace(/\\u003e/gi, ">").replace(/\\r\\n/g, "\n").replace(/\\n/g, "\n");
}

function select_change(name) {

    var finput, info, btnSend;
    switch (name) {
        case "sif2json":
            finput = $('#selectfile0');
            info = $('#info0');
            btnSend = $('#pub0');
            break;

        case "privacy":
            finput = $('#selectfile1');
            info = $('#info1');
            btnSend = $('#pub1');
            break;

        case "csv2json":
            finput = $('#selectfile2');
            info = $('#info2');
            btnSend = $('#pub2');
            break;

        default:
            alert("need correct service name [privacy, sif2json csv2json]");
            return;
    }

    finput.change(function (evt) {
        var files = evt.target.files; // files is a FileList of File objects. List some properties.
        var output = [];
        for (var i = 0, f; f = files[i]; i++) {
            output.push('<li><strong>', escape(f.name), '</strong> (', f.type || 'n/a', ') - ',
                f.size, ' bytes, last modified: ',
                f.lastModifiedDate ? f.lastModifiedDate.toLocaleDateString() : 'n/a',
                '</li>');
        }
        info.html('<ul>' + output.join('') + '</ul>');

        filename = finput.val();
        var disabled;
        if (name == 'sif2json') {
            disabled = (!filename || (!filename.endsWith('.xml') && !filename.endsWith('.json')))
        } else if (name == 'privacy') {
            disabled = (!filename || (!filename.endsWith('.json')))
        } else if (name == 'csv2json') {
            disabled = (!filename || (!filename.endsWith('.csv')))
        }
        btnSend.prop('disabled', disabled);
    });
}

function form_submit(name) {

    var uploadform;
    switch (name) {
        case "sif2json":
            uploadform = $("#uploadform0");
            break;

        case "privacy":
            uploadform = $("#uploadform1");
            break;

        case "csv2json":
            uploadform = $("#uploadform2");
            break;

        default:
            alert("need correct service name [privacy, sif2json csv2json]");
            return;
    }

    uploadform.on('submit', function (e) {
        e.preventDefault();  // avoid to execute the actual submit of the form

        var ip = location.host;
        var url = '';
        var file = $('input[type=file]')[0].files[0]

        switch (name) {
            case "sif2json":
                ip = '192.168.31.233:1324';
                if (filename.endsWith('.xml')) {
                    url = 'http://' + ip + '/sif2json/0.1.0'
                } else if (filename.endsWith('.json')) {
                    url = 'http://' + ip + '/json2sif/0.1.0'
                }
                break;

            case "privacy":
                ip = '192.168.31.233:1323';
                break;

            case "csv2json":
                ip = '192.168.31.233:1325';
                break;

            default:
                alert("need correct service name [privacy, sif2json csv2json]");
                return;
        }


        // -------------------------------------------------

        var xhr = new XMLHttpRequest;
        xhr.open('POST', url, true);
        // xhr.setRequestHeader("Authorization", "Basic " + btoa("user" + ":" + "user"));
        xhr.setRequestHeader('X-Requested-With', 'XMLHttpRequest');
        xhr.setRequestHeader('X-File-Name', file.name);
        xhr.setRequestHeader('Content-Type', file.type || 'application/octet-stream');
        xhr.onreadystatechange = function () {
            if (xhr.readyState === 4 && xhr.status === 200) { // 0:UNSENT  1:OPENED  2:HEADERS_RECEIVED  3:LOADING  4:DONE
                var response = xhr.responseText;
                response = txt2xml(response);
                console.log(response);
            }
        };
        xhr.send(file);

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
    select_change("sif2json");
    form_submit("sif2json");

    select_change("privacy");
    form_submit("privacy");

    select_change("csv2json");
    form_submit("csv2json");
}
