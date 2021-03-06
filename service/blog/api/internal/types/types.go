// Code generated by goctl. DO NOT EDIT.
package types

type GetBlogListReq struct {
	UserID int64 `form:"user_id, optional"`
	Page   int   `form:"page, default=1"`
	Size   int   `form:"size, default=20"`
}

type GetBlogListResp struct {
	Blogs []Blog `json:"blogs"`
}

type Blog struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`    // 作者ID
	Title     string `json:"title"`      // 博客标题/文件夹名
	IsFolder  bool   `json:"is_folder"`  // 是否是文件夹
	Content   string `json:"content"`    // 博客正文
	Status    int32  `json:"status"`     // 博客状态 1:所有人可见 2. 仅自己可见 3. 删除状态
	FolderID  int64  `json:"folderID"`   // 博客所属文件夹，0 则无类别或者说根目录
	CreatedAt int64  `json:"created_at"` // 创建时间
	UpdatedAt int64  `json:"updated_at"` // 更新时间
}
