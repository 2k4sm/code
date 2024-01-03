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


const main_cont = document.querySelector(".main-container");
let tcol = document.querySelector(".ticket-color");

let red = document.querySelector(".red");
let green = document.querySelector(".green");
let gray = document.querySelector(".gray");
let purple = document.querySelector(".purple");

let bcol;

red.addEventListener("click",function(){
    bcol = "red";
})

green.addEventListener("click",function(){
    bcol = "green";
})

gray.addEventListener("click",function(){
    bcol= "gray";
})

purple.addEventListener("click",function(){
    bcol = "purple";
})


function genRandId(){
    sym = "@#$%^&*()!"
    alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
    num = "1234567890";


}
let counter = 0;
modal.addEventListener("keypress",function(event){
    let text_area = document.querySelector("textarea");
    if (event.key === "Enter"){
        counter++;
        let new_ticket = document.createElement("div");
        new_ticket.className = "ticket-cont";
        
        new_ticket.id = `container-${counter}`;
        
        let ticket_color = document.createElement("div");
        ticket_color.className = "ticket-color";
        
        let ticket_id = document.createElement("div");
        ticket_id.className = "ticket-id";
        

        let ticket_area = document.createElement("div");
        ticket_area.className = "ticket-area";
        ticket_area.id = `textarea-${counter}`;
        ticket_area.innerText = text_area.value;
        text_area.value = "";

        new_ticket.appendChild(ticket_color);
        new_ticket.appendChild(ticket_id);
        new_ticket.appendChild(ticket_area);

        main_cont.appendChild(new_ticket);
        modal.style.display = 'none';
    }
})