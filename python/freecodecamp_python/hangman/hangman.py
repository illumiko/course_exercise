from words import words
import random
randomIndex = random.randrange(len(words))
randomWord = words[randomIndex]
print([word for word in randomWord if word == "-"])
filteredWord = ["_" for x in randomWord]
winCondition = False
i = 0
while not winCondition:
    print(filteredWord)
    userInput = input("Please enter a letter: ")
    if userInput in randomWord:
        correctIndex = [i for i,c in enumerate(randomWord) if c == userInput]
        for i in correctIndex:
            filteredWord[i] = userInput
            if "_" not in filteredWord:
                winCondition = True
                print("Congrats you guessed the word","[",randomWord,"]")
           

