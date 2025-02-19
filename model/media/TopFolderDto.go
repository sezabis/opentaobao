package media

// TopFolderDto 结构体
type TopFolderDto struct {
	// 文件夹名称
	FolderName string `json:"folder_name,omitempty" xml:"folder_name,omitempty"`
	// 父文件夹id，顶级文件夹的父文件夹id为null。顶级文件夹只有一个。
	ParentFolderId int64 `json:"parent_folder_id,omitempty" xml:"parent_folder_id,omitempty"`
	// 文件夹id
	FolderId int64 `json:"folder_id,omitempty" xml:"folder_id,omitempty"`
	// 文件夹的层级，顶级文件夹为0
	Level int64 `json:"level,omitempty" xml:"level,omitempty"`
}
