const inputBtn = document.querySelector('#inputBtn')
const input = document.querySelector('#inputText')
const gl = document.querySelector('#goalList')
console.log(inputBtn, input, gl)

inputBtn.addEventListener('click', () => {
  gl.innerHTML += `
        <div class="bg-gray-200 mt-4 relative ">
          <div class="h-2 w-2 inset-y-3 -left-4 absolute rounded-full bg-black"></div>
          <p class="text-center text-2xl"> ${input.value}</p>
        </div>
    `
})
