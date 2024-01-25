console.log(1)
function sum(a,b){
    return new Promise(resolve => {
        setTimeout(() => {
            resolve(a+b);
        },5000)
    });
}
async function get(){
    console.log(2);
    // let ans = await fetch("https://www.github.com/");
    let ans = await sum(2,4);
    console.log(ans);
    console.log(3);
}
get();
console.log(4);