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

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (p point) Add(q point) point {
	return point{p.i + q.i, p.j + q.j}
}

func WalkMaze(maze [][]int, start, end point) [][]int {
	//record the steps  of walking the maze
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	queue := []point{start} //建立队列

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == end {
			break
		}
		for _, dir := range dirs {
			next := current.Add(dir)
			if next.j < 0 || next.j >= len(maze[0]) || next.i < 0 || next.i >= len(maze) || next == start || steps[current.i][current.j] != 0 {
				continue
			}
			steps[next.i][next.j] = steps[current.i][current.j] + 1
			queue = append(queue, next)
		}
	}
	return steps

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

	currentDir, err := os.Getwd() //get current dir
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

	steps := WalkMaze(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	fmt.Println(steps)
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
