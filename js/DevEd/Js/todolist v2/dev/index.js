//for declaring vars
const buttonInputAdd = document.querySelector(".add-todo");
const nameInput = document.querySelector("#form-nameInput");
const todoList = document.querySelector(".todoList-list ul");
//const todoListItems = document.querySelectorAll(".todoItems");

console.log(todoList)

//adding items
const addItemsFunc = () => {
  todoList.innerHTML += `
    <li class="todoItems"> <p>${nameInput.value}</p> <button class="clear-task">X</button></li>
    `;
  popEffect();

  //setting eventlistener for remove btn
  document
    .querySelectorAll(".todoList-list ul li button")
    .forEach((i) => i.addEventListener("click", removeItem));

  document.querySelectorAll(".todoItems").forEach((i) => {
    i.addEventListener("click", () => {
      i.classList.toggle("todoItems-completed");
    });
  });

  //calling function
  saveLocalStorage();
};

//pop up effect when new item is added
const popEffect = () => {
  //removing preceeding todoItems-pop
  for (let i = 0; i < todoList.children.length; i++) {
    todoList.children[i].classList.remove("todoItems-pop");
  }
  //add todoItems-pop to the latest item added
  todoList.lastElementChild.classList.add("todoItems-pop");
};

//remove item
const removeItem = (e) => {
    e.stopPropagation();
  //removing todoItems-pop as it clashes with todoItems-remove
  e.target.parentElement.classList.remove("todoItems-pop");
  //animation to play when item deleted
  e.target.parentElement.classList.add("todoItems-remove");
  //callback function after animation ends
  e.target.parentElement.addEventListener("animationend", () => {
    //e.target.parentElement.classList.remove('todoItems-remove')
    //removes item
    todoList.removeChild(e.target.parentElement);
  });
  deleteLocalStorage(e);
};

//changing default behaviour of input
const whichkey = (e) => {
  e.preventDefault;
  if (e.keyCode == 13) {
    addItemsFunc();
  }
};

//local storage
//saving

const saveLocalStorage = () => {
  let todoArr;
  if (localStorage.length == 0) {
    todoArr = [];
  } else {
    todoArr = JSON.parse(localStorage.getItem("todo"));
  }
  todoArr.push(nameInput.value);
  localStorage.setItem("todo", JSON.stringify(todoArr));
  //console.log('hello')
};
//deleting

const deleteLocalStorage = (e) => {
  let todoArr;
  if (localStorage.length == 0) {
    todoArr = [];
  } else {
    todoArr = JSON.parse(localStorage.getItem("todo"));
  }

  selectedElementData =
    e.target.parentElement.firstElementChild.firstChild.data;

  let todoArrIndexOf = todoArr.indexOf(`${selectedElementData}`);
  //console.log(todoArrIndexOf);

  todoArr.splice(todoArrIndexOf, 1);
  localStorage.setItem("todo", JSON.stringify(todoArr));
};

//loading saved storage
const loadStorage = () => {
  if (localStorage.length == 0) {
    todoArr = [];
  } else {
    todoArr = JSON.parse(localStorage.getItem("todo"));
  }
  todoArr.forEach((i) => {
    todoList.innerHTML += `
    <li class="todoItems"> <p>${i}</p> <button class="clear-task">X</button></li>
    `;
  });
  document
    .querySelectorAll(".todoList-list ul li button")
    .forEach((i) => i.addEventListener("click", removeItem));

  document.querySelectorAll(".todoItems").forEach((i) => {
    //i.stopPropagation();
    i.addEventListener("click", () => {
      i.classList.toggle("todoItems-completed");
    });
  });
};
loadStorage();

buttonInputAdd.addEventListener("click", addItemsFunc);
nameInput.addEventListener("keypress", whichkey);
