package recursion

import "fmt"

const numDisks = 3

func push(post []int, disk int) []int {
	post = append([]int{disk}, post...)
	return post
}

func pop(post []int) (int, []int) {
	item := post[0]
	post = post[1:]
	return item, post
}

func moveDisk(posts [][]int, fromPost int, toPost int) {
	var disk int
	disk, posts[fromPost] = pop(posts[fromPost])
	posts[toPost] = push(posts[toPost], disk)
}

func drawPosts(posts [][]int) {
	// add 0s to the end of each post so they all have numDisks entries
	for p := 0; p < 3; p++ {
		for len(posts[p]) < numDisks {
			posts[p] = push(posts[p], 0)
		}
	}
	// draw the posts
	for row := 0; row < numDisks; row++ {
		for p := 0; p < 3; p++ {
			fmt.Printf("%d ", posts[p][row])
		}
		fmt.Println()
	}
	// draw a line between moves
	fmt.Println("-----")
	// remove the 0s
	for p := 0; p < 3; p++ {
		for len(posts[p]) > 0 && posts[p][0] == 0 {
			_, posts[p] = pop(posts[p])
		}
	}
}

func moveDisks(posts [][]int, numToMove int, fromPost int, toPost int, tempPost int) {
	if numToMove > 1 {
		moveDisks(posts, numToMove-1, fromPost, tempPost, toPost)
	}
	moveDisk(posts, fromPost, toPost)
	drawPosts(posts)
	if numToMove > 1 {
		moveDisks(posts, numToMove-1, tempPost, toPost, fromPost)
	}
}
