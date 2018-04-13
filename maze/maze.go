package main

import (
	"fmt"
	"os"
)

func ReadMaze(file string) [][]int {
	fmt.Println("fullfile", file)
	openFile, e := os.Open(file)
	if e != nil {
		panic(e)
	}

	var row, col, end int
	fmt.Fscanf(openFile, "%d %d", &row, &col)

	fmt.Println(row, col)
	fmt.Fscanf(openFile, "%d", &end) //它每个行末读进了一个0 将第一行末尾的0读到end变量中舍弃
	//这里把读取的数据后面的换行去掉，对于Mac是"\r"，Linux下面
	//是"\n"，Windows下面是"\r\n"，所以为了支持多平台，直接用
	//"\r\n"作为过滤字符
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(openFile, "%d", &maze[i][j])
			//fmt.Printf("%3d", maze[i][j])
		}
		fmt.Fscanf(openFile, "%d", &end) // 每一行末尾的0读到end变量中舍弃
	}
	//fmt.Println(maze)
	return maze
}

type point struct {
	i int
	j int
}

var dirs = [4]point{ //探寻next point 的四个方向
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (p point) Add(q point) point {
	return point{p.i + q.i, p.j + q.j}
}

func WalkMaze(maze [][]int, start, end point) ([][]int, int) {
	//record the steps  of walking the maze
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	queue := []point{start} //建立队列存放当前探寻到的点 从start开始

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			break
		}
		for _, dir := range dirs {
			next := current.Add(dir)
			if next.j < 0 || next.j >= len(maze[0]) || //超出边界，跳过探索下一点
				next.i < 0 || next.i >= len(maze) ||
				next == start || //回到原点，跳过，探索下一点
				maze[next.i][next.j] == 1 || //撞墙了，跳过，探索下一点
				steps[next.i][next.j] != 0 { //代表已经走过这个点了
				continue
			}
			steps[next.i][next.j] = steps[current.i][current.j] + 1 //计步器+1
			queue = append(queue, next)
		}
	}
	return steps, steps[len(maze)-1][len(maze[0])-1]

}

func main() {
	/*dir, e := filepath.Abs(filepath.Dir(os.Args[0]))
	if e != nil {
		panic(e)
	} else {
		fmt.Println("now path :", dir)
	}*/

	/*	ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		fmt.Println("exPath :", exPath)*/

	getenv := os.Getenv("GOPATH")
	fmt.Println(getenv)

	//get current dir
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(currentDir)
	}

	maze := ReadMaze(currentDir + "\\maze\\maze.in")

	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}

	steps, steps_num := WalkMaze(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	//fmt.Println(steps)
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
	fmt.Println("The min steps of the maze is ", steps_num)

	var que = []point{end}
	/*for i := len(steps) - 1; i > 0; i-- {
		for j := len(steps[0]) - 1; j > 0; j-- {
			que = append(que, point{i, j})
			if steps[i][j] > 0 {
				if steps[i][j-1] == steps[i][j]-1 {
					que = append(que, point{i, j - 1})
				}
				if steps[i-1][j] == steps[i][j]-1 {
					que = append(que, point{i - 1, j})
				}
			}
		}
	}*/

	currend_point := end
	//for !(currend_point.i == 1 && currend_point.j == 0) && !(currend_point.i == 0 && currend_point.j == 1) {
	for currend_point != start { //未到起点循环探寻
		for _, dir := range dirs {
			next := currend_point.Add(dir)
			if next.j < 0 || next.j >= len(steps[0]) || //超出边界，跳过探索下一点
				next.i < 0 || next.i >= len(steps) ||
				(steps[next.i][next.j] == 0 && next != start) || //未走过的点且不是start起点，跳过，探索下一点
				steps[next.i][next.j] != steps[currend_point.i][currend_point.j]-1 { //走到此点的步数==next点的步数+1
				continue
			}
			que = append(que, next)
			currend_point = next

		}
	}

	que = append(que, start)

	fmt.Println("The shortest path is:")
	for i := len(que) - 1; i >= 0; i-- {
		fmt.Printf("{%d,%d} ", que[i].i, que[i].j)
	}

}
