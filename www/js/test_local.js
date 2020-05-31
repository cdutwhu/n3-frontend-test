function init_local(name) {

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
                disabled |= (!fname.endsWith('.csv') && !fname.endsWith('.json'))
                break;
        }
        btnSend.disabled = disabled;
    });

    uploadform.addEventListener('submit', function (e) {
        e.preventDefault();  // avoid to execute the actual submit of the form
        
        var add_params = prompt("Enter additional parameters (&user= &ctx= &rw= | &sv= ...)", "");       
        
        fetch(
            'http://' + location.host + '/upload/?service=' + name + add_params, {
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
    });
}

window.onload = function () {
    init_local("privacy");
    init_local("sif2json");
    init_local("csv2json");
}
