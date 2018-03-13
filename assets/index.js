fetch('/hello',{
    method:'GET'
}).then(res => res.text())
.then(text => console.log(text));