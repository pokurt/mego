package rooms

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildRoomTree(t *testing.T) {
	roomList = [][]string{
		{"code", "building", "zone", "size", "display_name"},
		{"PAX-Hobby-8", "building 1", "Zone10", "8", "Hobby"},
		{"PAX-XHobby-9", "building 1", "Zone10", "9", "XHobby"},
		{"PAX-Floppy-10", "building 1", "Zone2", "10", "Floppy"},
		{"PAX-Zloopy-10", "building 2", "Zone2", "10", "Zloopy"},
		{"PAX-XZloopy-12", "building 2", "Zone2", "12", "XZloopy"},
		{"PAX-XXZloopy-12", "building 2", "Zone2", "12", "XXZloopy"},
	}

	buildRoomTree()

	logRootTreeForTest()

	assert.Len(t, roomTree.Root, 2)
	assert.Equal(t, *roomTree, Root{Root: []Node{
		{
			Key:   "building 1",
			Label: "building 1",
			Children: []Node{
				{
					Key:   "Zone10",
					Label: "Zone10",
					Children: []Node{
						{
							Key:   "8",
							Label: "8",
							Children: []Node{
								{
									Key:   "PAX-Hobby-8",
									Label: "Hobby",
								},
							},
						},
						{
							Key:   "9",
							Label: "9",
							Children: []Node{
								{
									Key:   "PAX-XHobby-9",
									Label: "XHobby",
								},
							},
						},
					},
				},
				{
					Key:   "Zone2",
					Label: "Zone2",
					Children: []Node{
						{
							Key:   "10",
							Label: "10",
							Children: []Node{
								{
									Key:   "PAX-Floppy-10",
									Label: "Floppy",
								},
							},
						},
					},
				},
			},
		},
		{
			Key:   "building 2",
			Label: "building 2",
			Children: []Node{
				{
					Key:   "Zone2",
					Label: "Zone2",
					Children: []Node{
						{
							Key:   "10",
							Label: "10",
							Children: []Node{
								{
									Key:   "PAX-Zloopy-10",
									Label: "Zloopy",
								},
							},
						},
						{
							Key:   "12",
							Label: "12",
							Children: []Node{
								{
									Key:   "PAX-XZloopy-12",
									Label: "XZloopy",
								}, {
									Key:   "PAX-XXZloopy-12",
									Label: "XXZloopy",
								},
							},
						},
					},
				},
			},
		},
	}})
}

func logRootTreeForTest() {
	for _, a := range roomTree.Root {
		fmt.Println(a.Key)

		for _, aa := range a.Children {
			fmt.Println("\t", aa.Key)

			for _, aaa := range aa.Children {
				fmt.Println("\t\t", aaa.Key)

				for _, aaaa := range aaa.Children {
					fmt.Println("\t\t\t", aaaa.Key, aaaa.Label)
				}
			}
		}
	}
}
