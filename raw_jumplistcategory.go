package electron

import "github.com/gopherjs/gopherjs/js"

// JumpListCategory a Structure
type JumpListCategory struct {
	*js.Object
	// One of the following:
	Type JumpListCategoryJumpListCategoryType `js:"type"`
	// Must be set if is , otherwise it should be omitted.
	Name string `js:"name"`
	// Array of objects if is or , otherwise it should be omitted.
	Items *js.Object `js:"items"`
}

type JumpListCategoryJumpListCategoryType string

// consts
const (
	JumpListCategoryJumpListCategoryTypeTasks    JumpListCategoryJumpListCategoryType = "tasks"
	JumpListCategoryJumpListCategoryTypeFrequent JumpListCategoryJumpListCategoryType = "frequent"
	JumpListCategoryJumpListCategoryTypeRecent   JumpListCategoryJumpListCategoryType = "recent"
	JumpListCategoryJumpListCategoryTypeCustom   JumpListCategoryJumpListCategoryType = "custom"
)
