let masaslider = document.getElementById("masalar");
let masalar = document.getElementsByClassName("masa");
let masaID = document.getElementById("search").value;
function ara(){
    let sonuc = false;
    for(let i = 0;i < masalar.length;i++){
        if(masalar[i].getAttribute("masano") == kw.bind("search")){
            masalar[i].style.display = "inline-block";
            sonuc = true;
        } else {
            masalar[i].style.display = "none";   
        }
    }
}
setInterval(()=>{
    console.clear();
},10000)
