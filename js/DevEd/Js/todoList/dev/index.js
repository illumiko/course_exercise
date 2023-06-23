//for declaring vars
const buttonInputAdd = document.querySelector(".add-todo");
const buttonInputRemove = document.querySelector(".clear-task");
const nameInput = document.querySelector("#form-nameInput");
const todoList = document.querySelector(".todoList-list ul");
//const todoListItems = document.querySelectorAll(".todoItems");

let todoListArr = [];

//function that adds styling when user clicks done
const toggleCompleteTask = (e) => {
  e.target.classList.toggle("todoItems-Clicked");
};

const removeItem = (e) => {
  localStorage.removeItem("todo");
  location.reload();
};

//for loading added items to local storage
const loadState = () => {
  if (localStorage.length !== 0) {
    //if (JSON.parse(localStorage.getItem("todo")).length !== 0) {
    JSON.parse(localStorage.getItem("todo")).forEach((i) => {
      todoListArr.push(i);
      todoList.innerHTML += `
		<li class="todoItems">${i} </li>

		`;
      //for adding styling when the task is done by user ON LOAD
      Object.values(todoList.children).forEach((i) =>
        i.addEventListener("click", toggleCompleteTask)
      );
    });
    console.log(todoListArr);
  } else {
    console.log("nothing to load");
  }
};
loadState();

//adding item from name field and adding item on enter press
const whichkey = (e) => {
  e.target.preventDefault;
  if (e.keyCode === 13) {
    const todoListLength = document.querySelectorAll(".todoItems").length;
    todoListArr.push(nameInput.value);
    updateLocalStorage();
    todoList.innerHTML += `
		<li class="todoItems"> ${
      JSON.parse(localStorage.getItem("todo"))[todoListLength]
    }</li>
		`;
  }
};

//for updating localStorage
const updateLocalStorage = () => {
  localStorage.setItem("todo", JSON.stringify(todoListArr));
};

//for adding new items and updating local storage
const addItemsFunc = (e) => {
  e.target.preventDefault;
  const todoListLength = document.querySelectorAll(".todoItems").length;
  todoListArr.push(nameInput.value);
  updateLocalStorage();
  todoList.innerHTML += `
		<li class="todoItems"> ${
      JSON.parse(localStorage.getItem("todo"))[todoListLength]
    } </li>
		`;

  //for adding styling when the task is done by user
  Object.values(todoList.children).forEach((i) =>
    i.addEventListener("click", toggleCompleteTask)
  );
};
buttonInputAdd.addEventListener("click", addItemsFunc);
nameInput.addEventListener("keypress", whichkey);
buttonInputRemove.addEventListener("click", removeItem);
