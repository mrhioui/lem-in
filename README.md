<h3 align="center">
Lem-in Project
</h3>

<p align="center">
This project is meant to make a digital version of an ant farm.<br>
The program reads from a file (describing the ants and the colony) given in the arguments.<br>
Upon successfully finding the quickest path, the program displays the content of the file passed as argument and each move the ants make from room to room.
</p>

## How does it work?

The argument file present an ant farm with tunnels and rooms.
We place the ants on one side and look at how they find the exit using the quickest way to get N ants across a colony (composed of rooms and tunnels).

At the beginning of the game, all the ants are in the room `##start`. The goal is to bring them to the room `##end` with as few moves as possible.
The shortest path is not necessarily the simplest.
Some colonies may have many rooms and many links, but no path between `##start` and `##end`.
Some may have rooms that link to themselves, sending the path-search spinning in circles. Some will have too many/too few ants, no `##start` or `##end`, duplicated rooms, links to unknown rooms, rooms with invalid coordinates and a variety of other invalid or poorly-formatted input. In those cases the program will return an error message `ERROR: invalid data format`. The are more specific error messages (example: `ERROR: the number of ants is not correct or missing` or `ERROR: invalid Tunnul format, expected format: 'RoomA-RoomB'`).

We display the results on the standard output in the following format :
```
(Input)
number_of_ants
the_rooms
the_links

(Solution)
Lx-y Lz-w Lr-o ...
```
- x, z, r represents the ants numbers (going from 1 to number_of_ants) and y, w, o represents the rooms names.

- A room is defined by "name coord_x coord_y", and will usually look like "Room 1 2", "nameoftheroom 1 6", "4 6 7".

- The links are defined by "name1-name2" and will usually look like "1-2", "2-5".

## Example in practice

```
7
##start
1 23 3
2 16 7
#comment
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
#another comment
4-2
2-1
7-6
7-2
7-4
6-5

L1-3 L2-2
L1-4 L2-7 L3-3 L4-2
L1-0 L2-6 L3-4 L4-7 L5-3 L6-2
L2-0 L3-0 L4-6 L5-4 L6-7 L7-3
L4-0 L5-0 L6-6 L7-4
L6-0 L7-0
```

- As we can see, after all the ants emerged from `##start` (room 1), they reached `##end` (room 0).

## Best combination algorithm

The algorithm calculates the minimum number of turns (referred to as "height") required for a group of ants to traverse a set of paths in various combinations. It iterates through all combinations of paths, calculates the time needed to equalize the lengths of the paths by filling them to the length of the longest path, and determines the remaining ants to distribute (`left`). The "height" is computed as the sum of the longest path length and the additional number of turns required to evenly distribute the remaining ants across the paths, with any remainder adding an extra unit of turns. Each computed height is stored in a map, linking it to the corresponding combination of paths. Finally, the algorithm finds the smallest height, which represents the fastest configuration for all ants to reach their destination.

- **Number of turns' formula**

```
turns := max + left/paths (floored division)
```

**Turns**: total number of all ants steps.  
**Max**: length of longest path in combination.  
**Left**: remaining ants, after even distribution.  
**Paths**: number of paths in the combination.  

- **Literal example**

```
Input:
ants = 10
combined = [[[1, 2], [1, 2, 3]], [[1, 2, 3], [1, 2, 3, 4]]]

Step-by-Step Execution,
For cmb = [[1, 2], [1, 2, 3]]:

paths = 2, max = 3
fill = (3-2) + (3-3) = 1
left = 10 - 1 = 9
turns = max + left/paths = 3 + 9/2 = 3 + 4 = 7
Store: heights[7] = [[1, 2], [1, 2, 3]], min = 7

Step-by-Step Execution,
For cmb = [[1, 2, 3], [1, 2, 3, 4]]:

paths = 2, max = 4
fill = (4-3) + (4-4) = 1
left = 10 - 1 = 9
turns = max + left/paths = 4 + 9/2 = 4 + 4 = 8
Store: heights[8] = [[1, 2, 3], [1, 2, 3, 4]], min = 7

Final Result:
heights = {7: [[1, 2], [1, 2, 3]], 8: [[1, 2, 3], [1, 2, 3, 4]]}
min = 7 (Minimal number of turns).
```

## Instructions for argument file

- You need to create tunnels and rooms.
- A room will never start with the letter `L` or with `#` and must have no spaces.
- You join the rooms together with as many tunnels as you need.
- A tunnel joins only two rooms together never more than that.
- A room can be linked to an infinite number of rooms and by as many tunnels as deemed necessary.
- Each room can only contain one ant at a time (except at ##start and ##end which can contain as many ants as necessary).
- To be the first to arrive, ants will need to take the shortest path or paths. They will also need to avoid traffic jams as well as walking all over their fellow ants.
- We display only the ants that moved at each turn, and each ant can move only once and through a tunnel (the room at the receiving end must be empty).
- The rooms names will not necessarily be numbers, and in order.
- Any unknown command will be ignored.
- The coordinates of the rooms should always be int.

## Authors

- [Sadiqui Abdelilah](https://learn.zone01oujda.ma/git/asadiqui)
- [Saadane Aymane](https://learn.zone01oujda.ma/git/asaadane)
- [Rhioui Mohamed](https://learn.zone01oujda.ma/git/mrhioui)
- [Otchoun Agiel](https://learn.zone01oujda.ma/git/aotchoun)

## License

This project is open-sourced under [the MIT License](https://opensource.org/license/mit).
