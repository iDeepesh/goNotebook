Pots of gold game: Two players A & B. There are pots of gold arranged in a line, each containing some gold coins (the players can see how many coins are there in each gold pot - perfect information). They get alternating turns in which the player can pick a pot from one of the ends of the line. The winner is the player which has a higher number of coins at the end. The objective is to "maximize" the number of coins collected by A, assuming B also plays optimally. A starts the game. 

The idea is to find an optimal strategy that makes A win knowing that B is playing optimally as well. How would you do that?

*Solution:*

For player A : he can pick up coin in first place or last place, for each of these case player B can further pick from first or last place from remaining coins. But for A to win/maximize coins, coins collected by B should be minimum while that of A should be maximized. 

If A selects coin[i], then B can choose coin[i+1] or coin[array_size] 

If A selects coin[array_last], then B can choose coin[i] or coin[array_last-1] 

We will simulate such that every function call will start from A's turn. This will give us the given recursive function. The reason for the min function being used is that the opponent B is also going to want to maximize B's winning when picking the gold pot, essentially leaving A with 
the minimum of the two options to choose from in the next round.