const add = document.querySelector("#plus");
const modal = document.querySelector(".modal-container");
add.addEventListener("click", ()=>{
    if(modal.style.display === "none"){
        modal.style.display = 'flex';
    }
    else{
        modal.style.display = 'none';
    }
})