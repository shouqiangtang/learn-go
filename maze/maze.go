package maze

import (
	"errors"
	"fmt"
	"os"
)

// Load : 加载源数据并存入二维数组
func Load() ([][]int, error) {
	f, err := os.Open("maze.in")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// 获取行数列数
	var rows, cols int
	fmt.Fscanf(f, "%d %d", &rows, &cols)

	// 申请内存
	data := make([][]int, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			fmt.Fscanf(f, "%d", &data[i][j])
		}
	}

	return data, nil
}

// Point : 表示迷宫里的格子
type Point struct {
	X, Y int
}

// 方向结点
var directions []Point = []Point{
	Point{0, -1}, // 上
	Point{-1, 0}, // 左
	Point{0, 1},  // 下
	Point{1, 0},  // 右
}

// Add : 两个结点相加
func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

// Equal : 判断两个结点是否相等
func (p Point) Equal(other Point) bool {
	return p == other
}

// Available : 判断结点是否有效
func (p Point) Available(maze, footPrints [][]int, start Point) bool {
	// fmt.Printf("%#v\n", p)
	// 判断是否越界
	rows, cols := len(maze), len(maze[0])
	if p.X < 0 || p.Y < 0 {
		return false
	}
	if p.X >= rows || p.Y >= cols {
		return false
	}

	// 不能是起点
	if p == start {
		return false
	}

	// 判断是不是墙
	if maze[p.X][p.Y] == 1 {
		return false
	}

	// 判断是否已走过
	if footPrints[p.X][p.Y] > 0 {
		return false
	}

	return true
}

// Walk : 走迷宫
func Walk(maze [][]int, start, end Point) ([][]int, error) {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return nil, errors.New("迷宫地图错误")
	}
	rows, cols := len(maze), len(maze[0])

	// 生成足迹表
	footPrints := make([][]int, rows)
	for i := 0; i < rows; i++ {
		footPrints[i] = make([]int, cols)
	}

	// 队列，用于存放待探索结点
	var queue []Point
	queue = append(queue, start)

	for {
		if len(queue) == 0 {
			break
		}
		curPoint := queue[0]
		queue = queue[1:]
		if curPoint.Equal(end) {
			break
		}

		// 获取当前结点足迹
		curVal := footPrints[curPoint.X][curPoint.Y]
		for _, direction := range directions {
			todoPoint := curPoint.Add(direction)
			if todoPoint.Available(maze, footPrints, start) {
				queue = append(queue, todoPoint)
				footPrints[todoPoint.X][todoPoint.Y] = curVal + 1
			}
		}
	}
	return footPrints, nil
}
