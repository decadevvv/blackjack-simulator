# BlackJack Simulator

This is a Golang project designed to test BlackJack strategies.

4 strategies are already implemented:

1. never-explode: never split and double, only hit on hard point < 12
2. dealer: mimic dealer, always hit on point < 17
3. basic: the basic strategy introduced in the book `Beat the Dealer: A Winning Strategy for the Game of Twenty-One` by `Edward Thorp`.
4. ask-user: you can also play with the simulator, here is an example

How to build:

```shell
go build -o blackjack-simulator.exe
```

How to check flags:

```shell
./blackjack-simulator.exe -h
```

How to run 10000 rounds of simulation using the `basic` strategy:

```shell
./blackjack-simulator.exe -r 10000 -s basic
```

Example output of running 10 rounds of simulation using the `basic` strategy:

```shell
$ .\blackjack-simulator.exe -r 10 -s basic
round 1 lose: [5 9 9] (23) V.S. [Q 6 9] (25) (player balance 0.0 - 1.0 => -1.0)       
round 2 lose: [9 J] (19) V.S. [9 3 9] (21) (player balance -1.0 - 1.0 => -2.0)        
round 3 lose: [4 5 4 K] (23) V.S. [K 10] (20) (player balance -2.0 - 1.0 => -3.0)     
round 4 lose: [7 3 2 2 10] (24) V.S. [K 8] (18) (player balance -3.0 - 1.0 => -4.0)   
round 5 win: [8 J] (18) V.S. [4 6 7] (17) (player balance -4.0 + 1.0 => -3.0)         
round 6 push: [7 2 J] (19) V.S. [8 A] (19) (player balance -3.0 unchanged)            
round 7 win: [10 Q] (20) V.S. [2 2 2 Q A] (17) (player balance -3.0 + 1.0 => -2.0)    
round 8 win: [A 5 8 A 6] (21) V.S. [7 Q] (17) (player balance -2.0 + 1.0 => -1.0)     
round 9 blackjack: [A Q] (21) V.S. [J Q] (20) (player balance -1.0 + 1.5 => 0.5)      
round 10 double win: [2 9 10] (21) V.S. [J 4 Q] (24) (player balance 0.5 + 2.0 => 2.5)
-----statistics-----
10 rounds, 10 hands, 5 win, 1 push, 4 lose
BlackJack analysis: 1 blackjacks (10.00%), 1 win, 0 push, margin 1.5, edge 150.00%, contribution 15.00%
Hit stand analysis: 8 hit/stand hands (80%), 3 win, 1 push, 4 lose, 37.50% win rate, margin -1, edge -12.50%, contribution -10.00%
Hit/Stand/BlackJack analysis: 9 hit/stand/blackjack hands (90%), 4 win, 1 push, 4 lose, 44.44% win rate, margin 0.5, edge 5.56%, contribition 5.00%
Double hand analysis: 1 doubles hands (10.00%), 1 win, 0 push, 0 lose, 100.00% win rate, margin 2, edge 200.00%, contribution 20.00%
Split analysis: 0 hands could split (0.00%), 0 splits (NaN%), 0 win, 0 push, 0 lose, NaN% win rate, margin 0, edge NaN%, contribution 0.00%
player edge = 25.000%
```

How to add new strategies:

1. Implement the `Strategy` interface defined in this repo.
2. Then run the `main.go` file to see the result.
