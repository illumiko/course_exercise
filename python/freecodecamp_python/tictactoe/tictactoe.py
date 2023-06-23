import random
def makeBoard(num=False):
    if num:
        return [str(x) for x in range(9)]
    else:
        return ["_" for x in range(9)]

def formatBoard(board):
    board = [board[i*3:(i+1)*3] for i  in range(3)]
    for i in range(3):
        print(" | " + " | ".join(board[i]) + " | ")


def updateBoard(board ,position, character):
    board[position] = character

def checkWin(board):
    row = [board[(i*3):(i+1)*3] for i in range(3)]
    col1 = [board[i] for i in range(9) if i%3==0]
    col2 = [board[i] for i in range(9) if i%3==1]
    col3 = [board[i] for i in range(9) if i%3==2]
    for i in [col1,col2,col3]:
        if i == ["x","x","x"] :
            return True
        elif i == ["o","o","o"] :
            return True
    for i in row:
        if i == ["x","x","x"] :
            return True
        elif i == ["o","o","o"] :
            return True

def playerInput():
    num = input("Enter a number corresponding to the sqaure you want to play: ").lower()
    return int(num)

def compInput(board):
    filteredBoard = [i for i,c in enumerate(board) if c != "o" and c !="x"]
    position = random.choice(filteredBoard)
    return position

def playerTurn(board):
    pI = playerInput()
    if board[pI] != "x" and board[pI] != "o":
        updateBoard(board,pI,'x')
        # print(checkWin(board))
        if checkWin(board):
            return True
    else:
        while board[pI] == "x" and board[pI] == "o":
            updateBoard(board,pI,'x')
            # print(checkWin(board))
            if checkWin(board):
                return True
def compTurn(board):
    updateBoard(board,compInput(board), 'o')
    # print(checkWin(board))
    if checkWin(board):
        return True

winner = False
board = makeBoard()
numBoard = makeBoard(True)
i = 0
while not winner:
    formatBoard(board)
    print("------------------------------------------------")
    formatBoard(numBoard)
    if playerTurn(board):
        print("player Wins")
        winner = True
    elif compTurn(board):
        print("computer Wins")
        winner = True










