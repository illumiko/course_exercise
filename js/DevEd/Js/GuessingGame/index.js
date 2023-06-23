let resultOutput = document.querySelector("#results");

const guessGame = () => {
  let userInput = document.querySelector("#userInput").value;
  const randomNr = (min, max) => {
    return Math.round(Math.random() * (max - min) + min);
  };

  const genNr = randomNr(1, 10);

    if (genNr === parseInt(userInput)) {
      resultOutput.style.color = "#BEF264";
      resultOutput.innerHTML = `Number generated: ${genNr}  Number guessed: ${userInput} YOU WIN`;
    } else {
      resultOutput.style.color = "#FCA6A5";
      resultOutput.innerHTML = `Number generated: ${genNr}  Number guessed: ${userInput} 
	  YOU LOSE`;
    }
};
document.querySelector("#guess").addEventListener("click", guessGame);
