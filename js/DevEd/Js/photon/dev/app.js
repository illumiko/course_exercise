class Picture {
  constructor() {
    this._c = document.querySelector(".picShowcase");
    this._button = document.querySelector(".btn");
    this._currentCuratedPage = 1;
    this._searchInput = document.querySelector(".searchInput");
    this._nextpage
    this._auth = "563492ad6f91700001000001625f5b73945a4df58dcd30b24d559869";
  }

  getAuth() {
    return this._auth;
  }

  //curated releted
  curatedUrl(pg=1) {
    return `https://api.pexels.com/v1/curated/?page=${pg}&per_page=5`;
  }
  setCurrentCuratedPage(v) {
    return (this._currentCuratedPage = v);
  }

  getCurrentCuratedPage() {
    return this._currentCuratedPage;
  }

  incrementedCuratedPage() {
    return this._currentCuratedPage+=1 ;
  }

  //search related
  setCurrentSearchedPage(v) {
    return this._currentSearchedPage = v;
  }

  getCurrentSearchedPage() {
    return this._currentSearchedPage;
  }

  getSearchInput() {
    return this._searchInput;
  }
  getSearchInputValue () {
    return this._searchInput.value
  }

  incrementedSearchPage() {
    return this._currentSearchedPage+= 1;
  }


  searchUrl(query, pg = 1) {
    //return `https://api.pexels.com/v1/search?query=${query}&page=${pg}&per_page=5`;
    return `https://api.pexels.com/v1/search?query=${query}&page=${pg}&per_page=5`;
  }

  get container() {
    return this._c;
  }

  get btn() {
    return this._button;
  }

  async getData(fetchUrl) {
    try {

    const getting = await fetch(`${fetchUrl}`, {
      method: "GET",
      headers: {
        Accept: "application/json",
        Authorization: this.getAuth(),
      },
    });
    const gotten = await getting.json();
    return gotten;
    } catch (e) {
      document.querySelector('body').innerHTML = `
      <h1 class='error'>SERVER LIMIT REACHED</h1>
      `
    }
  }

  createCuratedPic() {
    let container = this.container;
    this.getData(this.curatedUrl(this.getCurrentCuratedPage())).then((v) => {
      console.log(v)
      v.photos.forEach((i) => {
        ///*
        container.innerHTML += `
      <div class='picContainer'>
        <img class='picDisplay' src='${i.src.portrait}'/>
        <div class=picDetails>
          <a target='_blank' href='${i.photographer_url}' class='authorName'>Photographer:
          ${i.photographer}</a>
        </div>
      </div>
    `;
        //*/
      });
    });
  }

  createSearchedPic(url) {
    let container = this.container;
    console.log(this.getCurrentSearchedPage())
    this.getData(url).then(
      (v) => {
        this._nextpage = v.next_page
        v.photos.forEach((i) => {
          container.innerHTML += `
      <div class='picContainer'>
        <img class='picDisplay' src='${i.src.portrait}'/>
        <div class=picDetails>
          <a target='_blank' href='${i.photographer_url}' class='authorName'>Photographer:
          ${i.photographer}</a>
        </div>
      </div>
    `;
        });
      }
    );
  }

  clear() {
    let container = this.container;
    container.innerHTML = "";
    this._nextpage = ''
    this.setCurrentSearchedPage(1);
    //this.setCurrentCuratedPage(1);
  }
}
const pic = new Picture();
//initial load
pic.createCuratedPic();
//for Curated imgs
pic.btn.addEventListener("click", (e) => {
  if(!pic.btn.classList.contains('searching')){
    pic.setCurrentCuratedPage(pic.incrementedCuratedPage());
    pic.createCuratedPic();
  }else {
    pic.createSearchedPic(pic._nextpage)
  }
});

pic.getSearchInput().addEventListener("change", (e) => {
  if (e.target.value) {
    pic.clear()
    pic.btn.classList.add('searching')
    //pic.createSearchedPic(e.target.value, this.getCurrentSearchedPage);
    pic.createSearchedPic(pic.searchUrl(e.target.value))
  } else {
    pic.clear()
    pic.btn.classList.remove('searching')
    pic.createCuratedPic()
  }
});


//pic.createPic()
//pic.createPic()
//
//
