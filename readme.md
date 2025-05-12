# Tasking

You are to write a program that generates 11 random lowercase letters, and then checks to see if it makes the word shakespear. Points are equivalent to how many consecuitive letters are correct:

sh********* = 2 points
*hakespeare = 0 points
shake****** = 5 points

Your program should output the best attempt after 5 mins. Making the word is quite basic, the challenge is optimising your solution. The more attempts made the more likely you are to win.

## Solutions

1. Naive approach to just generate as many values as possible and after each iteration evaluate them. 
    1a. Have one function to evaluate the word and return the score
    1b. Have one function to generate the word and return a string/list of chars.

2. Now that naive approach is working, create a workflow to evaluate and run the program for 5 minutes and collect all the outputs.
    1a. Are we collecting and generating enough attempts? How close do we get?
    1b. How can we optimise and match strings better?
