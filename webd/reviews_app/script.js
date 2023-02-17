const apiLink = "https://api.themoviedb.org/3/discover/movie?sort_by=popularity.desc&api_key=5253a2cb4574d77ba564b7780e34f00e&page=1";
const imgPath = "https://image.tmdb.org/t/p/w1280";
const searchApi = "https://api.themoviedb.org/3/search/movie?&api_key=5253a2cb4574d77ba564b7780e34f00e&query=";


const main = document.getElementById("section-2")
const form = document.getElementById("form")
const search = document.getElementById("search")

returnMovies(apiLink)

function returnMovies(url){
    fetch(url).then(res => res.json())
    .then(function(data){
        console.log(data.results)
        data.results.forEach(element => {
            const div_card = document.createElement('div')
            div_card.setAttribute('class','card')
            const div_row = document.createElement('div')
            div_card.setAttribute('class','row')

            const div_column = document.createElement('div')
            div_card.setAttribute('class','column')

            const image = document.createElement('img')
            div_card.setAttribute('class','thumbnail')
            div_card.setAttribute('id','image')

            const title = document.createElement('h3')
            div_card.setAttribute('id','title')

            title.innerHTML = $`{element.title}`
            image.src = imgPath + element.poster_path

            div_card.appendChild(image)
            div_card.appendChild(title)
            div_column.appendChild(div_column)
            div_row.appendChild(div_column)
            main.appendChild(div_row)
        });
    });
}
form.addEventListener("submit",(e) => {
    e.preventDefault()
    main.innerHTML = ""
    const searchItem = search.value
    
    if (searchItem){
        returnMovies(searchApi+searchItem)
        search.value = ""
    }
})