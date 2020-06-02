function init_upload(name) {
    var finput, form;    

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
        var info = document.getElementById('info');
        info.innerHTML = ('<ul>' + output.join('') + '</ul>');
    });

    form.addEventListener('submit', function (e) {
        e.preventDefault();  // avoid to execute the actual submit of the form
        
        var add_params = prompt("&fn=*&user=*&ctx=*&rw=*&object=* | &sv=* ...", "");   
                
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
    init_upload("privacy");
    init_upload("sif2json");
    init_upload("csv2json");
}
