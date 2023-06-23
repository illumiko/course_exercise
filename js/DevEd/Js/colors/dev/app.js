//imported function
const lightOrDark = (color) => {
  let element, bgColor, brightness, r, g, b, hsp;

  // Check the format of the color, HEX or RGB?
  if (color.match(/^rgb/)) {
    // If HEX --> store the red, green, blue values in separate variables
    color = color.match(
      /^rgba?\((\d+),\s*(\d+),\s*(\d+)(?:,\s*(\d+(?:\.\d+)?))?\)$/
    );

    r = color[1];
    g = color[2];
    b = color[3];
  } else {
    // If RGB --> Convert it to HEX: http://gist.github.com/983661
    color = +("0x" + color.slice(1).replace(color.length < 5 && /./g, "$&$&"));

    r = color >> 16;
    g = (color >> 8) & 255;
    b = color & 255;
  }

  // HSP (Highly Sensitive Poo) equation from http://alienryderflex.com/hsp.html
  hsp = Math.sqrt(0.299 * (r * r) + 0.587 * (g * g) + 0.114 * (b * b));

  // Using the HSP value, determine whether the color is light or dark
  if (hsp > 127.5) {
    return "light";
  } else {
    return "dark";
  }
};

const colorToHex = (color) => {
    if (color.substr(0, 1) === '#') {
        return color;
    }
    var digits = /(.*?)rgb\((\d+), (\d+), (\d+)\)/.exec(color);
    
    var red = parseInt(digits[2]);
    var green = parseInt(digits[3]);
    var blue = parseInt(digits[4]);
    
    var rgb = blue | (green << 8) | (red << 16);
    return digits[1] + '#' + rgb.toString(16).padStart(6, '0');
};



const lock = document.querySelectorAll(".lock");
const colorChange = document.querySelectorAll(".inputColor");
const save = document.querySelector(".save");
const library = document.querySelector(".library");
const gen = document.querySelector(".gen");
const closeLibrary = document.querySelector(".closeLibrary");
const colorValAll = document.querySelectorAll(".colorVal");
const colorShowAll = document.querySelectorAll(".colorShow");
/*
const colorVal1 = document.querySelector('.color1 .colorVal')
const colorVal2 = document.querySelector('.color2 .colorVal')
const colorVal3 = document.querySelector('.color3 .colorVal')
const colorVal4 = document.querySelector('.color4 .colorVal')
const colorShow1 = document.querySelector('.color1 .colorShow')
const colorShow2 = document.querySelector('.color2 .colorShow')
const colorShow3 = document.querySelector('.color3 .colorShow')
const colorShow4 = document.querySelector('.color4 .colorShow')
*/

//functions

const FOReach = (array) => {
  array.forEach((i) => {
    return i;
  });
};

const loadLibraryOnce = () => {

  let createdElementContainer = document.querySelector(".libraryContainer .library");
  const paletteArr = []

  for (let i = 0; i < localStorage.length; i++) {
    let x = localStorage.getItem(`${localStorage.key(i)}`);
    let y = JSON.parse(x);

    const e = 
      `<div class="${localStorage.key(i)} savedPalettes">
          <h3 class="paletteName">${localStorage.key(i)}</h3>
          <div class="palette">
            <div style='background:${y[0]}'></div>
            <div style='background:${y[1]}'></div>
            <div style='background:${y[2]}'></div>
            <div style='background:${y[3]}'></div>
          </div>
      </div>`;
    paletteArr.unshift(e)
  }
  paletteArr
    .sort()
    .forEach(i => {
    createdElementContainer.innerHTML += i
  })

}

const appendLibrary = (x,y) => {

  let createdElementContainer = document.querySelector(".libraryContainer .library");
    createdElementContainer.innerHTML += 
      `<div class="${x} savedPalettes">
          <h3 class="paletteName">${x}</h3>
          <div class="palette">
            <div style='background:${y[0]}'></div>
            <div style='background:${y[1]}'></div>
            <div style='background:${y[2]}'></div>
            <div style='background:${y[3]}'></div>
          </div>
      </div>`;
}

const generatePalelet = () => {
  const letters = "0123456789ABCDEF";
  let generatedColor = "#";
  for (let i = 0; i < 6; i++) {
    generatedColor += `${letters[Math.floor(Math.random() * 16)]}`;
  }
  //console.log(generatedColor)
  return generatedColor;
};

const showGeneratedColor = () => {
  colorShowAll.forEach((i) => {
    let x = generatePalelet();

    //checks if the colorShow is locked or not if not generates color if yes then continues to
    //generate color
    if (!i.firstElementChild.children[1].classList.contains("locked")) {
      i.style.backgroundColor = `${x}`;
      i.lastElementChild.firstElementChild.innerText = x;
      i.lastElementChild.lastElementChild.value = x;
      //checking if generated color is light or dark
      if (lightOrDark(x) == "light") {
        //changing color of text
        i.style.color = "black";
      } else {
        //changing color of text
        i.style.color = "white";
        //changing color of lock icon
      }
    }
  });
};

const lockToggle = (e) => {
  e.stopPropagation();
  e.target.parentElement.parentElement.classList.toggle("locked");

  //ensures that if clicked outside the desired place it doesnt do anything
  if (e.target.parentElement.parentElement.classList.contains("lock")) {
    //checks if the lock is locked and changes icon and state accordingly
    if (e.target.parentElement.parentElement.classList.contains("locked")) {
      e.target.parentElement.parentElement.innerHTML = `
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M18 10v-4c0-3.313-2.687-6-6-6s-6 2.687-6 6v4h-3v14h18v-14h-3zm-10 0v-4c0-2.206 1.794-4 4-4s4 1.794 4 4v4h-8z"/></svg>
      `;
    } else {
      e.target.parentElement.parentElement.innerHTML = `
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M12 10v-4c0-3.313-2.687-6-6-6s-6 2.687-6 6v3h2v-3c0-2.206 1.794-4 4-4s4 1.794 4 4v4h-4v14h18v-14h-12z"/></svg>
      `;
    }
  }
};

const liveChange = (e) => {
  if (!e.target.parentElement.children[1].classList.contains("locked")) {
    e.target.parentElement.parentElement.style.background = e.target.value;
    e.target.parentElement.firstElementChild.innerText = e.target.value;
  }
};

const savePalelet = () => {
  let savedPaletteName = prompt("Enter Name")
  let savedPalette = [];
  colorValAll.forEach((i) => {
    savedPalette.unshift(i.innerText);
  });
  localStorage.setItem(savedPaletteName, JSON.stringify(savedPalette));
  appendLibrary(savedPaletteName, savedPalette)
};

const closingLibrary = (e) => {
  let targetedElement = e.target.parentElement.parentElement;
  targetedElement.classList.add("dissappear");
  targetedElement.addEventListener("animationend", () => {
    targetedElement.classList.remove("dissappear");
    targetedElement.classList.add("displayNone");
    //targetedElement.style.display = "none";
  });
};

const showingLibrary = () => {
  let targetedElement = document.querySelector(".libraryContainer");
  targetedElement.classList.remove("displayNone");

};


const loadPalette = (e) => {
  const palette = Array.from(e.target.lastElementChild.children)
  for (let i = 0; i < 4; i++){
    colorShowAll[i].style.background = palette[i].style.background 
    colorValAll[i].innerText = colorToHex(palette[i].style.background.toString())
    //console.log(colorValAll[i].innerText)
  }
}

//functions that load during initial page load
loadLibraryOnce()


//event add
save.addEventListener("click", savePalelet);

library.addEventListener("click", showingLibrary);

gen.addEventListener("click", showGeneratedColor);

lock.forEach((i) => {
  i.addEventListener("click", lockToggle);
});

colorChange.forEach((i) => {
  i.addEventListener("input", liveChange);
});

closeLibrary.addEventListener("click", closingLibrary);

window.setInterval(() => {
  const x = document.querySelectorAll(".libraryContainer .library > div");
  x.forEach(i => i.addEventListener('click', loadPalette))
}, 1500)


//copying in single tap
document.querySelectorAll(".colorShow").forEach((i) => {
  i.addEventListener("click", (e) => {
    //e.stopPropagation();
    if (e.target.classList.contains("colorAdjustments")) {
      //visual feedback for copying
      const adder = () => {
        if (
          document.querySelector(".copied").classList.contains("displayNone")
        ) {
          document.querySelector(".copied").classList.remove("displayNone");
        } else {
          document.querySelector(".copied").classList.add("displayNone");
        }
      };
      const a = window.setInterval(adder, 50);
      window.setTimeout(() => {
        window.clearInterval(a);
      }, 110);
      //copies text
      document.execCommand("Copy");
    }
  });
});
