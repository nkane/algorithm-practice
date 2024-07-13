package recursion

import "testing"

func TestTowerOfHanoi(t *testing.T) {
	posts := [][]int{}
	posts = append(posts, []int{})
	for disk := numDisks; disk > 0; disk-- {
		posts[0] = push(posts[0], disk)
	}
	for p := 1; p < 3; p++ {
		posts = append(posts, []int{})
	}
	drawPosts(posts)
	moveDisks(posts, numDisks, 0, 1, 2)
}
