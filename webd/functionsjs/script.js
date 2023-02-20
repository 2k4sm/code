let h1 = document.querySelector("#head")
let add7 = number => (h1.innerHTML = number+7);
function multiply(n1,n2){
    return h1.innerHTML = n1*n2
}
function capitalize(word){
    word = word.toLowerCase()
    word = word.charAt(0).toUpperCase()+word.slice(1)
    
    return h1.innerHTML = word
}
let lastLetter = word => h1.innerHTML = word[word.length-1]



